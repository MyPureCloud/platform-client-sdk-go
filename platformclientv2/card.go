package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Card
type Card struct { 
	// Title - Text to show in the title.
	Title *string `json:"title,omitempty"`


	// Description - Text to show in the description.
	Description *string `json:"description,omitempty"`


	// Url - URL of an image.
	Url *string `json:"url,omitempty"`


	// DefaultAction - The default action to be taken.
	DefaultAction *Cardaction `json:"defaultAction,omitempty"`


	// Actions - List of possible action objects.
	Actions *[]Cardaction `json:"actions,omitempty"`

}

func (o *Card) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Card
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		DefaultAction *Cardaction `json:"defaultAction,omitempty"`
		
		Actions *[]Cardaction `json:"actions,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Description: o.Description,
		
		Url: o.Url,
		
		DefaultAction: o.DefaultAction,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Card) UnmarshalJSON(b []byte) error {
	var CardMap map[string]interface{}
	err := json.Unmarshal(b, &CardMap)
	if err != nil {
		return err
	}
	
	if Title, ok := CardMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := CardMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Url, ok := CardMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if DefaultAction, ok := CardMap["defaultAction"].(map[string]interface{}); ok {
		DefaultActionString, _ := json.Marshal(DefaultAction)
		json.Unmarshal(DefaultActionString, &o.DefaultAction)
	}
	
	if Actions, ok := CardMap["actions"].([]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Card) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
