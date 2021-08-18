package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Transcriptiontopictranscriptresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptiontopictranscriptresult

	

	return json.Marshal(&struct { 
		UtteranceId *string `json:"utteranceId,omitempty"`
		
		IsFinal *bool `json:"isFinal,omitempty"`
		
		Channel *string `json:"channel,omitempty"`
		
		Alternatives *[]Transcriptiontopictranscriptalternative `json:"alternatives,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		
		EngineId *string `json:"engineId,omitempty"`
		
		Dialect *string `json:"dialect,omitempty"`
		
		SpeechTextAnalyticsProgramId *string `json:"speechTextAnalyticsProgramId,omitempty"`
		
		AgentAssistEnabled *bool `json:"agentAssistEnabled,omitempty"`
		
		VoiceTranscriptionEnabled *bool `json:"voiceTranscriptionEnabled,omitempty"`
		*Alias
	}{ 
		UtteranceId: u.UtteranceId,
		
		IsFinal: u.IsFinal,
		
		Channel: u.Channel,
		
		Alternatives: u.Alternatives,
		
		AgentAssistantId: u.AgentAssistantId,
		
		EngineId: u.EngineId,
		
		Dialect: u.Dialect,
		
		SpeechTextAnalyticsProgramId: u.SpeechTextAnalyticsProgramId,
		
		AgentAssistEnabled: u.AgentAssistEnabled,
		
		VoiceTranscriptionEnabled: u.VoiceTranscriptionEnabled,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
