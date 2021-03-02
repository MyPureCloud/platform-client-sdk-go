package platformclientv2
import (
	"encoding/json"
)

// Patchbureschedulingoptionsrequest
type Patchbureschedulingoptionsrequest struct { 
	// ManagementUnits - Per-management unit rescheduling options to update
	ManagementUnits *[]Patchbureschedulingoptionsmanagementunitrequest `json:"managementUnits,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchbureschedulingoptionsrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
