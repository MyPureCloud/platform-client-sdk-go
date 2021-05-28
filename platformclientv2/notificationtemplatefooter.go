package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Notificationtemplatefooter - Template footer object.
type Notificationtemplatefooter struct { 
	// Text - Footer text. For WhatsApp, ignored.
	Text *string `json:"text,omitempty"`

}

// String returns a JSON representation of the model
func (o *Notificationtemplatefooter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
