package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Punctualityevent
type Punctualityevent struct { 
	// DateScheduleStart - The scheduled activity start time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateScheduleStart *time.Time `json:"dateScheduleStart,omitempty"`


	// DateStart - The time the user started the activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthMinutes - The length of the activity in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Description - The description of the activity
	Description *string `json:"description,omitempty"`


	// ActivityCodeId - The ID of the activity code associated with this activity
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// ActivityCode - The activity code
	ActivityCode *string `json:"activityCode,omitempty"`


	// ActivityName - The activity name
	ActivityName *string `json:"activityName,omitempty"`


	// Category - The category for the activity
	Category *string `json:"category,omitempty"`


	// Points - The points earned for this activity
	Points *int `json:"points,omitempty"`


	// Delta - Difference between this activity and the last activity in seconds
	Delta *float64 `json:"delta,omitempty"`


	// Bullseye
	Bullseye *bool `json:"bullseye,omitempty"`

}

func (o *Punctualityevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Punctualityevent
	
	DateScheduleStart := new(string)
	if o.DateScheduleStart != nil {
		
		*DateScheduleStart = timeutil.Strftime(o.DateScheduleStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateScheduleStart = nil
	}
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		DateScheduleStart *string `json:"dateScheduleStart,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		ActivityCode *string `json:"activityCode,omitempty"`
		
		ActivityName *string `json:"activityName,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		Points *int `json:"points,omitempty"`
		
		Delta *float64 `json:"delta,omitempty"`
		
		Bullseye *bool `json:"bullseye,omitempty"`
		*Alias
	}{ 
		DateScheduleStart: DateScheduleStart,
		
		DateStart: DateStart,
		
		LengthMinutes: o.LengthMinutes,
		
		Description: o.Description,
		
		ActivityCodeId: o.ActivityCodeId,
		
		ActivityCode: o.ActivityCode,
		
		ActivityName: o.ActivityName,
		
		Category: o.Category,
		
		Points: o.Points,
		
		Delta: o.Delta,
		
		Bullseye: o.Bullseye,
		Alias:    (*Alias)(o),
	})
}

func (o *Punctualityevent) UnmarshalJSON(b []byte) error {
	var PunctualityeventMap map[string]interface{}
	err := json.Unmarshal(b, &PunctualityeventMap)
	if err != nil {
		return err
	}
	
	if dateScheduleStartString, ok := PunctualityeventMap["dateScheduleStart"].(string); ok {
		DateScheduleStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateScheduleStartString)
		o.DateScheduleStart = &DateScheduleStart
	}
	
	if dateStartString, ok := PunctualityeventMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if LengthMinutes, ok := PunctualityeventMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Description, ok := PunctualityeventMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ActivityCodeId, ok := PunctualityeventMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if ActivityCode, ok := PunctualityeventMap["activityCode"].(string); ok {
		o.ActivityCode = &ActivityCode
	}
    
	if ActivityName, ok := PunctualityeventMap["activityName"].(string); ok {
		o.ActivityName = &ActivityName
	}
    
	if Category, ok := PunctualityeventMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if Points, ok := PunctualityeventMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	
	if Delta, ok := PunctualityeventMap["delta"].(float64); ok {
		o.Delta = &Delta
	}
    
	if Bullseye, ok := PunctualityeventMap["bullseye"].(bool); ok {
		o.Bullseye = &Bullseye
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Punctualityevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
