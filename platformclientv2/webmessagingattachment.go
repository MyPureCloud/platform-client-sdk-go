package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webmessagingattachment - Attachment object.
type Webmessagingattachment struct { 
	// Id - The globally unique identifier for the object.
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


	// FileSize - The file size associated with the file
	FileSize *int `json:"fileSize,omitempty"`

}

func (o *Webmessagingattachment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webmessagingattachment
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Mime *string `json:"mime,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Sha256 *string `json:"sha256,omitempty"`
		
		Filename *string `json:"filename,omitempty"`
		
		FileSize *int `json:"fileSize,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		MediaType: o.MediaType,
		
		Url: o.Url,
		
		Mime: o.Mime,
		
		Text: o.Text,
		
		Sha256: o.Sha256,
		
		Filename: o.Filename,
		
		FileSize: o.FileSize,
		Alias:    (*Alias)(o),
	})
}

func (o *Webmessagingattachment) UnmarshalJSON(b []byte) error {
	var WebmessagingattachmentMap map[string]interface{}
	err := json.Unmarshal(b, &WebmessagingattachmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebmessagingattachmentMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if MediaType, ok := WebmessagingattachmentMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if Url, ok := WebmessagingattachmentMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if Mime, ok := WebmessagingattachmentMap["mime"].(string); ok {
		o.Mime = &Mime
	}
	
	if Text, ok := WebmessagingattachmentMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if Sha256, ok := WebmessagingattachmentMap["sha256"].(string); ok {
		o.Sha256 = &Sha256
	}
	
	if Filename, ok := WebmessagingattachmentMap["filename"].(string); ok {
		o.Filename = &Filename
	}
	
	if FileSize, ok := WebmessagingattachmentMap["fileSize"].(float64); ok {
		FileSizeInt := int(FileSize)
		o.FileSize = &FileSizeInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webmessagingattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
