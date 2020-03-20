package platformclientv2
import (
	"encoding/json"
)

// Documentthumbnail
type Documentthumbnail struct { 
	// Resolution
	Resolution *string `json:"resolution,omitempty"`


	// ImageUri
	ImageUri *string `json:"imageUri,omitempty"`


	// Height
	Height *int32 `json:"height,omitempty"`


	// Width
	Width *int32 `json:"width,omitempty"`

}

// String returns a JSON representation of the model
func (o *Documentthumbnail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
