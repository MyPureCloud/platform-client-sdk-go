package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Answeroption
type Answeroption struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// Value
	Value *int `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Answeroption) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
