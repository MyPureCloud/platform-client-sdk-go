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

func (o *Emailattachment) MarshalJSON() ([]byte, error) {
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
		Name: o.Name,
		
		ContentPath: o.ContentPath,
		
		ContentType: o.ContentType,
		
		AttachmentId: o.AttachmentId,
		
		ContentLength: o.ContentLength,
		Alias:    (*Alias)(o),
	})
}

func (o *Emailattachment) UnmarshalJSON(b []byte) error {
	var EmailattachmentMap map[string]interface{}
	err := json.Unmarshal(b, &EmailattachmentMap)
	if err != nil {
		return err
	}
	
	if Name, ok := EmailattachmentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ContentPath, ok := EmailattachmentMap["contentPath"].(string); ok {
		o.ContentPath = &ContentPath
	}
    
	if ContentType, ok := EmailattachmentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if AttachmentId, ok := EmailattachmentMap["attachmentId"].(string); ok {
		o.AttachmentId = &AttachmentId
	}
    
	if ContentLength, ok := EmailattachmentMap["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Emailattachment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
