package platformclientv2
import (
	"encoding/json"
)

// Flowdiagnosticinfo
type Flowdiagnosticinfo struct { 
	// LastActionId - The step number of the survey invite flow where the error occurred.
	LastActionId *int `json:"lastActionId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowdiagnosticinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
