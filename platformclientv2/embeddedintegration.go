package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Embeddedintegration
type Embeddedintegration struct { 
	// EnableWhitelist
	EnableWhitelist *bool `json:"enableWhitelist,omitempty"`


	// DomainWhitelist
	DomainWhitelist *[]string `json:"domainWhitelist,omitempty"`

}

// String returns a JSON representation of the model
func (o *Embeddedintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
