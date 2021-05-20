package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Buttonresponse
type Buttonresponse struct { 
	// Id - An ID assigned to the button response.
	Id *string `json:"id,omitempty"`


	// VarType - Button response type that captures Button and QuickReply type responses
	VarType *string `json:"type,omitempty"`


	// Text - Text to show inside the Button reply. This is also used as the response text after clicking on the Button.
	Text *string `json:"text,omitempty"`


	// Payload - Content of the textback payload after clicking a button
	Payload *string `json:"payload,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buttonresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
