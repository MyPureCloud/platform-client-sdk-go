package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Phonenumberstatus
type Phonenumberstatus struct { 
	// Callable - Indicates whether or not a phone number is callable.
	Callable *bool `json:"callable,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonenumberstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
