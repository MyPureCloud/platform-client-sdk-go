package platformclientv2
import (
	"time"
	"encoding/json"
)

// Bumanagementunitschedulesummary
type Bumanagementunitschedulesummary struct { 
	// ManagementUnit - The management unit to which this summary applies
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`


	// AgentCount - The number of agents from this management unit that are in the schedule
	AgentCount *int32 `json:"agentCount,omitempty"`


	// StartDate - The start of the schedule change in the management unit. Only populated in schedule update notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The end of the schedule change in the management unit. Only populated in schedule update notifications. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndDate *time.Time `json:"endDate,omitempty"`


	// Agents - The changed agents in the management unit. Only populated in schedule update notifications
	Agents *[]Userreference `json:"agents,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bumanagementunitschedulesummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
