package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmetrics
type Conversationmetrics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Conversation - The Conversation Reference
	Conversation *Addressableentityref `json:"conversation,omitempty"`

	// SentimentScore - The Sentiment Score
	SentimentScore *float64 `json:"sentimentScore,omitempty"`

	// SentimentTrend - The Sentiment Trend
	SentimentTrend *float64 `json:"sentimentTrend,omitempty"`

	// SentimentTrendClass - The Sentiment Trend Class
	SentimentTrendClass *string `json:"sentimentTrendClass,omitempty"`

	// EmpathyScores - The Empathy Scores
	EmpathyScores *[]Empathyscore `json:"empathyScores,omitempty"`

	// ParticipantMetrics - The Participant Metrics
	ParticipantMetrics *Participantmetrics `json:"participantMetrics,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationmetrics) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Conversationmetrics) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationmetrics
	
	return json.Marshal(&struct { 
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		
		SentimentScore *float64 `json:"sentimentScore,omitempty"`
		
		SentimentTrend *float64 `json:"sentimentTrend,omitempty"`
		
		SentimentTrendClass *string `json:"sentimentTrendClass,omitempty"`
		
		EmpathyScores *[]Empathyscore `json:"empathyScores,omitempty"`
		
		ParticipantMetrics *Participantmetrics `json:"participantMetrics,omitempty"`
		Alias
	}{ 
		Conversation: o.Conversation,
		
		SentimentScore: o.SentimentScore,
		
		SentimentTrend: o.SentimentTrend,
		
		SentimentTrendClass: o.SentimentTrendClass,
		
		EmpathyScores: o.EmpathyScores,
		
		ParticipantMetrics: o.ParticipantMetrics,
		Alias:    (Alias)(o),
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
    
	if EmpathyScores, ok := ConversationmetricsMap["empathyScores"].([]interface{}); ok {
		EmpathyScoresString, _ := json.Marshal(EmpathyScores)
		json.Unmarshal(EmpathyScoresString, &o.EmpathyScores)
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
