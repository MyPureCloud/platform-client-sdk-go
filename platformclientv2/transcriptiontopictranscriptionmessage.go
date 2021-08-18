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

func (u *Transcriptiontopictranscriptionmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptiontopictranscriptionmessage

	
	EventTime := new(string)
	if u.EventTime != nil {
		
		*EventTime = timeutil.Strftime(u.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		OrganizationId: u.OrganizationId,
		
		ConversationId: u.ConversationId,
		
		CommunicationId: u.CommunicationId,
		
		SessionStartTimeMs: u.SessionStartTimeMs,
		
		TranscriptionStartTimeMs: u.TranscriptionStartTimeMs,
		
		Transcripts: u.Transcripts,
		
		Status: u.Status,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptionmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
