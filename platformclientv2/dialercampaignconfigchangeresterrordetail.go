package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignconfigchangeresterrordetail
type Dialercampaignconfigchangeresterrordetail struct { 
	// VarError - name of the error
	VarError *string `json:"error,omitempty"`


	// Details - additional information regarding the error
	Details *string `json:"details,omitempty"`

}

func (o *Dialercampaignconfigchangeresterrordetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignconfigchangeresterrordetail
	
	return json.Marshal(&struct { 
		VarError *string `json:"error,omitempty"`
		
		Details *string `json:"details,omitempty"`
		*Alias
	}{ 
		VarError: o.VarError,
		
		Details: o.Details,
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangeresterrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
