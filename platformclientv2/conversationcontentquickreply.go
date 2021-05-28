package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentquickreply - Quick reply object.
type Conversationcontentquickreply struct { 
	// Text - Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
	Text *string `json:"text,omitempty"`


	// Payload - Content of the payload included in the quick reply response. Could be an ID identifying the quick reply response.
	Payload *string `json:"payload,omitempty"`


	// Image - URL of an image associated with the quick reply.
	Image *string `json:"image,omitempty"`


	// Action - Specifies the type of action that is triggered upon clicking the quick reply.
	Action *string `json:"action,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentquickreply) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
