package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationscreenshareeventtopicurireference
type Queueconversationscreenshareeventtopicurireference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicurireference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
