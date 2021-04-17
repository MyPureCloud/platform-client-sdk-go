package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsconversation
type Analyticsconversation struct { 
	// ConversationEnd - The end time of a conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConversationEnd *time.Time `json:"conversationEnd,omitempty"`


	// ConversationId - Unique identifier for the conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// ConversationStart - The start time of a conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConversationStart *time.Time `json:"conversationStart,omitempty"`


	// DivisionIds - Identifier(s) of division(s) associated with a conversation
	DivisionIds *[]string `json:"divisionIds,omitempty"`


	// ExternalTag - External tag for the conversation
	ExternalTag *string `json:"externalTag,omitempty"`


	// MediaStatsMinConversationMos - The lowest estimated average MOS among all the audio streams belonging to this conversation
	MediaStatsMinConversationMos *float64 `json:"mediaStatsMinConversationMos,omitempty"`


	// MediaStatsMinConversationRFactor - The lowest R-factor value among all of the audio streams belonging to this conversation
	MediaStatsMinConversationRFactor *float64 `json:"mediaStatsMinConversationRFactor,omitempty"`


	// OriginatingDirection - The original direction of the conversation
	OriginatingDirection *string `json:"originatingDirection,omitempty"`


	// Evaluations - Evaluations associated with this conversation
	Evaluations *[]Analyticsevaluation `json:"evaluations,omitempty"`


	// Surveys - Surveys associated with this conversation
	Surveys *[]Analyticssurvey `json:"surveys,omitempty"`


	// Resolutions - Resolutions associated with this conversation
	Resolutions *[]Analyticsresolution `json:"resolutions,omitempty"`


	// Participants - Participants in the conversation
	Participants *[]Analyticsparticipant `json:"participants,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
