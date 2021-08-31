package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicvoicemail
type Queueconversationeventtopicvoicemail struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// UploadStatus
	UploadStatus *string `json:"uploadStatus,omitempty"`

}

func (o *Queueconversationeventtopicvoicemail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopicvoicemail
	
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

func (o *Queueconversationeventtopicvoicemail) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopicvoicemailMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopicvoicemailMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationeventtopicvoicemailMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if UploadStatus, ok := QueueconversationeventtopicvoicemailMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicvoicemail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
