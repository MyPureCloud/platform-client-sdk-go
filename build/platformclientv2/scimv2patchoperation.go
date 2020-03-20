package platformclientv2
import (
	"encoding/json"
)

// Scimv2patchoperation - Defines a SCIM PATCH operation. The path and value follow very specific rules based on operation types. See section 3.5.2 \"Modifying with PATCH\" in RFC 7644 for details.
type Scimv2patchoperation struct { 
	// Op - The PATCH operation to perform.
	Op *string `json:"op,omitempty"`


	// Path - The attribute path that describes the target of the operation. Required for a \"remove\" operation.
	Path *string `json:"path,omitempty"`


	// Value - The value to set in the path.
	Value *Jsonnode `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2patchoperation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
