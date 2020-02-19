package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmbuscheduletopicbumanagementunitschedulesummary
type Wfmbuscheduletopicbumanagementunitschedulesummary struct { 
	// ManagementUnit
	ManagementUnit *Wfmbuscheduletopicmanagementunit `json:"managementUnit,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`


	// Agents
	Agents *[]Wfmbuscheduletopicuserreference `json:"agents,omitempty"`


	// AgentCount
	AgentCount *int32 `json:"agentCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbumanagementunitschedulesummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
