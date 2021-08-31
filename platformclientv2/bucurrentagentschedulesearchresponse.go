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

func (o *Bucurrentagentschedulesearchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bucurrentagentschedulesearchresponse
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		AgentSchedules: o.AgentSchedules,
		
		BusinessUnitTimeZone: o.BusinessUnitTimeZone,
		
		PublishedSchedules: o.PublishedSchedules,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Updates: o.Updates,
		Alias:    (*Alias)(o),
	})
}

func (o *Bucurrentagentschedulesearchresponse) UnmarshalJSON(b []byte) error {
	var BucurrentagentschedulesearchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BucurrentagentschedulesearchresponseMap)
	if err != nil {
		return err
	}
	
	if AgentSchedules, ok := BucurrentagentschedulesearchresponseMap["agentSchedules"].([]interface{}); ok {
		AgentSchedulesString, _ := json.Marshal(AgentSchedules)
		json.Unmarshal(AgentSchedulesString, &o.AgentSchedules)
	}
	
	if BusinessUnitTimeZone, ok := BucurrentagentschedulesearchresponseMap["businessUnitTimeZone"].(string); ok {
		o.BusinessUnitTimeZone = &BusinessUnitTimeZone
	}
	
	if PublishedSchedules, ok := BucurrentagentschedulesearchresponseMap["publishedSchedules"].([]interface{}); ok {
		PublishedSchedulesString, _ := json.Marshal(PublishedSchedules)
		json.Unmarshal(PublishedSchedulesString, &o.PublishedSchedules)
	}
	
	if startDateString, ok := BucurrentagentschedulesearchresponseMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := BucurrentagentschedulesearchresponseMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if Updates, ok := BucurrentagentschedulesearchresponseMap["updates"].([]interface{}); ok {
		UpdatesString, _ := json.Marshal(Updates)
		json.Unmarshal(UpdatesString, &o.Updates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bucurrentagentschedulesearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
