package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingsettings
type Recordingsettings struct { 
	// MaxSimultaneousStreams - Maximum number of simultaneous screen recording streams
	MaxSimultaneousStreams *int `json:"maxSimultaneousStreams,omitempty"`


	// MaxConfigurableScreenRecordingStreams - Upper limit that maxSimultaneousStreams can be configured
	MaxConfigurableScreenRecordingStreams *int `json:"maxConfigurableScreenRecordingStreams,omitempty"`

}

// String returns a JSON representation of the model
func (o *Recordingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
