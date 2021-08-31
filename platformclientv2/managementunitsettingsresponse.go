package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Managementunitsettingsresponse
type Managementunitsettingsresponse struct { 
	// Adherence - Adherence settings for this management unit
	Adherence *Adherencesettings `json:"adherence,omitempty"`


	// ShortTermForecasting - Short term forecasting settings for this management unit
	ShortTermForecasting *Shorttermforecastingsettings `json:"shortTermForecasting,omitempty"`


	// TimeOff - Time off request settings for this management unit
	TimeOff *Timeoffrequestsettings `json:"timeOff,omitempty"`


	// Scheduling - Scheduling settings for this management unit. These settings are only available if you have the permission wfm:managementUnit:view
	Scheduling *Schedulingsettingsresponse `json:"scheduling,omitempty"`


	// ShiftTrading - Shift trade settings for this management unit
	ShiftTrading *Shifttradesettings `json:"shiftTrading,omitempty"`


	// Metadata - Version info metadata for the associated management unit
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Managementunitsettingsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Managementunitsettingsresponse
	
	return json.Marshal(&struct { 
		Adherence *Adherencesettings `json:"adherence,omitempty"`
		
		ShortTermForecasting *Shorttermforecastingsettings `json:"shortTermForecasting,omitempty"`
		
		TimeOff *Timeoffrequestsettings `json:"timeOff,omitempty"`
		
		Scheduling *Schedulingsettingsresponse `json:"scheduling,omitempty"`
		
		ShiftTrading *Shifttradesettings `json:"shiftTrading,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Adherence: o.Adherence,
		
		ShortTermForecasting: o.ShortTermForecasting,
		
		TimeOff: o.TimeOff,
		
		Scheduling: o.Scheduling,
		
		ShiftTrading: o.ShiftTrading,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Managementunitsettingsresponse) UnmarshalJSON(b []byte) error {
	var ManagementunitsettingsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ManagementunitsettingsresponseMap)
	if err != nil {
		return err
	}
	
	if Adherence, ok := ManagementunitsettingsresponseMap["adherence"].(map[string]interface{}); ok {
		AdherenceString, _ := json.Marshal(Adherence)
		json.Unmarshal(AdherenceString, &o.Adherence)
	}
	
	if ShortTermForecasting, ok := ManagementunitsettingsresponseMap["shortTermForecasting"].(map[string]interface{}); ok {
		ShortTermForecastingString, _ := json.Marshal(ShortTermForecasting)
		json.Unmarshal(ShortTermForecastingString, &o.ShortTermForecasting)
	}
	
	if TimeOff, ok := ManagementunitsettingsresponseMap["timeOff"].(map[string]interface{}); ok {
		TimeOffString, _ := json.Marshal(TimeOff)
		json.Unmarshal(TimeOffString, &o.TimeOff)
	}
	
	if Scheduling, ok := ManagementunitsettingsresponseMap["scheduling"].(map[string]interface{}); ok {
		SchedulingString, _ := json.Marshal(Scheduling)
		json.Unmarshal(SchedulingString, &o.Scheduling)
	}
	
	if ShiftTrading, ok := ManagementunitsettingsresponseMap["shiftTrading"].(map[string]interface{}); ok {
		ShiftTradingString, _ := json.Marshal(ShiftTrading)
		json.Unmarshal(ShiftTradingString, &o.ShiftTrading)
	}
	
	if Metadata, ok := ManagementunitsettingsresponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Managementunitsettingsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
