package platformclientv2
import (
	"encoding/json"
)

// Connectrate
type Connectrate struct { 
	// Attempts - Number of call attempts made
	Attempts *int64 `json:"attempts,omitempty"`


	// Connects - Number of calls with a live voice detected
	Connects *int64 `json:"connects,omitempty"`


	// ConnectRatio - Ratio of connects to attempts
	ConnectRatio *float64 `json:"connectRatio,omitempty"`

}

// String returns a JSON representation of the model
func (o *Connectrate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
