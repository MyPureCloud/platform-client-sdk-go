package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricssubsystem
type Edgemetricssubsystem struct { 
	// DelayMs - Delay in milliseconds.
	DelayMs *int `json:"delayMs,omitempty"`


	// ProcessName - Name of the Edge process.
	ProcessName *string `json:"processName,omitempty"`


	// MediaSubsystem - Subsystem for an Edge device.
	MediaSubsystem **Edgemetricssubsystem `json:"mediaSubsystem,omitempty"`

}

func (u *Edgemetricssubsystem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricssubsystem

	

	return json.Marshal(&struct { 
		DelayMs *int `json:"delayMs,omitempty"`
		
		ProcessName *string `json:"processName,omitempty"`
		
		MediaSubsystem **Edgemetricssubsystem `json:"mediaSubsystem,omitempty"`
		*Alias
	}{ 
		DelayMs: u.DelayMs,
		
		ProcessName: u.ProcessName,
		
		MediaSubsystem: u.MediaSubsystem,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgemetricssubsystem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
