package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buttoncomponent - Structured template button object.
type Buttoncomponent struct { 
	// Title - Text to show inside the button.
	Title *string `json:"title,omitempty"`


	// Actions - The button actions (Deprecated).
	Actions *Contentactions `json:"actions,omitempty"`

}

func (o *Buttoncomponent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buttoncomponent
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Actions *Contentactions `json:"actions,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Buttoncomponent) UnmarshalJSON(b []byte) error {
	var ButtoncomponentMap map[string]interface{}
	err := json.Unmarshal(b, &ButtoncomponentMap)
	if err != nil {
		return err
	}
	
	if Title, ok := ButtoncomponentMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Actions, ok := ButtoncomponentMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buttoncomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
