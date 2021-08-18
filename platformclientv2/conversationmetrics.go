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

}

func (u *Conversationmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmetrics

	

	return json.Marshal(&struct { 
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		
		SentimentScore *float64 `json:"sentimentScore,omitempty"`
		
		SentimentTrend *float64 `json:"sentimentTrend,omitempty"`
		*Alias
	}{ 
		Conversation: u.Conversation,
		
		SentimentScore: u.SentimentScore,
		
		SentimentTrend: u.SentimentTrend,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
