package platformclientv2
import (
	"encoding/json"
)

// Patchbureschedulingoptionsmanagementunitrequest
type Patchbureschedulingoptionsmanagementunitrequest struct { 
	// ManagementUnitId - The management unit portion of the rescheduling run to update
	ManagementUnitId *string `json:"managementUnitId,omitempty"`


	// Applied - Whether to mark the run as applied.  Only applies to reschedule runs.  Once applied, a run cannot be un-marked as applied
	Applied *bool `json:"applied,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchbureschedulingoptionsmanagementunitrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
