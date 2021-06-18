package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptionmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
