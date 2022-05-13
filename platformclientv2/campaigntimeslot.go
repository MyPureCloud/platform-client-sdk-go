package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaigntimeslot
type Campaigntimeslot struct { 
	// StartTime - The start time of the interval as an ISO-8601 string, i.e. HH:mm:ss
	StartTime *string `json:"startTime,omitempty"`


	// StopTime - The end time of the interval as an ISO-8601 string, i.e. HH:mm:ss
	StopTime *string `json:"stopTime,omitempty"`


	// Day - The day of the interval. Valid values: [1-7], representing Monday through Sunday
	Day *int `json:"day,omitempty"`

}

func (o *Campaigntimeslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaigntimeslot
	
	return json.Marshal(&struct { 
		StartTime *string `json:"startTime,omitempty"`
		
		StopTime *string `json:"stopTime,omitempty"`
		
		Day *int `json:"day,omitempty"`
		*Alias
	}{ 
		StartTime: o.StartTime,
		
		StopTime: o.StopTime,
		
		Day: o.Day,
		Alias:    (*Alias)(o),
	})
}

func (o *Campaigntimeslot) UnmarshalJSON(b []byte) error {
	var CampaigntimeslotMap map[string]interface{}
	err := json.Unmarshal(b, &CampaigntimeslotMap)
	if err != nil {
		return err
	}
	
	if StartTime, ok := CampaigntimeslotMap["startTime"].(string); ok {
		o.StartTime = &StartTime
	}
    
	if StopTime, ok := CampaigntimeslotMap["stopTime"].(string); ok {
		o.StopTime = &StopTime
	}
    
	if Day, ok := CampaigntimeslotMap["day"].(float64); ok {
		DayInt := int(Day)
		o.Day = &DayInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaigntimeslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
