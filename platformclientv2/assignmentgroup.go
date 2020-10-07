package platformclientv2
import (
	"encoding/json"
)

// Assignmentgroup
type Assignmentgroup struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Assignmentgroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
