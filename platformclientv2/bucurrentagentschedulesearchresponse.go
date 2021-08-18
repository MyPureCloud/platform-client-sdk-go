package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bucurrentagentschedulesearchresponse
type Bucurrentagentschedulesearchresponse struct { 
	// AgentSchedules - The requested agent schedules
	AgentSchedules *[]Buagentschedulesearchresponse `json:"agentSchedules,omitempty"`


	// BusinessUnitTimeZone - The time zone configured for the business unit to which this schedule applies
	BusinessUnitTimeZone *string `json:"businessUnitTimeZone,omitempty"`


	// PublishedSchedules - References to all published week schedules overlapping the start/end date query parameters
	PublishedSchedules *[]Buagentschedulepublishedschedulereference `json:"publishedSchedules,omitempty"`


	// StartDate - The start date of the schedules. Only populated on notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end date of the schedules. Only populated on notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// Updates - The list of updates for the schedule. Only used in notifications
	Updates *[]Buagentscheduleupdate `json:"updates,omitempty"`

}

func (u *Bucurrentagentschedulesearchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bucurrentagentschedulesearchresponse

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if u.EndDate != nil {
		
		*EndDate = timeutil.Strftime(u.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	

	return json.Marshal(&struct { 
		AgentSchedules *[]Buagentschedulesearchresponse `json:"agentSchedules,omitempty"`
		
		BusinessUnitTimeZone *string `json:"businessUnitTimeZone,omitempty"`
		
		PublishedSchedules *[]Buagentschedulepublishedschedulereference `json:"publishedSchedules,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		Updates *[]Buagentscheduleupdate `json:"updates,omitempty"`
		*Alias
	}{ 
		AgentSchedules: u.AgentSchedules,
		
		BusinessUnitTimeZone: u.BusinessUnitTimeZone,
		
		PublishedSchedules: u.PublishedSchedules,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Updates: u.Updates,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bucurrentagentschedulesearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
