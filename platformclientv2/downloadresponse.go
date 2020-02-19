package platformclientv2
import (
	"encoding/json"
)

// Downloadresponse
type Downloadresponse struct { 
	// ContentLocationUri
	ContentLocationUri *string `json:"contentLocationUri,omitempty"`


	// ImageUri
	ImageUri *string `json:"imageUri,omitempty"`


	// Thumbnails
	Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`

}

// String returns a JSON representation of the model
func (o *Downloadresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
