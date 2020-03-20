package platformclientv2
import (
	"encoding/json"
)

// Scimv2memberreference - Defines a reference to SCIM group members.
type Scimv2memberreference struct { 
	// VarType - The type of SCIM resource.
	VarType *string `json:"type,omitempty"`


	// Value - The ID of the group member. Can be \"userId\" or \"groupId\".
	Value *string `json:"value,omitempty"`


	// Ref - The reference URI of the SCIM resource.
	Ref *string `json:"$ref,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2memberreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
