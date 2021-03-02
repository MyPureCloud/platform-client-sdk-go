package platformclientv2
import (
	"encoding/json"
)

// Wfmupdateagentdetailstopicwfmupdateagentdetailscomplete
type Wfmupdateagentdetailstopicwfmupdateagentdetailscomplete struct { 
	// Status
	Status *string `json:"status,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmupdateagentdetailstopicwfmupdateagentdetailscomplete) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
