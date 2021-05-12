package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Buagentschedulesqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
