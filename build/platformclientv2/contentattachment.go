package platformclientv2
import (
	"encoding/json"
)

// Contentattachment - Attachment object
type Contentattachment struct { 
	// Id - Vendor specific ID for media. For example, a LINE sticker ID
	Id *string `json:"id,omitempty"`


	// MediaType - The type of media this instance represents
	MediaType *string `json:"mediaType,omitempty"`


	// Url - Content element url
	Url *string `json:"url,omitempty"`


	// Mime - Content mime type from https://www.iana.org/assignments/media-types/media-types.xhtml
	Mime *string `json:"mime,omitempty"`


	// Text - Text message associated with media element: e.g. caption in case of image.
	Text *string `json:"text,omitempty"`


	// Sha256 - Secure hash of the media content
	Sha256 *string `json:"sha256,omitempty"`


	// Filename - Suggested file name for media file
	Filename *string `json:"filename,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentattachment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
