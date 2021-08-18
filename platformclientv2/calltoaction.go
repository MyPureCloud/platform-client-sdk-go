package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Calltoaction
type Calltoaction struct { 
	// Text - Text displayed on the call to action button.
	Text *string `json:"text,omitempty"`


	// Url - URL to open when user clicks on the call to action button.
	Url *string `json:"url,omitempty"`


	// Target - Where the URL should be opened when the user clicks on the call to action button.
	Target *string `json:"target,omitempty"`

}

func (u *Calltoaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Calltoaction

	

	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Target *string `json:"target,omitempty"`
		*Alias
	}{ 
		Text: u.Text,
		
		Url: u.Url,
		
		Target: u.Target,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Calltoaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
