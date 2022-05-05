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


	// Actions - The list actions (Deprecated).
	Actions *Contentactions `json:"actions,omitempty"`


	// Components - An array of component objects.
	Components *[]Listitemcomponent `json:"components,omitempty"`

}

func (o *Contentlist) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		ListType: o.ListType,
		
		Title: o.Title,
		
		Description: o.Description,
		
		SubmitLabel: o.SubmitLabel,
		
		Actions: o.Actions,
		
		Components: o.Components,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentlist) UnmarshalJSON(b []byte) error {
	var ContentlistMap map[string]interface{}
	err := json.Unmarshal(b, &ContentlistMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContentlistMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ListType, ok := ContentlistMap["listType"].(string); ok {
		o.ListType = &ListType
	}
	
	if Title, ok := ContentlistMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Description, ok := ContentlistMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if SubmitLabel, ok := ContentlistMap["submitLabel"].(string); ok {
		o.SubmitLabel = &SubmitLabel
	}
	
	if Actions, ok := ContentlistMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Components, ok := ContentlistMap["components"].([]interface{}); ok {
		ComponentsString, _ := json.Marshal(Components)
		json.Unmarshal(ComponentsString, &o.Components)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
