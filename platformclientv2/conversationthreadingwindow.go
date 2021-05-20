package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationthreadingwindow
type Conversationthreadingwindow struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Settings - The conversation threading window timeout (Minutes) for each messaging type
	Settings *[]Conversationthreadingwindowsetting `json:"settings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationthreadingwindow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
