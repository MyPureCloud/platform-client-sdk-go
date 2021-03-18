package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeversioninformation
type Edgeversioninformation struct { 
	// SoftwareVersion
	SoftwareVersion *string `json:"softwareVersion,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgeversioninformation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
