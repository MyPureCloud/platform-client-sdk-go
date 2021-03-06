package platformclientv2
import (
	"encoding/json"
)

// Createmanagementunitsettingsrequest - Management Unit Settings
type Createmanagementunitsettingsrequest struct { 
	// Adherence - Adherence settings for this management unit
	Adherence *Adherencesettings `json:"adherence,omitempty"`


	// ShortTermForecasting - Short term forecasting settings for this management unit.  Moving to Business Unit
	ShortTermForecasting *Shorttermforecastingsettings `json:"shortTermForecasting,omitempty"`


	// TimeOff - Time off request settings for this management unit
	TimeOff *Timeoffrequestsettings `json:"timeOff,omitempty"`


	// Scheduling - Scheduling settings for this management unit
	Scheduling *Schedulingsettingsrequest `json:"scheduling,omitempty"`


	// ShiftTrading - Shift trade settings for this management unit
	ShiftTrading *Shifttradesettings `json:"shiftTrading,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createmanagementunitsettingsrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
