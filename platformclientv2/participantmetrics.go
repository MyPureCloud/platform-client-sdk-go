package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Participantmetrics
type Participantmetrics struct { 
	// AgentDurationPercentage - Percentage of Agent duration in the conversation
	AgentDurationPercentage *float64 `json:"agentDurationPercentage,omitempty"`


	// CustomerDurationPercentage - Percentage of Customer duration in the conversation
	CustomerDurationPercentage *float64 `json:"customerDurationPercentage,omitempty"`


	// SilenceDurationPercentage - Percentage of Silence duration in the conversation
	SilenceDurationPercentage *float64 `json:"silenceDurationPercentage,omitempty"`


	// IvrDurationPercentage - Percentage of IVR duration in the conversation
	IvrDurationPercentage *float64 `json:"ivrDurationPercentage,omitempty"`


	// AcdDurationPercentage - Percentage of ACD duration in the conversation
	AcdDurationPercentage *float64 `json:"acdDurationPercentage,omitempty"`


	// OvertalkDurationPercentage - Percentage of Overtalk duration in the conversation
	OvertalkDurationPercentage *float64 `json:"overtalkDurationPercentage,omitempty"`


	// OtherDurationPercentage - Percentage of Other events duration in the conversation
	OtherDurationPercentage *float64 `json:"otherDurationPercentage,omitempty"`


	// OvertalkCount - Number of Overtalks in the conversation
	OvertalkCount *int `json:"overtalkCount,omitempty"`

}

func (o *Participantmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Participantmetrics
	
	return json.Marshal(&struct { 
		AgentDurationPercentage *float64 `json:"agentDurationPercentage,omitempty"`
		
		CustomerDurationPercentage *float64 `json:"customerDurationPercentage,omitempty"`
		
		SilenceDurationPercentage *float64 `json:"silenceDurationPercentage,omitempty"`
		
		IvrDurationPercentage *float64 `json:"ivrDurationPercentage,omitempty"`
		
		AcdDurationPercentage *float64 `json:"acdDurationPercentage,omitempty"`
		
		OvertalkDurationPercentage *float64 `json:"overtalkDurationPercentage,omitempty"`
		
		OtherDurationPercentage *float64 `json:"otherDurationPercentage,omitempty"`
		
		OvertalkCount *int `json:"overtalkCount,omitempty"`
		*Alias
	}{ 
		AgentDurationPercentage: o.AgentDurationPercentage,
		
		CustomerDurationPercentage: o.CustomerDurationPercentage,
		
		SilenceDurationPercentage: o.SilenceDurationPercentage,
		
		IvrDurationPercentage: o.IvrDurationPercentage,
		
		AcdDurationPercentage: o.AcdDurationPercentage,
		
		OvertalkDurationPercentage: o.OvertalkDurationPercentage,
		
		OtherDurationPercentage: o.OtherDurationPercentage,
		
		OvertalkCount: o.OvertalkCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Participantmetrics) UnmarshalJSON(b []byte) error {
	var ParticipantmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &ParticipantmetricsMap)
	if err != nil {
		return err
	}
	
	if AgentDurationPercentage, ok := ParticipantmetricsMap["agentDurationPercentage"].(float64); ok {
		o.AgentDurationPercentage = &AgentDurationPercentage
	}
	
	if CustomerDurationPercentage, ok := ParticipantmetricsMap["customerDurationPercentage"].(float64); ok {
		o.CustomerDurationPercentage = &CustomerDurationPercentage
	}
	
	if SilenceDurationPercentage, ok := ParticipantmetricsMap["silenceDurationPercentage"].(float64); ok {
		o.SilenceDurationPercentage = &SilenceDurationPercentage
	}
	
	if IvrDurationPercentage, ok := ParticipantmetricsMap["ivrDurationPercentage"].(float64); ok {
		o.IvrDurationPercentage = &IvrDurationPercentage
	}
	
	if AcdDurationPercentage, ok := ParticipantmetricsMap["acdDurationPercentage"].(float64); ok {
		o.AcdDurationPercentage = &AcdDurationPercentage
	}
	
	if OvertalkDurationPercentage, ok := ParticipantmetricsMap["overtalkDurationPercentage"].(float64); ok {
		o.OvertalkDurationPercentage = &OvertalkDurationPercentage
	}
	
	if OtherDurationPercentage, ok := ParticipantmetricsMap["otherDurationPercentage"].(float64); ok {
		o.OtherDurationPercentage = &OtherDurationPercentage
	}
	
	if OvertalkCount, ok := ParticipantmetricsMap["overtalkCount"].(float64); ok {
		OvertalkCountInt := int(OvertalkCount)
		o.OvertalkCount = &OvertalkCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Participantmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
