package platformclientv2
import (
	"encoding/json"
)

// Empty
type Empty struct { }

// String returns a JSON representation of the model
func (o *Empty) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
