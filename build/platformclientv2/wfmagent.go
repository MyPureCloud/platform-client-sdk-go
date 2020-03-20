package platformclientv2
import (
	"encoding/json"
)

// Wfmagent - Workforce management agent data
type Wfmagent struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// User - The user associated with this data
	User *Userreference `json:"user,omitempty"`


	// WorkPlan - The work plan associated with this agent
	WorkPlan *Workplanreference `json:"workPlan,omitempty"`


	// TimeZone - The time zone for this agent. Defaults to the time zone of the management unit to which the agent belongs
	TimeZone *Wfmtimezone `json:"timeZone,omitempty"`


	// AcceptDirectShiftTrades - Whether the agent accepts direct shift trade requests
	AcceptDirectShiftTrades *bool `json:"acceptDirectShiftTrades,omitempty"`


	// Metadata - Metadata for this agent
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// Queues - List of queues to which the agent belongs and which are defined in the service goal groups in this management unit
	Queues *[]Queuereference `json:"queues,omitempty"`


	// Languages - The list of languages
	Languages *[]Languagereference `json:"languages,omitempty"`


	// Skills - The list of skills
	Skills *[]Routingskillreference `json:"skills,omitempty"`


	// Schedulable - Whether the agent has the permission to be included in schedule generation
	Schedulable *bool `json:"schedulable,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
