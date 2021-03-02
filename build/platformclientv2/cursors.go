package platformclientv2
import (
	"encoding/json"
)

// Cursors
type Cursors struct { 
	// Before
	Before *string `json:"before,omitempty"`


	// After
	After *string `json:"after,omitempty"`

}

// String returns a JSON representation of the model
func (o *Cursors) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
