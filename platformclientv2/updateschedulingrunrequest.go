package platformclientv2
import (
	"encoding/json"
)

// Updateschedulingrunrequest
type Updateschedulingrunrequest struct { 
	// Applied - Mark the run as applied.  Request will be rejected if the value != true. Note: To discard a run without applying, you still need to mark it as applied so that other reschedule runs can be done
	Applied *bool `json:"applied,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updateschedulingrunrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
