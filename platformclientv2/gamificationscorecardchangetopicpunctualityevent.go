package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Gamificationscorecardchangetopicpunctualityevent
type Gamificationscorecardchangetopicpunctualityevent struct { 
	// DateStart
	DateStart *string `json:"dateStart,omitempty"`


	// DateScheduleStart
	DateScheduleStart *string `json:"dateScheduleStart,omitempty"`


	// ActivityCode
	ActivityCode *string `json:"activityCode,omitempty"`


	// Points
	Points *int `json:"points,omitempty"`

}

func (o *Gamificationscorecardchangetopicpunctualityevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Gamificationscorecardchangetopicpunctualityevent
	
	return json.Marshal(&struct { 
		DateStart *string `json:"dateStart,omitempty"`
		
		DateScheduleStart *string `json:"dateScheduleStart,omitempty"`
		
		ActivityCode *string `json:"activityCode,omitempty"`
		
		Points *int `json:"points,omitempty"`
		*Alias
	}{ 
		DateStart: o.DateStart,
		
		DateScheduleStart: o.DateScheduleStart,
		
		ActivityCode: o.ActivityCode,
		
		Points: o.Points,
		Alias:    (*Alias)(o),
	})
}

func (o *Gamificationscorecardchangetopicpunctualityevent) UnmarshalJSON(b []byte) error {
	var GamificationscorecardchangetopicpunctualityeventMap map[string]interface{}
	err := json.Unmarshal(b, &GamificationscorecardchangetopicpunctualityeventMap)
	if err != nil {
		return err
	}
	
	if DateStart, ok := GamificationscorecardchangetopicpunctualityeventMap["dateStart"].(string); ok {
		o.DateStart = &DateStart
	}
	
	if DateScheduleStart, ok := GamificationscorecardchangetopicpunctualityeventMap["dateScheduleStart"].(string); ok {
		o.DateScheduleStart = &DateScheduleStart
	}
	
	if ActivityCode, ok := GamificationscorecardchangetopicpunctualityeventMap["activityCode"].(string); ok {
		o.ActivityCode = &ActivityCode
	}
	
	if Points, ok := GamificationscorecardchangetopicpunctualityeventMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Gamificationscorecardchangetopicpunctualityevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
