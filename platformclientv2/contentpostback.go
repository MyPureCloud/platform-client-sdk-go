package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contentpostback - Postback response object representing the click of a rich media button (Deprecated).
type Contentpostback struct { 
	// Id - An ID assigned to the button response.
	Id *string `json:"id,omitempty"`


	// Text - The response text from the button click.
	Text *string `json:"text,omitempty"`


	// Payload - The response payload associated with the clicked button.
	Payload *string `json:"payload,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentpostback) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
