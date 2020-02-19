package platformclientv2
import (
	"encoding/json"
)

// Wfmmovemanagementunittopicmovemanagementunitnotification
type Wfmmovemanagementunittopicmovemanagementunitnotification struct { 
	// BusinessUnit
	BusinessUnit *Wfmmovemanagementunittopicbusinessunit `json:"businessUnit,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmmovemanagementunittopicmovemanagementunitnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
