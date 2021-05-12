package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Managementunitreference - Management unit reference object for Workforce Management (ID/selfUri only)
type Managementunitreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Managementunitreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
