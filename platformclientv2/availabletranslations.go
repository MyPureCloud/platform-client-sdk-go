package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availabletranslations
type Availabletranslations struct { 
	// OrgSpecific
	OrgSpecific *[]string `json:"orgSpecific,omitempty"`


	// Builtin
	Builtin *[]string `json:"builtin,omitempty"`

}

func (o *Availabletranslations) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availabletranslations
	
	return json.Marshal(&struct { 
		OrgSpecific *[]string `json:"orgSpecific,omitempty"`
		
		Builtin *[]string `json:"builtin,omitempty"`
		*Alias
	}{ 
		OrgSpecific: o.OrgSpecific,
		
		Builtin: o.Builtin,
		Alias:    (*Alias)(o),
	})
}

func (o *Availabletranslations) UnmarshalJSON(b []byte) error {
	var AvailabletranslationsMap map[string]interface{}
	err := json.Unmarshal(b, &AvailabletranslationsMap)
	if err != nil {
		return err
	}
	
	if OrgSpecific, ok := AvailabletranslationsMap["orgSpecific"].([]interface{}); ok {
		OrgSpecificString, _ := json.Marshal(OrgSpecific)
		json.Unmarshal(OrgSpecificString, &o.OrgSpecific)
	}
	
	if Builtin, ok := AvailabletranslationsMap["builtin"].([]interface{}); ok {
		BuiltinString, _ := json.Marshal(Builtin)
		json.Unmarshal(BuiltinString, &o.Builtin)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Availabletranslations) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
