package platformclientv2
import (
	"encoding/json"
)

// Greetingmediainfo
type Greetingmediainfo struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// MediaFileUri
	MediaFileUri *string `json:"mediaFileUri,omitempty"`


	// MediaImageUri
	MediaImageUri *string `json:"mediaImageUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Greetingmediainfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
