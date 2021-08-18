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


	// Actions - The button actions.
	Actions *Contentactions `json:"actions,omitempty"`

}

func (u *Buttoncomponent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buttoncomponent

	

	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Actions *Contentactions `json:"actions,omitempty"`
		*Alias
	}{ 
		Title: u.Title,
		
		Actions: u.Actions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buttoncomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
