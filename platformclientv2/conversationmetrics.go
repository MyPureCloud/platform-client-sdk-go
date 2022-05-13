package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmetrics
type Conversationmetrics struct { 
	// Conversation - The Conversation Reference
	Conversation *Addressableentityref `json:"conversation,omitempty"`


	// SentimentScore - The Sentiment Score
	SentimentScore *float64 `json:"sentimentScore,omitempty"`


	// SentimentTrend - The Sentiment Trend
	SentimentTrend *float64 `json:"sentimentTrend,omitempty"`


	// SentimentTrendClass - The Sentiment Trend Class
	SentimentTrendClass *string `json:"sentimentTrendClass,omitempty"`


	// ParticipantMetrics - The Participant Metrics
	ParticipantMetrics *Participantmetrics `json:"participantMetrics,omitempty"`

}

func (o *Conversationmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmetrics
	
	return json.Marshal(&struct { 
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		
		SentimentScore *float64 `json:"sentimentScore,omitempty"`
		
		SentimentTrend *float64 `json:"sentimentTrend,omitempty"`
		
		SentimentTrendClass *string `json:"sentimentTrendClass,omitempty"`
		
		ParticipantMetrics *Participantmetrics `json:"participantMetrics,omitempty"`
		*Alias
	}{ 
		Conversation: o.Conversation,
		
		SentimentScore: o.SentimentScore,
		
		SentimentTrend: o.SentimentTrend,
		
		SentimentTrendClass: o.SentimentTrendClass,
		
		ParticipantMetrics: o.ParticipantMetrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationmetrics) UnmarshalJSON(b []byte) error {
	var ConversationmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmetricsMap)
	if err != nil {
		return err
	}
	
	if Conversation, ok := ConversationmetricsMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if SentimentScore, ok := ConversationmetricsMap["sentimentScore"].(float64); ok {
		o.SentimentScore = &SentimentScore
	}
    
	if SentimentTrend, ok := ConversationmetricsMap["sentimentTrend"].(float64); ok {
		o.SentimentTrend = &SentimentTrend
	}
    
	if SentimentTrendClass, ok := ConversationmetricsMap["sentimentTrendClass"].(string); ok {
		o.SentimentTrendClass = &SentimentTrendClass
	}
    
	if ParticipantMetrics, ok := ConversationmetricsMap["participantMetrics"].(map[string]interface{}); ok {
		ParticipantMetricsString, _ := json.Marshal(ParticipantMetrics)
		json.Unmarshal(ParticipantMetricsString, &o.ParticipantMetrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
