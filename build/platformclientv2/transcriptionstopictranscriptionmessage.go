package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptionstopictranscriptionmessage
type Transcriptionstopictranscriptionmessage struct { 
	// EventTime
	EventTime *time.Time `json:"eventTime,omitempty"`


	// BatchIntervalMs
	BatchIntervalMs *int `json:"batchIntervalMs,omitempty"`


	// OrganizationId
	OrganizationId *string `json:"organizationId,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// CommunicationId
	CommunicationId *string `json:"communicationId,omitempty"`


	// SessionId
	SessionId *string `json:"sessionId,omitempty"`


	// Transcripts
	Transcripts *[]Transcriptionstopictranscriptresult `json:"transcripts,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcriptionstopictranscriptionmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
