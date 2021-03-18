package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkrecordingenabledcount
type Trunkrecordingenabledcount struct { 
	// EnabledCount - The amount of trunks that have recording enabled
	EnabledCount *int `json:"enabledCount,omitempty"`


	// DisabledCount - The amount of trunks that do not have recording enabled
	DisabledCount *int `json:"disabledCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkrecordingenabledcount) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
