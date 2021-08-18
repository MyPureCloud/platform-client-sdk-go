package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Contentlist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentlist

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ListType *string `json:"listType,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SubmitLabel *string `json:"submitLabel,omitempty"`
		
		Actions *Contentactions `json:"actions,omitempty"`
		
		Components *[]Listitemcomponent `json:"components,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		ListType: u.ListType,
		
		Title: u.Title,
		
		Description: u.Description,
		
		SubmitLabel: u.SubmitLabel,
		
		Actions: u.Actions,
		
		Components: u.Components,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contentlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
