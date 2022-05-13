package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulessearchresponse
type Buagentschedulessearchresponse struct { 
	// AgentSchedules - The requested agent schedules
	AgentSchedules *[]Buagentschedulesearchresponse `json:"agentSchedules,omitempty"`


	// BusinessUnitTimeZone - The time zone configured for the business unit to which this schedule applies
	BusinessUnitTimeZone *string `json:"businessUnitTimeZone,omitempty"`


	// PublishedSchedules - References to all published week schedules overlapping the start/end date query parameters
	PublishedSchedules *[]Buagentschedulepublishedschedulereference `json:"publishedSchedules,omitempty"`

}

func (o *Buagentschedulessearchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulessearchresponse
	
	return json.Marshal(&struct { 
		AgentSchedules *[]Buagentschedulesearchresponse `json:"agentSchedules,omitempty"`
		
		BusinessUnitTimeZone *string `json:"businessUnitTimeZone,omitempty"`
		
		PublishedSchedules *[]Buagentschedulepublishedschedulereference `json:"publishedSchedules,omitempty"`
		*Alias
	}{ 
		AgentSchedules: o.AgentSchedules,
		
		BusinessUnitTimeZone: o.BusinessUnitTimeZone,
		
		PublishedSchedules: o.PublishedSchedules,
		Alias:    (*Alias)(o),
	})
}

func (o *Buagentschedulessearchresponse) UnmarshalJSON(b []byte) error {
	var BuagentschedulessearchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentschedulessearchresponseMap)
	if err != nil {
		return err
	}
	
	if AgentSchedules, ok := BuagentschedulessearchresponseMap["agentSchedules"].([]interface{}); ok {
		AgentSchedulesString, _ := json.Marshal(AgentSchedules)
		json.Unmarshal(AgentSchedulesString, &o.AgentSchedules)
	}
	
	if BusinessUnitTimeZone, ok := BuagentschedulessearchresponseMap["businessUnitTimeZone"].(string); ok {
		o.BusinessUnitTimeZone = &BusinessUnitTimeZone
	}
    
	if PublishedSchedules, ok := BuagentschedulessearchresponseMap["publishedSchedules"].([]interface{}); ok {
		PublishedSchedulesString, _ := json.Marshal(PublishedSchedules)
		json.Unmarshal(PublishedSchedulesString, &o.PublishedSchedules)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentschedulessearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
