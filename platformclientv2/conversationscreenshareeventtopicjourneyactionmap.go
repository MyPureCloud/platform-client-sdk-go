package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationscreenshareeventtopicjourneyactionmap
type Conversationscreenshareeventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
