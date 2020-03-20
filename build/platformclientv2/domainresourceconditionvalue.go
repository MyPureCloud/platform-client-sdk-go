package platformclientv2
import (
	"encoding/json"
)

// Domainresourceconditionvalue
type Domainresourceconditionvalue struct { 
	// User
	User *User `json:"user,omitempty"`


	// Queue
	Queue *Queue `json:"queue,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Domainresourceconditionvalue) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
