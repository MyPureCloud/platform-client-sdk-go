package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Openmessagecontent - Message content element.
type Openmessagecontent struct { 
	// ContentType - Type of this content element. If contentType = \"Attachment\" only one item is allowed.
	ContentType *string `json:"contentType,omitempty"`


	// Attachment - Attachment content.
	Attachment *Contentattachment `json:"attachment,omitempty"`

}

func (o *Openmessagecontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Openmessagecontent
	
	return json.Marshal(&struct { 
		ContentType *string `json:"contentType,omitempty"`
		
		Attachment *Contentattachment `json:"attachment,omitempty"`
		*Alias
	}{ 
		ContentType: o.ContentType,
		
		Attachment: o.Attachment,
		Alias:    (*Alias)(o),
	})
}

func (o *Openmessagecontent) UnmarshalJSON(b []byte) error {
	var OpenmessagecontentMap map[string]interface{}
	err := json.Unmarshal(b, &OpenmessagecontentMap)
	if err != nil {
		return err
	}
	
	if ContentType, ok := OpenmessagecontentMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Attachment, ok := OpenmessagecontentMap["attachment"].(map[string]interface{}); ok {
		AttachmentString, _ := json.Marshal(Attachment)
		json.Unmarshal(AttachmentString, &o.Attachment)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Openmessagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
