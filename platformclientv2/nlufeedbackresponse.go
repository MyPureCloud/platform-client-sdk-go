package platformclientv2
import (
	"time"
	"encoding/json"
)

// Nlufeedbackresponse
type Nlufeedbackresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Text - The feedback text.
	Text *string `json:"text,omitempty"`


	// Intents - Detected intent of the utterance
	Intents *[]Intentfeedback `json:"intents,omitempty"`


	// Version - The domain version of the feedback.
	Version **Nludomainversion `json:"version,omitempty"`


	// DateCreated - The date when the feedback was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nlufeedbackresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
