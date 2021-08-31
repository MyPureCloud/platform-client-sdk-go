package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingjobfailedrecording
type Recordingjobfailedrecording struct { 
	// Conversation - Conversation
	Conversation *Addressableentityref `json:"conversation,omitempty"`


	// Recording - Recording
	Recording *Addressableentityref `json:"recording,omitempty"`

}

func (o *Recordingjobfailedrecording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingjobfailedrecording
	
	return json.Marshal(&struct { 
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		
		Recording *Addressableentityref `json:"recording,omitempty"`
		*Alias
	}{ 
		Conversation: o.Conversation,
		
		Recording: o.Recording,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingjobfailedrecording) UnmarshalJSON(b []byte) error {
	var RecordingjobfailedrecordingMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingjobfailedrecordingMap)
	if err != nil {
		return err
	}
	
	if Conversation, ok := RecordingjobfailedrecordingMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if Recording, ok := RecordingjobfailedrecordingMap["recording"].(map[string]interface{}); ok {
		RecordingString, _ := json.Marshal(Recording)
		json.Unmarshal(RecordingString, &o.Recording)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingjobfailedrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
