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


	// ConversationInitiator - Indicates the participant purpose of the participant initiating a message conversation
	ConversationInitiator *string `json:"conversationInitiator,omitempty"`


	// ConversationStart - The start time of a conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConversationStart *time.Time `json:"conversationStart,omitempty"`


	// CustomerParticipation - Indicates a messaging conversation in which the customer participated by sending at least one message
	CustomerParticipation *bool `json:"customerParticipation,omitempty"`


	// DivisionIds - Identifier(s) of division(s) associated with a conversation
	DivisionIds *[]string `json:"divisionIds,omitempty"`


	// ExternalTag - External tag for the conversation
	ExternalTag *string `json:"externalTag,omitempty"`


	// KnowledgeBaseIds - The unique identifier(s) of the knowledge base(s) used
	KnowledgeBaseIds *[]string `json:"knowledgeBaseIds,omitempty"`


	// MediaStatsMinConversationMos - The lowest estimated average MOS among all the audio streams belonging to this conversation
	MediaStatsMinConversationMos *float64 `json:"mediaStatsMinConversationMos,omitempty"`


	// MediaStatsMinConversationRFactor - The lowest R-factor value among all of the audio streams belonging to this conversation
	MediaStatsMinConversationRFactor *float64 `json:"mediaStatsMinConversationRFactor,omitempty"`


	// OriginatingDirection - The original direction of the conversation
	OriginatingDirection *string `json:"originatingDirection,omitempty"`


	// SelfServed - Indicates whether all flow sessions were self serviced
	SelfServed *bool `json:"selfServed,omitempty"`


	// Evaluations - Evaluations associated with this conversation
	Evaluations *[]Analyticsevaluation `json:"evaluations,omitempty"`


	// Surveys - Surveys associated with this conversation
	Surveys *[]Analyticssurvey `json:"surveys,omitempty"`


	// Resolutions - Resolutions associated with this conversation
	Resolutions *[]Analyticsresolution `json:"resolutions,omitempty"`


	// Participants - Participants in the conversation
	Participants *[]Analyticsparticipantwithoutattributes `json:"participants,omitempty"`

}

