package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webmessaginggeneric - Generic content object. (Deprecated
type Webmessaginggeneric struct { 
	// Title - Text to show in the title.
	Title *string `json:"title,omitempty"`


	// Description - Text to show in the description.
	Description *string `json:"description,omitempty"`


	// Image - URL of an image.
	Image *string `json:"image,omitempty"`


	// Actions - Actions to be taken.
	Actions *Contentactions `json:"actions,omitempty"`


	// Components - An array of component objects.
	Components *[]Buttoncomponent `json:"components,omitempty"`

}

func (o *Webmessaginggeneric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webmessaginggeneric
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Actions *Contentactions `json:"actions,omitempty"`
		
		Components *[]Buttoncomponent `json:"components,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Description: o.Description,
		
		Image: o.Image,
		
		Actions: o.Actions,
		
		Components: o.Components,
		Alias:    (*Alias)(o),
	})
}

func (o *Webmessaginggeneric) UnmarshalJSON(b []byte) error {
	var WebmessaginggenericMap map[string]interface{}
	err := json.Unmarshal(b, &WebmessaginggenericMap)
	if err != nil {
		return err
	}
	
	if Title, ok := WebmessaginggenericMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := WebmessaginggenericMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Image, ok := WebmessaginggenericMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if Actions, ok := WebmessaginggenericMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Components, ok := WebmessaginggenericMap["components"].([]interface{}); ok {
		ComponentsString, _ := json.Marshal(Components)
		json.Unmarshal(ComponentsString, &o.Components)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webmessaginggeneric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
