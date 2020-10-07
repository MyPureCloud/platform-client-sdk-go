package platformclientv2
import (
	"encoding/json"
)

// Constraintconflictmessage
type Constraintconflictmessage struct { 
	// Message - Message for how to resolve a set of conflicted work plan constraints
	Message *Workplanconstraintconflictmessage `json:"message,omitempty"`


	// ConflictedConstraintMessages - Messages for the set of conflicted work plan constraints. Each element indicates the message of a work plan constraint that is conflicted in the set
	ConflictedConstraintMessages *[]Workplanconstraintmessage `json:"conflictedConstraintMessages,omitempty"`

}

// String returns a JSON representation of the model
func (o *Constraintconflictmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
