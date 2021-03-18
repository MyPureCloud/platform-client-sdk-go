package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailmessagestopicvoicemailmessage
type Voicemailmessagestopicvoicemailmessage struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`


	// AudioRecordingDurationSeconds
	AudioRecordingDurationSeconds *int `json:"audioRecordingDurationSeconds,omitempty"`


	// AudioRecordingSizeBytes
	AudioRecordingSizeBytes *int `json:"audioRecordingSizeBytes,omitempty"`


	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// CallerAddress
	CallerAddress *string `json:"callerAddress,omitempty"`


	// CallerName
	CallerName *string `json:"callerName,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`


	// Note
	Note *string `json:"note,omitempty"`


	// Deleted
	Deleted *bool `json:"deleted,omitempty"`


	// ModifiedByUserId
	ModifiedByUserId *string `json:"modifiedByUserId,omitempty"`


	// CopiedTo
	CopiedTo *[]Voicemailmessagestopicvoicemailcopyrecord `json:"copiedTo,omitempty"`


	// CopiedFrom
	CopiedFrom *Voicemailmessagestopicvoicemailcopyrecord `json:"copiedFrom,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailmessagestopicvoicemailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
