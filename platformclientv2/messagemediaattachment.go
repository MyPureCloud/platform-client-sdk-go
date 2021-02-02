package platformclientv2
import (
	"encoding/json"
)

// Messagemediaattachment
type Messagemediaattachment struct { 
	// Url - The location of the media, useful for retrieving it
	Url *string `json:"url,omitempty"`


	// MediaType - The optional internet media type of the the media object.If null then the media type should be dictated by the url.
	MediaType *string `json:"mediaType,omitempty"`


	// ContentLength - The optional content length of the the media object, in bytes.
	ContentLength *int `json:"contentLength,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagemediaattachment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
