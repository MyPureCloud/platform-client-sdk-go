package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulesqueryresponse
type Buagentschedulesqueryresponse struct { 
	// AgentSchedules - The requested agent schedules
	AgentSchedules *[]Buagentschedulequeryresponse `json:"agentSchedules,omitempty"`


	// BusinessUnitTimeZone - The time zone configured for the business unit to which these schedules apply
	BusinessUnitTimeZone *string `json:"businessUnitTimeZone,omitempty"`

}

func (u *Buagentschedulesqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulesqueryresponse

	

	return json.Marshal(&struct { 
		AgentSchedules *[]Buagentschedulequeryresponse `json:"agentSchedules,omitempty"`
		
		BusinessUnitTimeZone *string `json:"businessUnitTimeZone,omitempty"`
		*Alias
	}{ 
		AgentSchedules: u.AgentSchedules,
		
		BusinessUnitTimeZone: u.BusinessUnitTimeZone,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buagentschedulesqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
