package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Edgemetricstopicedgemetricsubsystem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricstopicedgemetricsubsystem

	

	return json.Marshal(&struct { 
		ProcessName *string `json:"processName,omitempty"`
		
		DelayMs *int `json:"delayMs,omitempty"`
		
		MediaSubsystem *Edgemetricstopicedgemetricsubsystemmedia `json:"mediaSubsystem,omitempty"`
		*Alias
	}{ 
		ProcessName: u.ProcessName,
		
		DelayMs: u.DelayMs,
		
		MediaSubsystem: u.MediaSubsystem,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricsubsystem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
