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

func (o *Transcriptiontopictranscriptresult) MarshalJSON() ([]byte, error) {
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
		UtteranceId: o.UtteranceId,
		
		IsFinal: o.IsFinal,
		
		Channel: o.Channel,
		
		Alternatives: o.Alternatives,
		
		AgentAssistantId: o.AgentAssistantId,
		
		EngineId: o.EngineId,
		
		Dialect: o.Dialect,
		
		SpeechTextAnalyticsProgramId: o.SpeechTextAnalyticsProgramId,
		
		AgentAssistEnabled: o.AgentAssistEnabled,
		
		VoiceTranscriptionEnabled: o.VoiceTranscriptionEnabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcriptiontopictranscriptresult) UnmarshalJSON(b []byte) error {
	var TranscriptiontopictranscriptresultMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptiontopictranscriptresultMap)
	if err != nil {
		return err
	}
	
	if UtteranceId, ok := TranscriptiontopictranscriptresultMap["utteranceId"].(string); ok {
		o.UtteranceId = &UtteranceId
	}
    
	if IsFinal, ok := TranscriptiontopictranscriptresultMap["isFinal"].(bool); ok {
		o.IsFinal = &IsFinal
	}
    
	if Channel, ok := TranscriptiontopictranscriptresultMap["channel"].(string); ok {
		o.Channel = &Channel
	}
    
	if Alternatives, ok := TranscriptiontopictranscriptresultMap["alternatives"].([]interface{}); ok {
		AlternativesString, _ := json.Marshal(Alternatives)
		json.Unmarshal(AlternativesString, &o.Alternatives)
	}
	
	if AgentAssistantId, ok := TranscriptiontopictranscriptresultMap["agentAssistantId"].(string); ok {
		o.AgentAssistantId = &AgentAssistantId
	}
    
	if EngineId, ok := TranscriptiontopictranscriptresultMap["engineId"].(string); ok {
		o.EngineId = &EngineId
	}
    
	if Dialect, ok := TranscriptiontopictranscriptresultMap["dialect"].(string); ok {
		o.Dialect = &Dialect
	}
    
	if SpeechTextAnalyticsProgramId, ok := TranscriptiontopictranscriptresultMap["speechTextAnalyticsProgramId"].(string); ok {
		o.SpeechTextAnalyticsProgramId = &SpeechTextAnalyticsProgramId
	}
    
	if AgentAssistEnabled, ok := TranscriptiontopictranscriptresultMap["agentAssistEnabled"].(bool); ok {
		o.AgentAssistEnabled = &AgentAssistEnabled
	}
    
	if VoiceTranscriptionEnabled, ok := TranscriptiontopictranscriptresultMap["voiceTranscriptionEnabled"].(bool); ok {
		o.VoiceTranscriptionEnabled = &VoiceTranscriptionEnabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
