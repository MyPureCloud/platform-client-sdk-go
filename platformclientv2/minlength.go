package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Minlength
type Minlength struct { 
	// Min - A non-negative integer for a text-based schema field denoting the minimum smallest length a string field can contain for a schema instance.
	Min *int `json:"min,omitempty"`


	// Max - A non-negative integer for a text-based schema field denoting the maximum smallest length string the field can contain for a schema instance.
	Max *int `json:"max,omitempty"`

}

// String returns a JSON representation of the model
func (o *Minlength) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
