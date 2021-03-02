package platformclientv2
import (
	"encoding/json"
)

// Patchbuschedulerunrequest
type Patchbuschedulerunrequest struct { 
	// ReschedulingOptions - The rescheduling options to update
	ReschedulingOptions *Patchbureschedulingoptionsrequest `json:"reschedulingOptions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchbuschedulerunrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
