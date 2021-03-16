package platformclientv2
import (
	"encoding/json"
)

// Void
type Void struct { }

// String returns a JSON representation of the model
func (o *Void) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
