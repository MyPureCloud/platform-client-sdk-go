package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Copyattachmentsrequest
type Copyattachmentsrequest struct { 
	// SourceMessage - A reference to the email message within the current conversation that owns the attachments to be copied
	SourceMessage *Domainentityref `json:"sourceMessage,omitempty"`


	// Attachments - A list of attachments that will be copied from the source message to the current draft
	Attachments *[]Attachment `json:"attachments,omitempty"`

}

func (u *Copyattachmentsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Copyattachmentsrequest

	

	return json.Marshal(&struct { 
		SourceMessage *Domainentityref `json:"sourceMessage,omitempty"`
		
		Attachments *[]Attachment `json:"attachments,omitempty"`
		*Alias
	}{ 
		SourceMessage: u.SourceMessage,
		
		Attachments: u.Attachments,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Copyattachmentsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
