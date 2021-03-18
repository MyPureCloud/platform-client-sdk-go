package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Connectrate
type Connectrate struct { 
	// Attempts - Number of call attempts made
	Attempts *int `json:"attempts,omitempty"`


	// Connects - Number of calls with a live voice detected
	Connects *int `json:"connects,omitempty"`


	// ConnectRatio - Ratio of connects to attempts
	ConnectRatio *float64 `json:"connectRatio,omitempty"`

}

// String returns a JSON representation of the model
func (o *Connectrate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
