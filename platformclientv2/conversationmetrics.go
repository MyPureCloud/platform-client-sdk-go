package platformclientv2
import (
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

}

// String returns a JSON representation of the model
func (o *Conversationmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
