package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingjobfailedrecording
type Recordingjobfailedrecording struct { 
	// Conversation - Conversation
	Conversation *Addressableentityref `json:"conversation,omitempty"`


	// Recording - Recording
	Recording *Addressableentityref `json:"recording,omitempty"`

}

func (u *Recordingjobfailedrecording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingjobfailedrecording

	

	return json.Marshal(&struct { 
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		
		Recording *Addressableentityref `json:"recording,omitempty"`
		*Alias
	}{ 
		Conversation: u.Conversation,
		
		Recording: u.Recording,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Recordingjobfailedrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
