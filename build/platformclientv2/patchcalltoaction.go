package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Patchcalltoaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
