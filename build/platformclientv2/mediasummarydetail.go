package platformclientv2
import (
	"encoding/json"
)

// Mediasummarydetail
type Mediasummarydetail struct { 
	// Active
	Active *int `json:"active,omitempty"`


	// Acw
	Acw *int `json:"acw,omitempty"`

}

// String returns a JSON representation of the model
func (o *Mediasummarydetail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
