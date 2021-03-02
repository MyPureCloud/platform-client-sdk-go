package platformclientv2
import (
	"encoding/json"
)

// Createmanagementunitapirequest - Create Management Unit
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


	// BusinessUnitId - The id of the business unit to which this management unit belongs.  Required after business unit launch
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createmanagementunitapirequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
