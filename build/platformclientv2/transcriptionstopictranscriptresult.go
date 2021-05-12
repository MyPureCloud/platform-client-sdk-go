package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptionstopictranscriptresult
type Transcriptionstopictranscriptresult struct { 
	// UtteranceId
	UtteranceId *string `json:"utteranceId,omitempty"`


	// IsFinal
	IsFinal *bool `json:"isFinal,omitempty"`


	// Channel
	Channel *string `json:"channel,omitempty"`


	// Alternatives
	Alternatives *[]Transcriptionstopictranscriptalternative `json:"alternatives,omitempty"`


	// AgentAssistantId
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`


	// EngineId
	EngineId *string `json:"engineId,omitempty"`


	// Dialect
	Dialect *string `json:"dialect,omitempty"`


	// SpeechTextAnalyticsProgramId
	SpeechTextAnalyticsProgramId *string `json:"speechTextAnalyticsProgramId,omitempty"`


	// StartTimeMs
	StartTimeMs *int `json:"startTimeMs,omitempty"`


	// OffsetMs
	OffsetMs *int `json:"offsetMs,omitempty"`


	// DurationMs
	DurationMs *int `json:"durationMs,omitempty"`


	// AgentAssistEnabled
	AgentAssistEnabled *bool `json:"agentAssistEnabled,omitempty"`


	// VoiceTranscriptionEnabled
	VoiceTranscriptionEnabled *bool `json:"voiceTranscriptionEnabled,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcriptionstopictranscriptresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
