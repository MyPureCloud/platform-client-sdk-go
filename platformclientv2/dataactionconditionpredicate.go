package platformclientv2
import (
	"encoding/json"
)

// Dataactionconditionpredicate
type Dataactionconditionpredicate struct { }

// String returns a JSON representation of the model
func (o *Dataactionconditionpredicate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
