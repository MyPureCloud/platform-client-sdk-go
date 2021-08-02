package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Buttoncomponent - Structured template button object.
type Buttoncomponent struct { 
	// Title - Text to show inside the button.
	Title *string `json:"title,omitempty"`


	// Actions - The button actions.
	Actions *Contentactions `json:"actions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buttoncomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
