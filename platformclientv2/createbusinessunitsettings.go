package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createbusinessunitsettings
type Createbusinessunitsettings struct { 
	// StartDayOfWeek - The start day of week for this business unit
	StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`


	// TimeZone - The time zone for this business unit, using the Olsen tz database format
	TimeZone *string `json:"timeZone,omitempty"`


	// ShortTermForecasting - Short term forecasting settings
	ShortTermForecasting *Bushorttermforecastingsettings `json:"shortTermForecasting,omitempty"`


	// Scheduling - Scheduling settings
	Scheduling *Buschedulingsettings `json:"scheduling,omitempty"`

}

func (o *Createbusinessunitsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createbusinessunitsettings
	
	return json.Marshal(&struct { 
		StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		ShortTermForecasting *Bushorttermforecastingsettings `json:"shortTermForecasting,omitempty"`
		
		Scheduling *Buschedulingsettings `json:"scheduling,omitempty"`
		*Alias
	}{ 
		StartDayOfWeek: o.StartDayOfWeek,
		
		TimeZone: o.TimeZone,
		
		ShortTermForecasting: o.ShortTermForecasting,
		
		Scheduling: o.Scheduling,
		Alias:    (*Alias)(o),
	})
}

func (o *Createbusinessunitsettings) UnmarshalJSON(b []byte) error {
	var CreatebusinessunitsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &CreatebusinessunitsettingsMap)
	if err != nil {
		return err
	}
	
	if StartDayOfWeek, ok := CreatebusinessunitsettingsMap["startDayOfWeek"].(string); ok {
		o.StartDayOfWeek = &StartDayOfWeek
	}
	
	if TimeZone, ok := CreatebusinessunitsettingsMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
	
	if ShortTermForecasting, ok := CreatebusinessunitsettingsMap["shortTermForecasting"].(map[string]interface{}); ok {
		ShortTermForecastingString, _ := json.Marshal(ShortTermForecasting)
		json.Unmarshal(ShortTermForecastingString, &o.ShortTermForecasting)
	}
	
	if Scheduling, ok := CreatebusinessunitsettingsMap["scheduling"].(map[string]interface{}); ok {
		SchedulingString, _ := json.Marshal(Scheduling)
		json.Unmarshal(SchedulingString, &o.Scheduling)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createbusinessunitsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
