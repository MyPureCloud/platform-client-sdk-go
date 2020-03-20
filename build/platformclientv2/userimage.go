package platformclientv2
import (
	"encoding/json"
)

// Userimage
type Userimage struct { 
	// Resolution - Height and/or width of image. ex: 640x480 or x128
	Resolution *string `json:"resolution,omitempty"`


	// ImageUri
	ImageUri *string `json:"imageUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userimage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
