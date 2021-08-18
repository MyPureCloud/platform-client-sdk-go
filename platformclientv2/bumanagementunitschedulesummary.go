package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bumanagementunitschedulesummary
type Bumanagementunitschedulesummary struct { 
	// ManagementUnit - The management unit to which this summary applies
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`


	// AgentCount - The number of agents from this management unit that are in the schedule
	AgentCount *int `json:"agentCount,omitempty"`


	// StartDate - The start of the schedule change in the management unit. Only populated in schedule update notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end of the schedule change in the management unit. Only populated in schedule update notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// Agents - The agents in the management unit who are part of this schedule, or in schedule change notifications, the agents that were changed. Note this will come back as an empty list unless the appropriate expand query parameter is passed
	Agents *[]Userreference `json:"agents,omitempty"`

}

func (u *Bumanagementunitschedulesummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bumanagementunitschedulesummary

	
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
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		AgentCount *int `json:"agentCount,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		Agents *[]Userreference `json:"agents,omitempty"`
		*Alias
	}{ 
		ManagementUnit: u.ManagementUnit,
		
		AgentCount: u.AgentCount,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Agents: u.Agents,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bumanagementunitschedulesummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
