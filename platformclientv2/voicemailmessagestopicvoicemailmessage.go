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

func (o *Voicemailmessagestopicvoicemailmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailmessagestopicvoicemailmessage
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Read: o.Read,
		
		AudioRecordingDurationSeconds: o.AudioRecordingDurationSeconds,
		
		AudioRecordingSizeBytes: o.AudioRecordingSizeBytes,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		
		CallerAddress: o.CallerAddress,
		
		CallerName: o.CallerName,
		
		Action: o.Action,
		
		Note: o.Note,
		
		Deleted: o.Deleted,
		
		ModifiedByUserId: o.ModifiedByUserId,
		
		CopiedTo: o.CopiedTo,
		
		CopiedFrom: o.CopiedFrom,
		Alias:    (*Alias)(o),
	})
}

func (o *Voicemailmessagestopicvoicemailmessage) UnmarshalJSON(b []byte) error {
	var VoicemailmessagestopicvoicemailmessageMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailmessagestopicvoicemailmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := VoicemailmessagestopicvoicemailmessageMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Read, ok := VoicemailmessagestopicvoicemailmessageMap["read"].(bool); ok {
		o.Read = &Read
	}
	
	if AudioRecordingDurationSeconds, ok := VoicemailmessagestopicvoicemailmessageMap["audioRecordingDurationSeconds"].(float64); ok {
		AudioRecordingDurationSecondsInt := int(AudioRecordingDurationSeconds)
		o.AudioRecordingDurationSeconds = &AudioRecordingDurationSecondsInt
	}
	
	if AudioRecordingSizeBytes, ok := VoicemailmessagestopicvoicemailmessageMap["audioRecordingSizeBytes"].(float64); ok {
		AudioRecordingSizeBytesInt := int(AudioRecordingSizeBytes)
		o.AudioRecordingSizeBytes = &AudioRecordingSizeBytesInt
	}
	
	if createdDateString, ok := VoicemailmessagestopicvoicemailmessageMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := VoicemailmessagestopicvoicemailmessageMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if CallerAddress, ok := VoicemailmessagestopicvoicemailmessageMap["callerAddress"].(string); ok {
		o.CallerAddress = &CallerAddress
	}
	
	if CallerName, ok := VoicemailmessagestopicvoicemailmessageMap["callerName"].(string); ok {
		o.CallerName = &CallerName
	}
	
	if Action, ok := VoicemailmessagestopicvoicemailmessageMap["action"].(string); ok {
		o.Action = &Action
	}
	
	if Note, ok := VoicemailmessagestopicvoicemailmessageMap["note"].(string); ok {
		o.Note = &Note
	}
	
	if Deleted, ok := VoicemailmessagestopicvoicemailmessageMap["deleted"].(bool); ok {
		o.Deleted = &Deleted
	}
	
	if ModifiedByUserId, ok := VoicemailmessagestopicvoicemailmessageMap["modifiedByUserId"].(string); ok {
		o.ModifiedByUserId = &ModifiedByUserId
	}
	
	if CopiedTo, ok := VoicemailmessagestopicvoicemailmessageMap["copiedTo"].([]interface{}); ok {
		CopiedToString, _ := json.Marshal(CopiedTo)
		json.Unmarshal(CopiedToString, &o.CopiedTo)
	}
	
	if CopiedFrom, ok := VoicemailmessagestopicvoicemailmessageMap["copiedFrom"].(map[string]interface{}); ok {
		CopiedFromString, _ := json.Marshal(CopiedFrom)
		json.Unmarshal(CopiedFromString, &o.CopiedFrom)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailmessagestopicvoicemailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
