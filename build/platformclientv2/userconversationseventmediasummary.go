package platformclientv2
import (
	"encoding/json"
)

// Userconversationseventmediasummary
type Userconversationseventmediasummary struct { 
	// ContactCenter
	ContactCenter *Userconversationseventmediasummarydetail `json:"contactCenter,omitempty"`


	// Enterprise
	Enterprise *Userconversationseventmediasummarydetail `json:"enterprise,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userconversationseventmediasummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
