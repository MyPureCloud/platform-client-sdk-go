package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmscheduleactivity
type Wfmagentscheduleupdatetopicwfmscheduleactivity struct { 
	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// CountsAsPaidTime
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`


	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// TimeOffRequestId
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`

}

func (o *Wfmagentscheduleupdatetopicwfmscheduleactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagentscheduleupdatetopicwfmscheduleactivity
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		ActivityCodeId: o.ActivityCodeId,
		
		StartDate: StartDate,
		
		CountsAsPaidTime: o.CountsAsPaidTime,
		
		LengthInMinutes: o.LengthInMinutes,
		
		TimeOffRequestId: o.TimeOffRequestId,
		
		Description: o.Description,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmagentscheduleupdatetopicwfmscheduleactivity) UnmarshalJSON(b []byte) error {
	var WfmagentscheduleupdatetopicwfmscheduleactivityMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentscheduleupdatetopicwfmscheduleactivityMap)
	if err != nil {
		return err
	}
	
	if ActivityCodeId, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
	
	if startDateString, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if CountsAsPaidTime, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
	
	if LengthInMinutes, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if TimeOffRequestId, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
	
	if Description, ok := WfmagentscheduleupdatetopicwfmscheduleactivityMap["description"].(string); ok {
		o.Description = &Description
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
