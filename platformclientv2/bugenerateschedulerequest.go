package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bugenerateschedulerequest
type Bugenerateschedulerequest struct { 
	// Description - The description for the schedule
	Description *string `json:"description,omitempty"`


	// ShortTermForecast - The forecast to use when generating the schedule.  Note that the forecast must fully encompass the schedule's start week + week count
	ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`


	// WeekCount - The number of weeks in the schedule. One extra day is added at the end
	WeekCount *int `json:"weekCount,omitempty"`


	// Options - Additional scheduling options
	Options *Schedulingoptionsrequest `json:"options,omitempty"`

}

func (o *Bugenerateschedulerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bugenerateschedulerequest
	
	return json.Marshal(&struct { 
		Description *string `json:"description,omitempty"`
		
		ShortTermForecast *Bushorttermforecastreference `json:"shortTermForecast,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Options *Schedulingoptionsrequest `json:"options,omitempty"`
		*Alias
	}{ 
		Description: o.Description,
		
		ShortTermForecast: o.ShortTermForecast,
		
		WeekCount: o.WeekCount,
		
		Options: o.Options,
		Alias:    (*Alias)(o),
	})
}

func (o *Bugenerateschedulerequest) UnmarshalJSON(b []byte) error {
	var BugenerateschedulerequestMap map[string]interface{}
	err := json.Unmarshal(b, &BugenerateschedulerequestMap)
	if err != nil {
		return err
	}
	
	if Description, ok := BugenerateschedulerequestMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if ShortTermForecast, ok := BugenerateschedulerequestMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	
	if WeekCount, ok := BugenerateschedulerequestMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if Options, ok := BugenerateschedulerequestMap["options"].(map[string]interface{}); ok {
		OptionsString, _ := json.Marshal(Options)
		json.Unmarshal(OptionsString, &o.Options)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bugenerateschedulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
