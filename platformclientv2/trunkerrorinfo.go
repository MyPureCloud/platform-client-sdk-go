package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkerrorinfo
type Trunkerrorinfo struct { 
	// Text
	Text *string `json:"text,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`


	// Details
	Details *Trunkerrorinfodetails `json:"details,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkerrorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
