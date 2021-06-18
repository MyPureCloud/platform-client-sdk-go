package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptiontopictranscriptresult
type Transcriptiontopictranscriptresult struct { 
	// UtteranceId
	UtteranceId *string `json:"utteranceId,omitempty"`


	// IsFinal
	IsFinal *bool `json:"isFinal,omitempty"`


	// Channel
	Channel *string `json:"channel,omitempty"`


	// Alternatives
	Alternatives *[]Transcriptiontopictranscriptalternative `json:"alternatives,omitempty"`


	// AgentAssistantId
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`


	// EngineId
	EngineId *string `json:"engineId,omitempty"`


	// Dialect
	Dialect *string `json:"dialect,omitempty"`


	// SpeechTextAnalyticsProgramId
	SpeechTextAnalyticsProgramId *string `json:"speechTextAnalyticsProgramId,omitempty"`


	// AgentAssistEnabled
	AgentAssistEnabled *bool `json:"agentAssistEnabled,omitempty"`


	// VoiceTranscriptionEnabled
	VoiceTranscriptionEnabled *bool `json:"voiceTranscriptionEnabled,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
