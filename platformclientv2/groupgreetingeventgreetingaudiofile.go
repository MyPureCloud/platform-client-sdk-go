package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Groupgreetingeventgreetingaudiofile
type Groupgreetingeventgreetingaudiofile struct { 
	// DurationMilliseconds
	DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`


	// SizeBytes
	SizeBytes *int `json:"sizeBytes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Groupgreetingeventgreetingaudiofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
