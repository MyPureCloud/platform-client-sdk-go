package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Scimuserrole - Defines a user role.
type Scimuserrole struct { 
	// Value - The role of the Genesys Cloud user.
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimuserrole) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
