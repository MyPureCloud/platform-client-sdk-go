package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Footertemplate - The Footer template identifies the Footer type and its footerUsage
type Footertemplate struct { 
	// VarType - Specifies the type represented by Footer.
	VarType *string `json:"type,omitempty"`


	// ApplicableResources - Specifies the canned response template where the footer can be used.
	ApplicableResources *[]string `json:"applicableResources,omitempty"`

}

func (o *Footertemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Footertemplate
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ApplicableResources *[]string `json:"applicableResources,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		ApplicableResources: o.ApplicableResources,
		Alias:    (*Alias)(o),
	})
}

func (o *Footertemplate) UnmarshalJSON(b []byte) error {
	var FootertemplateMap map[string]interface{}
	err := json.Unmarshal(b, &FootertemplateMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := FootertemplateMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if ApplicableResources, ok := FootertemplateMap["applicableResources"].([]interface{}); ok {
		ApplicableResourcesString, _ := json.Marshal(ApplicableResources)
		json.Unmarshal(ApplicableResourcesString, &o.ApplicableResources)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Footertemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
