package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Openmessagecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
