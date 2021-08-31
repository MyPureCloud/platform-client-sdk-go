package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatebusinessunitrequest
type Updatebusinessunitrequest struct { 
	// Name - The name of the business unit
	Name *string `json:"name,omitempty"`


	// DivisionId - The ID of the division to which the business unit should be moved
	DivisionId *string `json:"divisionId,omitempty"`


	// Settings - Configuration for the business unit
	Settings *Updatebusinessunitsettings `json:"settings,omitempty"`

}

func (o *Updatebusinessunitrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatebusinessunitrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		Settings *Updatebusinessunitsettings `json:"settings,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		DivisionId: o.DivisionId,
		
		Settings: o.Settings,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatebusinessunitrequest) UnmarshalJSON(b []byte) error {
	var UpdatebusinessunitrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatebusinessunitrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := UpdatebusinessunitrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if DivisionId, ok := UpdatebusinessunitrequestMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
	
	if Settings, ok := UpdatebusinessunitrequestMap["settings"].(map[string]interface{}); ok {
		SettingsString, _ := json.Marshal(Settings)
		json.Unmarshal(SettingsString, &o.Settings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatebusinessunitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
