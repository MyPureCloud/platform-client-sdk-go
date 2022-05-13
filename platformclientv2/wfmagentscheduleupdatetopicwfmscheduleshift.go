package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmscheduleshift
type Wfmagentscheduleupdatetopicwfmscheduleshift struct { 
	// WeekDate
	WeekDate *string `json:"weekDate,omitempty"`


	// WeekScheduleId
	WeekScheduleId *string `json:"weekScheduleId,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Activities
	Activities *[]Wfmagentscheduleupdatetopicwfmscheduleactivity `json:"activities,omitempty"`

}

func (o *Wfmagentscheduleupdatetopicwfmscheduleshift) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagentscheduleupdatetopicwfmscheduleshift
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		WeekDate *string `json:"weekDate,omitempty"`
		
		WeekScheduleId *string `json:"weekScheduleId,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Activities *[]Wfmagentscheduleupdatetopicwfmscheduleactivity `json:"activities,omitempty"`
		*Alias
	}{ 
		WeekDate: o.WeekDate,
		
		WeekScheduleId: o.WeekScheduleId,
		
		Id: o.Id,
		
		StartDate: StartDate,
		
		LengthInMinutes: o.LengthInMinutes,
		
		Activities: o.Activities,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmagentscheduleupdatetopicwfmscheduleshift) UnmarshalJSON(b []byte) error {
	var WfmagentscheduleupdatetopicwfmscheduleshiftMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentscheduleupdatetopicwfmscheduleshiftMap)
	if err != nil {
		return err
	}
	
	if WeekDate, ok := WfmagentscheduleupdatetopicwfmscheduleshiftMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
    
	if WeekScheduleId, ok := WfmagentscheduleupdatetopicwfmscheduleshiftMap["weekScheduleId"].(string); ok {
		o.WeekScheduleId = &WeekScheduleId
	}
    
	if Id, ok := WfmagentscheduleupdatetopicwfmscheduleshiftMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if startDateString, ok := WfmagentscheduleupdatetopicwfmscheduleshiftMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthInMinutes, ok := WfmagentscheduleupdatetopicwfmscheduleshiftMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if Activities, ok := WfmagentscheduleupdatetopicwfmscheduleshiftMap["activities"].([]interface{}); ok {
		ActivitiesString, _ := json.Marshal(Activities)
		json.Unmarshal(ActivitiesString, &o.Activities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmscheduleshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
