package platformclientv2
import (
	"encoding/json"
)

// Messagemediadata
type Messagemediadata struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Url - The location of the media, useful for retrieving it
	Url *string `json:"url,omitempty"`


	// MediaType - The optional internet media type of the the media object.  If null then the media type should be dictated by the url.
	MediaType *string `json:"mediaType,omitempty"`


	// ContentLengthBytes - The optional content length of the the media object, in bytes.
	ContentLengthBytes *int32 `json:"contentLengthBytes,omitempty"`


	// UploadUrl - The URL returned to upload an attachment
	UploadUrl *string `json:"uploadUrl,omitempty"`


	// Status - The status of the media, indicates if the media is in the process of uploading. If the upload fails, the media becomes invalid
	Status *string `json:"status,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagemediadata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