func (o *Analyticsconversationwithoutattributes) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsconversationwithoutattributes
	
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
		
		ConversationInitiator *string `json:"conversationInitiator,omitempty"`
		
		ConversationStart *string `json:"conversationStart,omitempty"`
		
		CustomerParticipation *bool `json:"customerParticipation,omitempty"`
		
		DivisionIds *[]string `json:"divisionIds,omitempty"`
		
		ExternalTag *string `json:"externalTag,omitempty"`
		
		KnowledgeBaseIds *[]string `json:"knowledgeBaseIds,omitempty"`
		
		MediaStatsMinConversationMos *float64 `json:"mediaStatsMinConversationMos,omitempty"`
		
		MediaStatsMinConversationRFactor *float64 `json:"mediaStatsMinConversationRFactor,omitempty"`
		
		OriginatingDirection *string `json:"originatingDirection,omitempty"`
		
		SelfServed *bool `json:"selfServed,omitempty"`
		
		Evaluations *[]Analyticsevaluation `json:"evaluations,omitempty"`
		
		Surveys *[]Analyticssurvey `json:"surveys,omitempty"`
		
		Resolutions *[]Analyticsresolution `json:"resolutions,omitempty"`
		
		Participants *[]Analyticsparticipantwithoutattributes `json:"participants,omitempty"`
		*Alias
	}{ 
		ConversationEnd: ConversationEnd,
		
		ConversationId: o.ConversationId,
		
		ConversationInitiator: o.ConversationInitiator,
		
		ConversationStart: ConversationStart,
		
		CustomerParticipation: o.CustomerParticipation,
		
		DivisionIds: o.DivisionIds,
		
		ExternalTag: o.ExternalTag,
		
		KnowledgeBaseIds: o.KnowledgeBaseIds,
		
		MediaStatsMinConversationMos: o.MediaStatsMinConversationMos,
		
		MediaStatsMinConversationRFactor: o.MediaStatsMinConversationRFactor,
		
		OriginatingDirection: o.OriginatingDirection,
		
		SelfServed: o.SelfServed,
		
		Evaluations: o.Evaluations,
		
		Surveys: o.Surveys,
		
		Resolutions: o.Resolutions,
		
		Participants: o.Participants,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsconversationwithoutattributes) UnmarshalJSON(b []byte) error {
	var AnalyticsconversationwithoutattributesMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsconversationwithoutattributesMap)
	if err != nil {
		return err
	}
	
	if conversationEndString, ok := AnalyticsconversationwithoutattributesMap["conversationEnd"].(string); ok {
		ConversationEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", conversationEndString)
		o.ConversationEnd = &ConversationEnd
	}
	
	if ConversationId, ok := AnalyticsconversationwithoutattributesMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if ConversationInitiator, ok := AnalyticsconversationwithoutattributesMap["conversationInitiator"].(string); ok {
		o.ConversationInitiator = &ConversationInitiator
	}
	
	if conversationStartString, ok := AnalyticsconversationwithoutattributesMap["conversationStart"].(string); ok {
		ConversationStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", conversationStartString)
		o.ConversationStart = &ConversationStart
	}
	
	if CustomerParticipation, ok := AnalyticsconversationwithoutattributesMap["customerParticipation"].(bool); ok {
		o.CustomerParticipation = &CustomerParticipation
	}
	
	if DivisionIds, ok := AnalyticsconversationwithoutattributesMap["divisionIds"].([]interface{}); ok {
		DivisionIdsString, _ := json.Marshal(DivisionIds)
		json.Unmarshal(DivisionIdsString, &o.DivisionIds)
	}
	
	if ExternalTag, ok := AnalyticsconversationwithoutattributesMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
	
	if KnowledgeBaseIds, ok := AnalyticsconversationwithoutattributesMap["knowledgeBaseIds"].([]interface{}); ok {
		KnowledgeBaseIdsString, _ := json.Marshal(KnowledgeBaseIds)
		json.Unmarshal(KnowledgeBaseIdsString, &o.KnowledgeBaseIds)
	}
	
	if MediaStatsMinConversationMos, ok := AnalyticsconversationwithoutattributesMap["mediaStatsMinConversationMos"].(float64); ok {
		o.MediaStatsMinConversationMos = &MediaStatsMinConversationMos
	}
	
	if MediaStatsMinConversationRFactor, ok := AnalyticsconversationwithoutattributesMap["mediaStatsMinConversationRFactor"].(float64); ok {
		o.MediaStatsMinConversationRFactor = &MediaStatsMinConversationRFactor
	}
	
	if OriginatingDirection, ok := AnalyticsconversationwithoutattributesMap["originatingDirection"].(string); ok {
		o.OriginatingDirection = &OriginatingDirection
	}
	
	if SelfServed, ok := AnalyticsconversationwithoutattributesMap["selfServed"].(bool); ok {
		o.SelfServed = &SelfServed
	}
	
	if Evaluations, ok := AnalyticsconversationwithoutattributesMap["evaluations"].([]interface{}); ok {
		EvaluationsString, _ := json.Marshal(Evaluations)
		json.Unmarshal(EvaluationsString, &o.Evaluations)
	}
	
	if Surveys, ok := AnalyticsconversationwithoutattributesMap["surveys"].([]interface{}); ok {
		SurveysString, _ := json.Marshal(Surveys)
		json.Unmarshal(SurveysString, &o.Surveys)
	}
	
	if Resolutions, ok := AnalyticsconversationwithoutattributesMap["resolutions"].([]interface{}); ok {
		ResolutionsString, _ := json.Marshal(Resolutions)
		json.Unmarshal(ResolutionsString, &o.Resolutions)
	}
	
	if Participants, ok := AnalyticsconversationwithoutattributesMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsconversationwithoutattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
