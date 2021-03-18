package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationemaileventtopicjourneyaction
type Queueconversationemaileventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationemaileventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
