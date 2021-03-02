package platformclientv2
import (
	"encoding/json"
)

// Patchaction
type Patchaction struct { 
	// MediaType - Media type of action.
	MediaType *string `json:"mediaType,omitempty"`


	// ActionTemplate - Action template associated with the action map.
	ActionTemplate *Actionmapactiontemplate `json:"actionTemplate,omitempty"`


	// ArchitectFlowFields - Architect Flow Id and input contract.
	ArchitectFlowFields *Architectflowfields `json:"architectFlowFields,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
