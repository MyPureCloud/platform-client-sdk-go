package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Edgemetricssubsystem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
