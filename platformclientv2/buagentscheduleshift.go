package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentscheduleshift
type Buagentscheduleshift struct { 
	// Id - The ID of the shift
	Id *string `json:"id,omitempty"`


	// StartDate - The start date of this shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// LengthMinutes - The length of this shift in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`


	// Activities - The activities associated with this shift
	Activities *[]Buagentscheduleactivity `json:"activities,omitempty"`


	// ManuallyEdited - Whether this shift was manually edited. This is only set by clients and is used for rescheduling
	ManuallyEdited *bool `json:"manuallyEdited,omitempty"`


	// Schedule - The schedule to which this shift belongs
	Schedule *Buschedulereference `json:"schedule,omitempty"`

}

func (o *Buagentscheduleshift) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentscheduleshift
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		Activities *[]Buagentscheduleactivity `json:"activities,omitempty"`
		
		ManuallyEdited *bool `json:"manuallyEdited,omitempty"`
		
		Schedule *Buschedulereference `json:"schedule,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		StartDate: StartDate,
		
		LengthMinutes: o.LengthMinutes,
		
		Activities: o.Activities,
		
		ManuallyEdited: o.ManuallyEdited,
		
		Schedule: o.Schedule,
		Alias:    (*Alias)(o),
	})
}

func (o *Buagentscheduleshift) UnmarshalJSON(b []byte) error {
	var BuagentscheduleshiftMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentscheduleshiftMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BuagentscheduleshiftMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if startDateString, ok := BuagentscheduleshiftMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if LengthMinutes, ok := BuagentscheduleshiftMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if Activities, ok := BuagentscheduleshiftMap["activities"].([]interface{}); ok {
		ActivitiesString, _ := json.Marshal(Activities)
		json.Unmarshal(ActivitiesString, &o.Activities)
	}
	
	if ManuallyEdited, ok := BuagentscheduleshiftMap["manuallyEdited"].(bool); ok {
		o.ManuallyEdited = &ManuallyEdited
	}
    
	if Schedule, ok := BuagentscheduleshiftMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentscheduleshift) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
