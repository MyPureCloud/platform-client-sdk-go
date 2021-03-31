package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contentquickreply - Quick reply object
type Contentquickreply struct { 
	// Id - An ID assigned to the quick reply. Each object inside the content array has a unique ID.
	Id *string `json:"id,omitempty"`


	// Text - Text to show inside the quick reply. This is also used as the response text after clicking on the quick reply.
	Text *string `json:"text,omitempty"`


	// Payload - Content of the textback payload after clicking a quick reply
	Payload *string `json:"payload,omitempty"`


	// Image - Path or URI to an image file associated with quick reply
	Image *string `json:"image,omitempty"`


	// Action - Specifies the type of action that is triggered upon clicking the quick reply. Currently, the only supported action is \"Message\" which sends a message using the quick reply text.
	Action *string `json:"action,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentquickreply) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
