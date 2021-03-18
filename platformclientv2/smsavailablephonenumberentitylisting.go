package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Smsavailablephonenumberentitylisting
type Smsavailablephonenumberentitylisting struct { 
	// Entities
	Entities *[]Smsavailablephonenumber `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Smsavailablephonenumberentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
