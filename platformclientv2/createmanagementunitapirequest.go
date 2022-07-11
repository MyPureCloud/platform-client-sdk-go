package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createmanagementunitapirequest
type Createmanagementunitapirequest struct { 
	// Name - The name of the management unit
	Name *string `json:"name,omitempty"`


	// TimeZone - The default time zone to use for this management unit.  Moving to Business Unit
	TimeZone *string `json:"timeZone,omitempty"`


	// StartDayOfWeek - The configured first day of the week for scheduling and forecasting purposes. Moving to Business Unit
	StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`


	// Settings - The configuration for the management unit.  If omitted, reasonable defaults will be assigned
	Settings *Createmanagementunitsettingsrequest `json:"settings,omitempty"`


	// DivisionId - The id of the division to which this management unit belongs.  Defaults to home division ID
	DivisionId *string `json:"divisionId,omitempty"`


	// BusinessUnitId - The id of the business unit to which this management unit belongs
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

}

func (o *Createmanagementunitapirequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createmanagementunitapirequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`
		
		Settings *Createmanagementunitsettingsrequest `json:"settings,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		TimeZone: o.TimeZone,
		
		StartDayOfWeek: o.StartDayOfWeek,
		
		Settings: o.Settings,
		
		DivisionId: o.DivisionId,
		
		BusinessUnitId: o.BusinessUnitId,
		Alias:    (*Alias)(o),
	})
}

func (o *Createmanagementunitapirequest) UnmarshalJSON(b []byte) error {
	var CreatemanagementunitapirequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatemanagementunitapirequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreatemanagementunitapirequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if TimeZone, ok := CreatemanagementunitapirequestMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if StartDayOfWeek, ok := CreatemanagementunitapirequestMap["startDayOfWeek"].(string); ok {
		o.StartDayOfWeek = &StartDayOfWeek
	}
    
	if Settings, ok := CreatemanagementunitapirequestMap["settings"].(map[string]interface{}); ok {
		SettingsString, _ := json.Marshal(Settings)
		json.Unmarshal(SettingsString, &o.Settings)
	}
	
	if DivisionId, ok := CreatemanagementunitapirequestMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if BusinessUnitId, ok := CreatemanagementunitapirequestMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createmanagementunitapirequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
