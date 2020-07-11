package platformclientv2
import (
	"encoding/json"
)

// Buttoncomponent - Structured template button object
type Buttoncomponent struct { 
	// Id - An ID assigned to this component
	Id *string `json:"id,omitempty"`


	// Text - Text to show inside the button
	Text *string `json:"text,omitempty"`


	// Title - Tooltip text to show when hovering over the button
	Title *string `json:"title,omitempty"`


	// Actions - User actions available on the content. All actions are optional and all actions are executed simultaneously.
	Actions *Contentactions `json:"actions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buttoncomponent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
