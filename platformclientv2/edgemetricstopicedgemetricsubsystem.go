package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricstopicedgemetricsubsystem
type Edgemetricstopicedgemetricsubsystem struct { 
	// ProcessName
	ProcessName *string `json:"processName,omitempty"`


	// DelayMs
	DelayMs *int `json:"delayMs,omitempty"`


	// MediaSubsystem
	MediaSubsystem *Edgemetricstopicedgemetricsubsystemmedia `json:"mediaSubsystem,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricsubsystem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
