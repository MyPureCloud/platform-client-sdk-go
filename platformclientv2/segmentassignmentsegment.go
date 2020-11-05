package platformclientv2
import (
	"encoding/json"
)

// Segmentassignmentsegment
type Segmentassignmentsegment struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Scope - The target entity that a segment applies to.
	Scope *string `json:"scope,omitempty"`


	// Version - The version of the segment.
	Version *int32 `json:"version,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Segmentassignmentsegment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
