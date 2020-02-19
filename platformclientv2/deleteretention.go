package platformclientv2
import (
	"encoding/json"
)

// Deleteretention
type Deleteretention struct { 
	// Days
	Days *int32 `json:"days,omitempty"`

}

// String returns a JSON representation of the model
func (o *Deleteretention) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
