package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Generictemplate
type Generictemplate struct { 
	// Title - Text to show in the title.
	Title *string `json:"title,omitempty"`


	// Description - Text to show in the description.
	Description *string `json:"description,omitempty"`


	// Url - URL of an image.
	Url *string `json:"url,omitempty"`


	// Components - List of button components offered with this message content.
	Components *[]Recordingbuttoncomponent `json:"components,omitempty"`


	// Actions - Actions to be taken.
	Actions *Recordingcontentactions `json:"actions,omitempty"`

}

func (o *Generictemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Generictemplate
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Components *[]Recordingbuttoncomponent `json:"components,omitempty"`
		
		Actions *Recordingcontentactions `json:"actions,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Description: o.Description,
		
		Url: o.Url,
		
		Components: o.Components,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Generictemplate) UnmarshalJSON(b []byte) error {
	var GenerictemplateMap map[string]interface{}
	err := json.Unmarshal(b, &GenerictemplateMap)
	if err != nil {
		return err
	}
	
	if Title, ok := GenerictemplateMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Description, ok := GenerictemplateMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Url, ok := GenerictemplateMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if Components, ok := GenerictemplateMap["components"].([]interface{}); ok {
		ComponentsString, _ := json.Marshal(Components)
		json.Unmarshal(ComponentsString, &o.Components)
	}
	
	if Actions, ok := GenerictemplateMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Generictemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
