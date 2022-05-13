package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scheduleactivity
type Scheduleactivity struct { 
	// DateStart - The start date/time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthMinutes - The length of this activity in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Description - The description of this activity
	Description *string `json:"description,omitempty"`


	// ActivityCodeId - The ID of the activity code associated with this activity
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// Paid - Whether this activity is paid
	Paid *bool `json:"paid,omitempty"`


	// TimeOffRequestId - The ID of the time off request associated with this activity, if applicable
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`


	// ExternalActivityId - The ID of the external activity associated with this activity, if applicable
	ExternalActivityId *string `json:"externalActivityId,omitempty"`


	// ExternalActivityType - The type of the external activity associated with this activity, if applicable
	ExternalActivityType *string `json:"externalActivityType,omitempty"`

}

func (o *Scheduleactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scheduleactivity
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		ExternalActivityId *string `json:"externalActivityId,omitempty"`
		
		ExternalActivityType *string `json:"externalActivityType,omitempty"`
		*Alias
	}{ 
		DateStart: DateStart,
		
		LengthMinutes: o.LengthMinutes,
		
		Description: o.Description,
		
		ActivityCodeId: o.ActivityCodeId,
		
		Paid: o.Paid,
		
		TimeOffRequestId: o.TimeOffRequestId,
		
		ExternalActivityId: o.ExternalActivityId,
		
		ExternalActivityType: o.ExternalActivityType,
		Alias:    (*Alias)(o),
	})
}

func (o *Scheduleactivity) UnmarshalJSON(b []byte) error {
	var ScheduleactivityMap map[string]interface{}
	err := json.Unmarshal(b, &ScheduleactivityMap)
	if err != nil {
		return err
	}
	
	if dateStartString, ok := ScheduleactivityMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if LengthMinutes, ok := ScheduleactivityMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Description, ok := ScheduleactivityMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ActivityCodeId, ok := ScheduleactivityMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if Paid, ok := ScheduleactivityMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
    
	if TimeOffRequestId, ok := ScheduleactivityMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
    
	if ExternalActivityId, ok := ScheduleactivityMap["externalActivityId"].(string); ok {
		o.ExternalActivityId = &ExternalActivityId
	}
    
	if ExternalActivityType, ok := ScheduleactivityMap["externalActivityType"].(string); ok {
		o.ExternalActivityType = &ExternalActivityType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Scheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
