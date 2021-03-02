package platformclientv2
import (
	"encoding/json"
)

// Userconversationseventmediasummarydetail
type Userconversationseventmediasummarydetail struct { 
	// Active
	Active *int `json:"active,omitempty"`


	// Acw
	Acw *int `json:"acw,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userconversationseventmediasummarydetail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
