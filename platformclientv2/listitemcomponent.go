package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Listitemcomponent - An entry in List template
type Listitemcomponent struct { 
	// Id - An ID assigned to this component
	Id *string `json:"id,omitempty"`


	// Rmid - An ID of the rich message instance
	Rmid *string `json:"rmid,omitempty"`


	// VarType - The type of component to render
	VarType *string `json:"type,omitempty"`


	// Image - Path or URI to an image file
	Image *string `json:"image,omitempty"`


	// Title - The main headline of the list item
	Title *string `json:"title,omitempty"`


	// Description - Text to show in the description row
	Description *string `json:"description,omitempty"`


	// Actions - User actions available on the content. All actions are optional and all actions are executed simultaneously.
	Actions *Contentactions `json:"actions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Listitemcomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
