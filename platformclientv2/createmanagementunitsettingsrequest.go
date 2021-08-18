package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Createmanagementunitsettingsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createmanagementunitsettingsrequest

	

	return json.Marshal(&struct { 
		Adherence *Adherencesettings `json:"adherence,omitempty"`
		
		ShortTermForecasting *Shorttermforecastingsettings `json:"shortTermForecasting,omitempty"`
		
		TimeOff *Timeoffrequestsettings `json:"timeOff,omitempty"`
		
		Scheduling *Schedulingsettingsrequest `json:"scheduling,omitempty"`
		
		ShiftTrading *Shifttradesettings `json:"shiftTrading,omitempty"`
		*Alias
	}{ 
		Adherence: u.Adherence,
		
		ShortTermForecasting: u.ShortTermForecasting,
		
		TimeOff: u.TimeOff,
		
		Scheduling: u.Scheduling,
		
		ShiftTrading: u.ShiftTrading,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createmanagementunitsettingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
