package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulerule
type Learningmodulerule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// IsActive - If true, rule is active
	IsActive *bool `json:"isActive,omitempty"`


	// Parts - The parts of a learning module rule
	Parts *[]Learningmoduleruleparts `json:"parts,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningmodulerule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
