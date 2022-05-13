package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchcalltoaction
type Patchcalltoaction struct { 
	// Text - Text displayed on the call to action button.
	Text *string `json:"text,omitempty"`


	// Url - URL to open when user clicks on the call to action button.
	Url *string `json:"url,omitempty"`


	// Target - Where the URL should be opened when the user clicks on the call to action button.
	Target *string `json:"target,omitempty"`

}

func (o *Patchcalltoaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchcalltoaction
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Target *string `json:"target,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Url: o.Url,
		
		Target: o.Target,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchcalltoaction) UnmarshalJSON(b []byte) error {
	var PatchcalltoactionMap map[string]interface{}
	err := json.Unmarshal(b, &PatchcalltoactionMap)
	if err != nil {
		return err
	}
	
	if Text, ok := PatchcalltoactionMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Url, ok := PatchcalltoactionMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Target, ok := PatchcalltoactionMap["target"].(string); ok {
		o.Target = &Target
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Patchcalltoaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
