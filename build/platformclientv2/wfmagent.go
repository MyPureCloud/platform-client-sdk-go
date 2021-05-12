package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagent - Workforce management agent data
type Wfmagent struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// User - The user associated with this data
	User *Userreference `json:"user,omitempty"`


	// WorkPlan - The work plan associated with this agent, if applicable
	WorkPlan *Workplanreference `json:"workPlan,omitempty"`


	// WorkPlanRotation - The work plan rotation associated with this agent, if applicable
	WorkPlanRotation *Workplanrotationreference `json:"workPlanRotation,omitempty"`


	// AcceptDirectShiftTrades - Whether the agent accepts direct shift trade requests
	AcceptDirectShiftTrades *bool `json:"acceptDirectShiftTrades,omitempty"`


	// Queues - List of queues to which this agent is capable of handling
	Queues *[]Queuereference `json:"queues,omitempty"`


	// Languages - The list of languages this agent is capable of handling
	Languages *[]Languagereference `json:"languages,omitempty"`


	// Skills - The list of skills this agent is capable of handling
	Skills *[]Routingskillreference `json:"skills,omitempty"`


	// Schedulable - Whether the agent has the permission to be included in schedule generation
	Schedulable *bool `json:"schedulable,omitempty"`


	// Metadata - Metadata for this agent
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
