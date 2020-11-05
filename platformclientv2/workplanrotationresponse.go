package platformclientv2
import (
	"encoding/json"
)

// Workplanrotationresponse
type Workplanrotationresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Enabled - Whether the work plan rotation is enabled for scheduling
	Enabled *bool `json:"enabled,omitempty"`


	// DateRange - The date range to which this work plan rotation applies
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`


	// Pattern - Pattern with ordered list of work plans that rotate on a weekly basis
	Pattern *Workplanpatternresponse `json:"pattern,omitempty"`


	// AgentCount - Number of agents in this work plan rotation
	AgentCount *int32 `json:"agentCount,omitempty"`


	// Agents - Agents in this work plan rotation. Populate with expand=agents for GET WorkPlanRotationsList (defaults to empty list)
	Agents *[]Workplanrotationagentresponse `json:"agents,omitempty"`


	// Metadata - Version metadata for this work plan rotation
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanrotationresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
