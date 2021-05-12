package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Maxparticipants
type Maxparticipants struct { 
	// MaxParticipants - The maximum number of participants that are allowed on a conversation.
	MaxParticipants *int `json:"maxParticipants,omitempty"`

}

// String returns a JSON representation of the model
func (o *Maxparticipants) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
