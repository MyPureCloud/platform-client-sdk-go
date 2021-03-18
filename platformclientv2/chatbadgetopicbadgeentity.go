package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Chatbadgetopicbadgeentity
type Chatbadgetopicbadgeentity struct { 
	// JabberId
	JabberId *string `json:"jabberId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Chatbadgetopicbadgeentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
