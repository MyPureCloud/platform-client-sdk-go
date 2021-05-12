package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Manager - Defines a SCIM manager.
type Manager struct { 
	// Value - The ID of the manager.
	Value *string `json:"value,omitempty"`


	// Ref - The reference URI of the manager's user record.
	Ref *string `json:"$ref,omitempty"`

}

// String returns a JSON representation of the model
func (o *Manager) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
