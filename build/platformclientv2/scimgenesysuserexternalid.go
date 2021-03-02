package platformclientv2
import (
	"encoding/json"
)

// Scimgenesysuserexternalid - External Identifiers of user. The external identifier must be unique within the organization and the 'authority'
type Scimgenesysuserexternalid struct { 
	// Authority - Authority, or scope, of \"externalId\". Allows multiple external identifiers to be defined. Represents the source of the external identifier.
	Authority *string `json:"authority,omitempty"`


	// Value - Identifier of the user in an external system.
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimgenesysuserexternalid) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
