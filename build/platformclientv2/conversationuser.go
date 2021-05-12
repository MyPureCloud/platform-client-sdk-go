package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationuser
type Conversationuser struct { 
	// Id - The globally unique identifier for this user.
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
