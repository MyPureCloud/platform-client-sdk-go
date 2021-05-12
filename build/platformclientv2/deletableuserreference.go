package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Deletableuserreference - User reference with delete flag to remove the user from an associated entity
type Deletableuserreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Delete - If marked true, the user will be removed an associated entity
	Delete *bool `json:"delete,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Deletableuserreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
