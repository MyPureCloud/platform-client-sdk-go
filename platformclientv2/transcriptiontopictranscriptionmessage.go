package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptiontopictranscriptionmessage
type Transcriptiontopictranscriptionmessage struct { 
	// EventTime
	EventTime *time.Time `json:"eventTime,omitempty"`


	// OrganizationId
	OrganizationId *string `json:"organizationId,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// CommunicationId
	CommunicationId *string `json:"communicationId,omitempty"`


	// SessionStartTimeMs
	SessionStartTimeMs *int `json:"sessionStartTimeMs,omitempty"`


	// TranscriptionStartTimeMs
	TranscriptionStartTimeMs *int `json:"transcriptionStartTimeMs,omitempty"`


	// Transcripts
	Transcripts *[]Transcriptiontopictranscriptresult `json:"transcripts,omitempty"`


	// Status
	Status *Transcriptiontopictranscriptionrequeststatus `json:"status,omitempty"`

}

func (o *Transcriptiontopictranscriptionmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptiontopictranscriptionmessage
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		EventTime *string `json:"eventTime,omitempty"`
		
		OrganizationId *string `json:"organizationId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		SessionStartTimeMs *int `json:"sessionStartTimeMs,omitempty"`
		
		TranscriptionStartTimeMs *int `json:"transcriptionStartTimeMs,omitempty"`
		
		Transcripts *[]Transcriptiontopictranscriptresult `json:"transcripts,omitempty"`
		
		Status *Transcriptiontopictranscriptionrequeststatus `json:"status,omitempty"`
		*Alias
	}{ 
		EventTime: EventTime,
		
		OrganizationId: o.OrganizationId,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		SessionStartTimeMs: o.SessionStartTimeMs,
		
		TranscriptionStartTimeMs: o.TranscriptionStartTimeMs,
		
		Transcripts: o.Transcripts,
		
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcriptiontopictranscriptionmessage) UnmarshalJSON(b []byte) error {
	var TranscriptiontopictranscriptionmessageMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptiontopictranscriptionmessageMap)
	if err != nil {
		return err
	}
	
	if eventTimeString, ok := TranscriptiontopictranscriptionmessageMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if OrganizationId, ok := TranscriptiontopictranscriptionmessageMap["organizationId"].(string); ok {
		o.OrganizationId = &OrganizationId
	}
	
	if ConversationId, ok := TranscriptiontopictranscriptionmessageMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if CommunicationId, ok := TranscriptiontopictranscriptionmessageMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
	
	if SessionStartTimeMs, ok := TranscriptiontopictranscriptionmessageMap["sessionStartTimeMs"].(float64); ok {
		SessionStartTimeMsInt := int(SessionStartTimeMs)
		o.SessionStartTimeMs = &SessionStartTimeMsInt
	}
	
	if TranscriptionStartTimeMs, ok := TranscriptiontopictranscriptionmessageMap["transcriptionStartTimeMs"].(float64); ok {
		TranscriptionStartTimeMsInt := int(TranscriptionStartTimeMs)
		o.TranscriptionStartTimeMs = &TranscriptionStartTimeMsInt
	}
	
	if Transcripts, ok := TranscriptiontopictranscriptionmessageMap["transcripts"].([]interface{}); ok {
		TranscriptsString, _ := json.Marshal(Transcripts)
		json.Unmarshal(TranscriptsString, &o.Transcripts)
	}
	
	if Status, ok := TranscriptiontopictranscriptionmessageMap["status"].(map[string]interface{}); ok {
		StatusString, _ := json.Marshal(Status)
		json.Unmarshal(StatusString, &o.Status)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptionmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
