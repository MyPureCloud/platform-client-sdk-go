package platformclientv2
import (
	"encoding/json"
)

// Transcriptproperty
type Transcriptproperty struct { 
	// Url - The pre-signed S3 URL of the transcript
	Url *string `json:"url,omitempty"`


	// Conversation - The conversation reference
	Conversation *Addressableentityref `json:"conversation,omitempty"`


	// CommunicationId - The communication ID
	CommunicationId *string `json:"communicationId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcriptproperty) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
