package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Publishform
type Publishform struct { 
	// Published - Is this form published
	Published *bool `json:"published,omitempty"`


	// Id - Unique Id for this version of this form
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Publishform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
