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

func (o *Createmanagementunitsettingsrequest) MarshalJSON() ([]byte, error) {
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
		Adherence: o.Adherence,
		
		ShortTermForecasting: o.ShortTermForecasting,
		
		TimeOff: o.TimeOff,
		
		Scheduling: o.Scheduling,
		
		ShiftTrading: o.ShiftTrading,
		Alias:    (*Alias)(o),
	})
}

func (o *Createmanagementunitsettingsrequest) UnmarshalJSON(b []byte) error {
	var CreatemanagementunitsettingsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatemanagementunitsettingsrequestMap)
	if err != nil {
		return err
	}
	
	if Adherence, ok := CreatemanagementunitsettingsrequestMap["adherence"].(map[string]interface{}); ok {
		AdherenceString, _ := json.Marshal(Adherence)
		json.Unmarshal(AdherenceString, &o.Adherence)
	}
	
	if ShortTermForecasting, ok := CreatemanagementunitsettingsrequestMap["shortTermForecasting"].(map[string]interface{}); ok {
		ShortTermForecastingString, _ := json.Marshal(ShortTermForecasting)
		json.Unmarshal(ShortTermForecastingString, &o.ShortTermForecasting)
	}
	
	if TimeOff, ok := CreatemanagementunitsettingsrequestMap["timeOff"].(map[string]interface{}); ok {
		TimeOffString, _ := json.Marshal(TimeOff)
		json.Unmarshal(TimeOffString, &o.TimeOff)
	}
	
	if Scheduling, ok := CreatemanagementunitsettingsrequestMap["scheduling"].(map[string]interface{}); ok {
		SchedulingString, _ := json.Marshal(Scheduling)
		json.Unmarshal(SchedulingString, &o.Scheduling)
	}
	
	if ShiftTrading, ok := CreatemanagementunitsettingsrequestMap["shiftTrading"].(map[string]interface{}); ok {
		ShiftTradingString, _ := json.Marshal(ShiftTrading)
		json.Unmarshal(ShiftTradingString, &o.ShiftTrading)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createmanagementunitsettingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
