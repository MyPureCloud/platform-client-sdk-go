package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationscreenshareeventtopicurireference
type Conversationscreenshareeventtopicurireference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicurireference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
