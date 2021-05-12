package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkinstancetopictrunkerrorinfodetails
type Trunkinstancetopictrunkerrorinfodetails struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// Hostname
	Hostname *string `json:"hostname,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkerrorinfodetails) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
