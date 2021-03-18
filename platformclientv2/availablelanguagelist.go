package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Availablelanguagelist
type Availablelanguagelist struct { 
	// Languages
	Languages *[]string `json:"languages,omitempty"`

}

// String returns a JSON representation of the model
func (o *Availablelanguagelist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
