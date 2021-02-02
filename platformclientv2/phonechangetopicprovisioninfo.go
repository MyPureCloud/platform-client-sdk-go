package platformclientv2
import (
	"time"
	"encoding/json"
)

// Phonechangetopicprovisioninfo
type Phonechangetopicprovisioninfo struct { 
	// Time
	Time *time.Time `json:"time,omitempty"`


	// Source
	Source *string `json:"source,omitempty"`


	// ErrorInfo
	ErrorInfo *string `json:"errorInfo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonechangetopicprovisioninfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
