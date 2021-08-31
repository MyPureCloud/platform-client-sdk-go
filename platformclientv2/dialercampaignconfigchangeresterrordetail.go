package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignconfigchangeresterrordetail
type Dialercampaignconfigchangeresterrordetail struct { 
	// VarError
	VarError *string `json:"error,omitempty"`


	// Details
	Details *string `json:"details,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialercampaignconfigchangeresterrordetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignconfigchangeresterrordetail
	
	return json.Marshal(&struct { 
		VarError *string `json:"error,omitempty"`
		
		Details *string `json:"details,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		VarError: o.VarError,
		
		Details: o.Details,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignconfigchangeresterrordetail) UnmarshalJSON(b []byte) error {
	var DialercampaignconfigchangeresterrordetailMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignconfigchangeresterrordetailMap)
	if err != nil {
		return err
	}
	
	if VarError, ok := DialercampaignconfigchangeresterrordetailMap["error"].(string); ok {
		o.VarError = &VarError
	}
	
	if Details, ok := DialercampaignconfigchangeresterrordetailMap["details"].(string); ok {
		o.Details = &Details
	}
	
	if AdditionalProperties, ok := DialercampaignconfigchangeresterrordetailMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangeresterrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
