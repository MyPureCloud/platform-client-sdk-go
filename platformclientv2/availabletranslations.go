package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Availabletranslations
type Availabletranslations struct { 
	// OrgSpecific
	OrgSpecific *[]string `json:"orgSpecific,omitempty"`


	// Builtin
	Builtin *[]string `json:"builtin,omitempty"`

}

// String returns a JSON representation of the model
func (o *Availabletranslations) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
