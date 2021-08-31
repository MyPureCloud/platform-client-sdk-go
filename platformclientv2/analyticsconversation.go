package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (o *Analyticsconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsconversation
	
	ConversationEnd := new(string)
	if o.ConversationEnd != nil {
		
		*ConversationEnd = timeutil.Strftime(o.ConversationEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConversationEnd = nil
	}
	
	ConversationStart := new(string)
	if o.ConversationStart != nil {
		
		*ConversationStart = timeutil.Strftime(o.ConversationStart, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		Participants *[]Analyticsparticipant `json:"participants,omitempty"`
		*Alias
	}{ 
		ConversationEnd: ConversationEnd,
		
		ConversationId: o.ConversationId,
		
		ConversationStart: ConversationStart,
		
		DivisionIds: o.DivisionIds,
		
		ExternalTag: o.ExternalTag,
		
		MediaStatsMinConversationMos: o.MediaStatsMinConversationMos,
		
		MediaStatsMinConversationRFactor: o.MediaStatsMinConversationRFactor,
		
		OriginatingDirection: o.OriginatingDirection,
		
		Evaluations: o.Evaluations,
		
		Surveys: o.Surveys,
		
		Resolutions: o.Resolutions,
		
		Participants: o.Participants,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsconversation) UnmarshalJSON(b []byte) error {
	var AnalyticsconversationMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsconversationMap)
	if err != nil {
		return err
	}
	
	if conversationEndString, ok := AnalyticsconversationMap["conversationEnd"].(string); ok {
		ConversationEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", conversationEndString)
		o.ConversationEnd = &ConversationEnd
	}
	
	if ConversationId, ok := AnalyticsconversationMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if conversationStartString, ok := AnalyticsconversationMap["conversationStart"].(string); ok {
		ConversationStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", conversationStartString)
		o.ConversationStart = &ConversationStart
	}
	
	if DivisionIds, ok := AnalyticsconversationMap["divisionIds"].([]interface{}); ok {
		DivisionIdsString, _ := json.Marshal(DivisionIds)
		json.Unmarshal(DivisionIdsString, &o.DivisionIds)
	}
	
	if ExternalTag, ok := AnalyticsconversationMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
	
	if MediaStatsMinConversationMos, ok := AnalyticsconversationMap["mediaStatsMinConversationMos"].(float64); ok {
		o.MediaStatsMinConversationMos = &MediaStatsMinConversationMos
	}
	
	if MediaStatsMinConversationRFactor, ok := AnalyticsconversationMap["mediaStatsMinConversationRFactor"].(float64); ok {
		o.MediaStatsMinConversationRFactor = &MediaStatsMinConversationRFactor
	}
	
	if OriginatingDirection, ok := AnalyticsconversationMap["originatingDirection"].(string); ok {
		o.OriginatingDirection = &OriginatingDirection
	}
	
	if Evaluations, ok := AnalyticsconversationMap["evaluations"].([]interface{}); ok {
		EvaluationsString, _ := json.Marshal(Evaluations)
		json.Unmarshal(EvaluationsString, &o.Evaluations)
	}
	
	if Surveys, ok := AnalyticsconversationMap["surveys"].([]interface{}); ok {
		SurveysString, _ := json.Marshal(Surveys)
		json.Unmarshal(SurveysString, &o.Surveys)
	}
	
	if Resolutions, ok := AnalyticsconversationMap["resolutions"].([]interface{}); ok {
		ResolutionsString, _ := json.Marshal(Resolutions)
		json.Unmarshal(ResolutionsString, &o.Resolutions)
	}
	
	if Participants, ok := AnalyticsconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
