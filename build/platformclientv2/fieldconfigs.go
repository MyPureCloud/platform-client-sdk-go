package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Fieldconfigs
type Fieldconfigs struct { 
	// Org
	Org *Fieldconfig `json:"org,omitempty"`


	// Person
	Person *Fieldconfig `json:"person,omitempty"`


	// Group
	Group *Fieldconfig `json:"group,omitempty"`


	// ExternalContact
	ExternalContact *Fieldconfig `json:"externalContact,omitempty"`

}

// String returns a JSON representation of the model
func (o *Fieldconfigs) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
