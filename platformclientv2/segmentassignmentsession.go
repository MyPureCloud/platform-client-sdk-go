package platformclientv2
import (
	"encoding/json"
)

// Segmentassignmentsession
type Segmentassignmentsession struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// VarType - The type or category of session (e.g. web, ticket, delivery, atm).
	VarType *string `json:"type,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Segmentassignmentsession) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
