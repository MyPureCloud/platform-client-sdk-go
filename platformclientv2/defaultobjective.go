package platformclientv2
import (
	"encoding/json"
)

// Defaultobjective
type Defaultobjective struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// TemplateId - The id of this objective's base template
	TemplateId *string `json:"templateId,omitempty"`


	// Zones - Objective zone specifies min,max points and values for the associated metric
	Zones *[]Objectivezone `json:"zones,omitempty"`


	// Enabled - A flag for whether this objective is enabled for the related metric
	Enabled *bool `json:"enabled,omitempty"`

}

// String returns a JSON representation of the model
func (o *Defaultobjective) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
