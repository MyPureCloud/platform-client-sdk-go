package platformclientv2
import (
	"encoding/json"
)

// Contentpostback - The postback object result of a user clicking in a button
type Contentpostback struct { 
	// Id - An ID assigned to the postback reply. Each object inside the content array has a unique ID.
	Id *string `json:"id,omitempty"`


	// Text - The text inside the button clicked (in the structured message template)
	Text *string `json:"text,omitempty"`


	// Payload - Content of the textback payload after clicking a quick reply
	Payload *string `json:"payload,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentpostback) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
