package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Digits
type Digits struct { 
	// Digits - A string representing the digits pressed on phone.
	Digits *string `json:"digits,omitempty"`

}

// String returns a JSON representation of the model
func (o *Digits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
