package platformclientv2
import (
	"encoding/json"
)

// Nlufeedbackrequest
type Nlufeedbackrequest struct { 
	// Text - The feedback text.
	Text *string `json:"text,omitempty"`


	// Intents - Detected intent of the utterance
	Intents *[]Intentfeedback `json:"intents,omitempty"`


	// VersionId - The domain version ID of the feedback.
	VersionId *string `json:"versionId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nlufeedbackrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
