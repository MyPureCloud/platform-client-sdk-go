package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Helplink - Link to a help or support resource
type Helplink struct { 
	// Uri - URI of the help resource
	Uri *string `json:"uri,omitempty"`


	// Title - Link text of the resource
	Title *string `json:"title,omitempty"`


	// Description - Description of the document or resource
	Description *string `json:"description,omitempty"`

}

// String returns a JSON representation of the model
func (o *Helplink) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
