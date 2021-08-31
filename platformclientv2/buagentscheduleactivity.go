package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentscheduleactivity
type Buagentscheduleactivity struct { 
	// StartDate - The start date/time of this activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


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

func (o *Buagentscheduleactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentscheduleactivity
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		ExternalActivityId *string `json:"externalActivityId,omitempty"`
		
		ExternalActivityType *string `json:"externalActivityType,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
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

func (o *Buagentscheduleactivity) UnmarshalJSON(b []byte) error {
	var BuagentscheduleactivityMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentscheduleactivityMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := BuagentscheduleactivityMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := BuagentscheduleactivityMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Description, ok := BuagentscheduleactivityMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if ActivityCodeId, ok := BuagentscheduleactivityMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
	
	if Paid, ok := BuagentscheduleactivityMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
	
	if TimeOffRequestId, ok := BuagentscheduleactivityMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
	
	if ExternalActivityId, ok := BuagentscheduleactivityMap["externalActivityId"].(string); ok {
		o.ExternalActivityId = &ExternalActivityId
	}
	
	if ExternalActivityType, ok := BuagentscheduleactivityMap["externalActivityType"].(string); ok {
		o.ExternalActivityType = &ExternalActivityType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
