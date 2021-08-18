package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Recordingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingsettings

	

	return json.Marshal(&struct { 
		MaxSimultaneousStreams *int `json:"maxSimultaneousStreams,omitempty"`
		
		MaxConfigurableScreenRecordingStreams *int `json:"maxConfigurableScreenRecordingStreams,omitempty"`
		*Alias
	}{ 
		MaxSimultaneousStreams: u.MaxSimultaneousStreams,
		
		MaxConfigurableScreenRecordingStreams: u.MaxConfigurableScreenRecordingStreams,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Recordingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
