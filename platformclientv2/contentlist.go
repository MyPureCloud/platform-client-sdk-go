package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contentlist - List content object.
type Contentlist struct { 
	// Id - A unique ID assigned to this rich message content.
	Id *string `json:"id,omitempty"`


	// ListType - The type of list this instance represents.
	ListType *string `json:"listType,omitempty"`


	// Title - Text to show in the title.
	Title *string `json:"title,omitempty"`


	// Description - Text to show in the description.
	Description *string `json:"description,omitempty"`


	// SubmitLabel - Label for Submit button.
	SubmitLabel *string `json:"submitLabel,omitempty"`


	// Actions - The list actions.
	Actions *Contentactions `json:"actions,omitempty"`


	// Components - An array of component objects.
	Components *[]Listitemcomponent `json:"components,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
