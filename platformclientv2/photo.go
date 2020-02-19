package platformclientv2
import (
	"encoding/json"
)

// Photo - Defines a SCIM photo.
type Photo struct { 
	// Value - The URI of the photo. Photos are limited to 240 KB and JPG, GIF, and PNG formats. Returns a JPG.
	Value *string `json:"value,omitempty"`


	// VarType - The type of photo.
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Photo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
