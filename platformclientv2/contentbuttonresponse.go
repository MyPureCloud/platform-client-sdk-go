package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contentbuttonresponse - Button response object representing the click of a structured message button, such as a quick reply.
type Contentbuttonresponse struct { 
	// Id - An ID assigned to the button response (Deprecated).
	Id *string `json:"id,omitempty"`


	// VarType - Describes the button that resulted in the Button Response.
	VarType *string `json:"type,omitempty"`


	// Text - The response text from the button click.
	Text *string `json:"text,omitempty"`


	// Payload - The response payload associated with the clicked button.
	Payload *string `json:"payload,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentbuttonresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
