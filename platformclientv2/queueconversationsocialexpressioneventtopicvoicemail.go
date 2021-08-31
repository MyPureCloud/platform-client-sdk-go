package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicvoicemail
type Queueconversationsocialexpressioneventtopicvoicemail struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// UploadStatus
	UploadStatus *string `json:"uploadStatus,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicvoicemail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicvoicemail
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		UploadStatus *string `json:"uploadStatus,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		UploadStatus: o.UploadStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationsocialexpressioneventtopicvoicemail) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicvoicemailMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicvoicemailMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationsocialexpressioneventtopicvoicemailMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if UploadStatus, ok := QueueconversationsocialexpressioneventtopicvoicemailMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicvoicemail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
