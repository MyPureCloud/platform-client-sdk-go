package platformclientv2
import (
	"encoding/json"
)

// Architectflowfields
type Architectflowfields struct { 
	// ArchitectFlow - The architect flow.
	ArchitectFlow *Addressableentityref `json:"architectFlow,omitempty"`


	// FlowRequestMappings - Collection of Architect Flow Request Mappings to use.
	FlowRequestMappings *[]Requestmapping `json:"flowRequestMappings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflowfields) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
