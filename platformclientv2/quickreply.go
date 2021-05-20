package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Quickreply
type Quickreply struct { 
	// Id - An ID assigned to the quick reply.
	Id *string `json:"id,omitempty"`


	// Text - Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
	Text *string `json:"text,omitempty"`


	// Payload - Content of the textback payload after clicking a quick reply
	Payload *string `json:"payload,omitempty"`


	// Url - The location of the image file associated with quick reply
	Url *string `json:"url,omitempty"`


	// Action - Specifies the type of action that is triggered upon clicking the quick reply. Currently, the only supported action is \"Message\" which sends a message using the quick reply text.
	Action *string `json:"action,omitempty"`


	// IsSelected - Indicates if the quick reply option is selected by end customer
	IsSelected *bool `json:"isSelected,omitempty"`

}

// String returns a JSON representation of the model
func (o *Quickreply) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
