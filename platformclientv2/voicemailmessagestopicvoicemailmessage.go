package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Voicemailmessagestopicvoicemailmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailmessagestopicvoicemailmessage

	
	CreatedDate := new(string)
	if u.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(u.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		AudioRecordingDurationSeconds *int `json:"audioRecordingDurationSeconds,omitempty"`
		
		AudioRecordingSizeBytes *int `json:"audioRecordingSizeBytes,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		CallerAddress *string `json:"callerAddress,omitempty"`
		
		CallerName *string `json:"callerName,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		Note *string `json:"note,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		ModifiedByUserId *string `json:"modifiedByUserId,omitempty"`
		
		CopiedTo *[]Voicemailmessagestopicvoicemailcopyrecord `json:"copiedTo,omitempty"`
		
		CopiedFrom *Voicemailmessagestopicvoicemailcopyrecord `json:"copiedFrom,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Read: u.Read,
		
		AudioRecordingDurationSeconds: u.AudioRecordingDurationSeconds,
		
		AudioRecordingSizeBytes: u.AudioRecordingSizeBytes,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		
		CallerAddress: u.CallerAddress,
		
		CallerName: u.CallerName,
		
		Action: u.Action,
		
		Note: u.Note,
		
		Deleted: u.Deleted,
		
		ModifiedByUserId: u.ModifiedByUserId,
		
		CopiedTo: u.CopiedTo,
		
		CopiedFrom: u.CopiedFrom,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Voicemailmessagestopicvoicemailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
