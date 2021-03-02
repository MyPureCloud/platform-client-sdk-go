package platformclientv2
import (
	"encoding/json"
)

// Edgemetricstopicedgemetricsubsystemmedia
type Edgemetricstopicedgemetricsubsystemmedia struct { 
	// ProcessName
	ProcessName *string `json:"processName,omitempty"`


	// DelayMs
	DelayMs *int `json:"delayMs,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricsubsystemmedia) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
