package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
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
	AgentCount *int `json:"agentCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbumanagementunitschedulesummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
