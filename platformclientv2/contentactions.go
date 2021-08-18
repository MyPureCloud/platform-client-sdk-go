package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentactions - User actions available on the content. All actions are optional and all actions are executed simultaneously.
type Contentactions struct { 
	// Url - A URL of a web page to direct the user to.
	Url *string `json:"url,omitempty"`


	// UrlTarget - The target window in which to open the URL. If empty will open a blank page or tab.
	UrlTarget *string `json:"urlTarget,omitempty"`


	// Textback - Text to be sent back in reply when the item is selected.
	Textback *string `json:"textback,omitempty"`

}

func (u *Contentactions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentactions

	

	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		UrlTarget *string `json:"urlTarget,omitempty"`
		
		Textback *string `json:"textback,omitempty"`
		*Alias
	}{ 
		Url: u.Url,
		
		UrlTarget: u.UrlTarget,
		
		Textback: u.Textback,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contentactions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
