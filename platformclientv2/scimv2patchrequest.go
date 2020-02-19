package platformclientv2
import (
	"encoding/json"
)

// Scimv2patchrequest - Defines a SCIM PATCH request. See section 3.5.2 \"Modifying with PATCH\" in RFC 7644 for details.
type Scimv2patchrequest struct { 
	// Schemas - The list of schemas used in the PATCH request.
	Schemas *[]string `json:"schemas,omitempty"`


	// Operations - The list of operations to perform for the PATCH request.
	Operations *[]Scimv2patchoperation `json:"Operations,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2patchrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
