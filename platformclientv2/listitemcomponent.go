package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Listitemcomponent - An entry in a List template.
type Listitemcomponent struct { 
	// Id - An ID assigned to this list item.
	Id *string `json:"id,omitempty"`


	// Rmid - An ID of the rich message instance.
	Rmid *string `json:"rmid,omitempty"`


	// VarType - The type of list item to render.
	VarType *string `json:"type,omitempty"`


	// Image - URL of an image.
	Image *string `json:"image,omitempty"`


	// Title - The main headline of the list item.
	Title *string `json:"title,omitempty"`


	// Description - Text to show in the list item description.
	Description *string `json:"description,omitempty"`


	// Actions - The list item actions.
	Actions *Contentactions `json:"actions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Listitemcomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
