package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Articlecontent
type Articlecontent struct { 
	// Body - Body of the article content.
	Body *Articlecontentbody `json:"body,omitempty"`

}

// String returns a JSON representation of the model
func (o *Articlecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
