package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicvoicemail - The voicemail data to be used when this callback is an ACD voicemail.
type Conversationeventtopicvoicemail struct { 
	// Id - The voicemail id
	Id *string `json:"id,omitempty"`


	// UploadStatus - current state of the voicemail upload
	UploadStatus *string `json:"uploadStatus,omitempty"`

}

func (o *Conversationeventtopicvoicemail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicvoicemail
	
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

func (o *Conversationeventtopicvoicemail) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicvoicemailMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicvoicemailMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationeventtopicvoicemailMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if UploadStatus, ok := ConversationeventtopicvoicemailMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicvoicemail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
