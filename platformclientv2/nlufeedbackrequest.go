package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Nlufeedbackrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nlufeedbackrequest

	

	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Intents *[]Intentfeedback `json:"intents,omitempty"`
		
		VersionId *string `json:"versionId,omitempty"`
		*Alias
	}{ 
		Text: u.Text,
		
		Intents: u.Intents,
		
		VersionId: u.VersionId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nlufeedbackrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
