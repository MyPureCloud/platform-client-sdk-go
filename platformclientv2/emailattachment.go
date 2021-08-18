package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailattachment
type Emailattachment struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// ContentPath
	ContentPath *string `json:"contentPath,omitempty"`


	// ContentType
	ContentType *string `json:"contentType,omitempty"`


	// AttachmentId
	AttachmentId *string `json:"attachmentId,omitempty"`


	// ContentLength
	ContentLength *int `json:"contentLength,omitempty"`

}

func (u *Emailattachment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailattachment

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ContentPath *string `json:"contentPath,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		AttachmentId *string `json:"attachmentId,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		ContentPath: u.ContentPath,
		
		ContentType: u.ContentType,
		
		AttachmentId: u.AttachmentId,
		
		ContentLength: u.ContentLength,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Emailattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
