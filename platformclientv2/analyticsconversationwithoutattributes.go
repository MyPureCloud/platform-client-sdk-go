package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsconversationwithoutattributes
type Analyticsconversationwithoutattributes struct { 
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
	Participants *[]Analyticsparticipantwithoutattributes `json:"participants,omitempty"`

}

func (u *Analyticsconversationwithoutattributes) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsconversationwithoutattributes

	
	ConversationEnd := new(string)
	if u.ConversationEnd != nil {
		
		*ConversationEnd = timeutil.Strftime(u.ConversationEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConversationEnd = nil
	}
	
	ConversationStart := new(string)
	if u.ConversationStart != nil {
		
		*ConversationStart = timeutil.Strftime(u.ConversationStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConversationStart = nil
	}
	

	return json.Marshal(&struct { 
		ConversationEnd *string `json:"conversationEnd,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ConversationStart *string `json:"conversationStart,omitempty"`
		
		DivisionIds *[]string `json:"divisionIds,omitempty"`
		
		ExternalTag *string `json:"externalTag,omitempty"`
		
		MediaStatsMinConversationMos *float64 `json:"mediaStatsMinConversationMos,omitempty"`
		
		MediaStatsMinConversationRFactor *float64 `json:"mediaStatsMinConversationRFactor,omitempty"`
		
		OriginatingDirection *string `json:"originatingDirection,omitempty"`
		
		Evaluations *[]Analyticsevaluation `json:"evaluations,omitempty"`
		
		Surveys *[]Analyticssurvey `json:"surveys,omitempty"`
		
		Resolutions *[]Analyticsresolution `json:"resolutions,omitempty"`
		
		Participants *[]Analyticsparticipantwithoutattributes `json:"participants,omitempty"`
		*Alias
	}{ 
		ConversationEnd: ConversationEnd,
		
		ConversationId: u.ConversationId,
		
		ConversationStart: ConversationStart,
		
		DivisionIds: u.DivisionIds,
		
		ExternalTag: u.ExternalTag,
		
		MediaStatsMinConversationMos: u.MediaStatsMinConversationMos,
		
		MediaStatsMinConversationRFactor: u.MediaStatsMinConversationRFactor,
		
		OriginatingDirection: u.OriginatingDirection,
		
		Evaluations: u.Evaluations,
		
		Surveys: u.Surveys,
		
		Resolutions: u.Resolutions,
		
		Participants: u.Participants,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticsconversationwithoutattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
