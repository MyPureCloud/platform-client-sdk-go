package platformclientv2
import (
	"time"
	"encoding/json"
)

// Analyticsconversationwithoutattributes
type Analyticsconversationwithoutattributes struct { 
	// ConversationId - Unique identifier for the conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// ConversationStart - Date/time the conversation started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ConversationStart *time.Time `json:"conversationStart,omitempty"`


	// ConversationEnd - Date/time the conversation ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ConversationEnd *time.Time `json:"conversationEnd,omitempty"`


	// MediaStatsMinConversationMos - The lowest estimated average MOS among all the audio streams belonging to this conversation
	MediaStatsMinConversationMos *float64 `json:"mediaStatsMinConversationMos,omitempty"`


	// MediaStatsMinConversationRFactor - The lowest R-factor value among all of the audio streams belonging to this conversation
	MediaStatsMinConversationRFactor *float64 `json:"mediaStatsMinConversationRFactor,omitempty"`


	// OriginatingDirection - The original direction of the conversation
	OriginatingDirection *string `json:"originatingDirection,omitempty"`


	// Evaluations - Evaluations tied to this conversation
	Evaluations *[]Analyticsevaluation `json:"evaluations,omitempty"`


	// Surveys - Surveys tied to this conversation
	Surveys *[]Analyticssurvey `json:"surveys,omitempty"`


	// DivisionIds - Identifiers of divisions associated with this conversation
	DivisionIds *[]string `json:"divisionIds,omitempty"`


	// Participants - Participants in the conversation
	Participants *[]Analyticsparticipantwithoutattributes `json:"participants,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationwithoutattributes) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
