package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Recordingjobfailedrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
