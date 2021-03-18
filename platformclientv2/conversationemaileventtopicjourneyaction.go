package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationemaileventtopicjourneyaction
type Conversationemaileventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Conversationemaileventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationemaileventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
