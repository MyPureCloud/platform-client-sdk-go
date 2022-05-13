package platformclientv2
import (
	"github.com/leekchan/timeutil"
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


	// Actions - The list item actions (Deprecated).
	Actions *Contentactions `json:"actions,omitempty"`

}

func (o *Listitemcomponent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Listitemcomponent
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Rmid *string `json:"rmid,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Actions *Contentactions `json:"actions,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Rmid: o.Rmid,
		
		VarType: o.VarType,
		
		Image: o.Image,
		
		Title: o.Title,
		
		Description: o.Description,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Listitemcomponent) UnmarshalJSON(b []byte) error {
	var ListitemcomponentMap map[string]interface{}
	err := json.Unmarshal(b, &ListitemcomponentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ListitemcomponentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Rmid, ok := ListitemcomponentMap["rmid"].(string); ok {
		o.Rmid = &Rmid
	}
    
	if VarType, ok := ListitemcomponentMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Image, ok := ListitemcomponentMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if Title, ok := ListitemcomponentMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := ListitemcomponentMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Actions, ok := ListitemcomponentMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Listitemcomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
