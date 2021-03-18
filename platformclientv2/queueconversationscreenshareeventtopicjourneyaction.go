package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationscreenshareeventtopicjourneyaction
type Queueconversationscreenshareeventtopicjourneyaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ActionMap
	ActionMap *Queueconversationscreenshareeventtopicjourneyactionmap `json:"actionMap,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicjourneyaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
