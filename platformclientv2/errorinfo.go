package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Errorinfo
type Errorinfo struct { 
	// Message
	Message *string `json:"message,omitempty"`


	// Code
	Code *string `json:"code,omitempty"`

}

// String returns a JSON representation of the model
func (o *Errorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
