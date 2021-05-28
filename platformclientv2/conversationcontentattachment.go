package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentattachment - Attachment object.
type Conversationcontentattachment struct { 
	// Id - Provider specific ID for attachment. For example, a LINE sticker ID.
	Id *string `json:"id,omitempty"`


	// MediaType - The type of attachment this instance represents.
	MediaType *string `json:"mediaType,omitempty"`


	// Url - URL of the attachment.
	Url *string `json:"url,omitempty"`


	// Mime - Attachment mime type (https://www.iana.org/assignments/media-types/media-types.xhtml).
	Mime *string `json:"mime,omitempty"`


	// Text - Text associated with attachment such as an image caption.
	Text *string `json:"text,omitempty"`


	// Sha256 - Secure hash of the attachment content.
	Sha256 *string `json:"sha256,omitempty"`


	// Filename - Suggested file name for attachment.
	Filename *string `json:"filename,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcontentattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
