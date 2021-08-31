package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcallbackeventtopicvoicemail
type Conversationcallbackeventtopicvoicemail struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// UploadStatus
	UploadStatus *string `json:"uploadStatus,omitempty"`

}

func (o *Conversationcallbackeventtopicvoicemail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcallbackeventtopicvoicemail
	
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

func (o *Conversationcallbackeventtopicvoicemail) UnmarshalJSON(b []byte) error {
	var ConversationcallbackeventtopicvoicemailMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcallbackeventtopicvoicemailMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationcallbackeventtopicvoicemailMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if UploadStatus, ok := ConversationcallbackeventtopicvoicemailMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcallbackeventtopicvoicemail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
