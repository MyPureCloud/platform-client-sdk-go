package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createbusinessunitrequest
type Createbusinessunitrequest struct { 
	// Name - The name of the business unit
	Name *string `json:"name,omitempty"`


	// DivisionId - The ID of the division to which the business unit should be added
	DivisionId *string `json:"divisionId,omitempty"`


	// Settings - Configuration for the business unit
	Settings *Createbusinessunitsettings `json:"settings,omitempty"`

}

func (o *Createbusinessunitrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createbusinessunitrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		Settings *Createbusinessunitsettings `json:"settings,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		DivisionId: o.DivisionId,
		
		Settings: o.Settings,
		Alias:    (*Alias)(o),
	})
}

func (o *Createbusinessunitrequest) UnmarshalJSON(b []byte) error {
	var CreatebusinessunitrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatebusinessunitrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreatebusinessunitrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if DivisionId, ok := CreatebusinessunitrequestMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if Settings, ok := CreatebusinessunitrequestMap["settings"].(map[string]interface{}); ok {
		SettingsString, _ := json.Marshal(Settings)
		json.Unmarshal(SettingsString, &o.Settings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createbusinessunitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
