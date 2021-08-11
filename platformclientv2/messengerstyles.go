package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Messengerstyles
type Messengerstyles struct { 
	// PrimaryColor - The primary color of messenger in hexadecimal
	PrimaryColor *string `json:"primaryColor,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messengerstyles) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
