package platformclientv2
import (
	"time"
	"encoding/json"
)

// Analyticsconversationsegment
type Analyticsconversationsegment struct { 
	// SegmentStart - The timestamp when this segment began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SegmentStart *time.Time `json:"segmentStart,omitempty"`


	// SegmentEnd - The timestamp when this segment ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SegmentEnd *time.Time `json:"segmentEnd,omitempty"`


	// QueueId - Queue identifier
	QueueId *string `json:"queueId,omitempty"`


	// WrapUpCode - Wrapup Code id
	WrapUpCode *string `json:"wrapUpCode,omitempty"`


	// WrapUpNote - Note entered by an agent during after-call work
	WrapUpNote *string `json:"wrapUpNote,omitempty"`


	// WrapUpTags
	WrapUpTags *[]string `json:"wrapUpTags,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// DisconnectType - A description of the event that disconnected the segment
	DisconnectType *string `json:"disconnectType,omitempty"`


	// SegmentType - The activity taking place for the participant in the segment
	SegmentType *string `json:"segmentType,omitempty"`


	// RequestedRoutingUserIds
	RequestedRoutingUserIds *[]string `json:"requestedRoutingUserIds,omitempty"`


	// RequestedRoutingSkillIds
	RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`


	// RequestedLanguageId - A unique identifier for the language requested for an interaction.
	RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Analyticsscoredagent `json:"scoredAgents,omitempty"`


	// Properties
	Properties *[]Analyticsproperty `json:"properties,omitempty"`


	// SourceConversationId
	SourceConversationId *string `json:"sourceConversationId,omitempty"`


	// DestinationConversationId
	DestinationConversationId *string `json:"destinationConversationId,omitempty"`


	// SourceSessionId
	SourceSessionId *string `json:"sourceSessionId,omitempty"`


	// DestinationSessionId
	DestinationSessionId *string `json:"destinationSessionId,omitempty"`


	// SipResponseCodes
	SipResponseCodes *[]int64 `json:"sipResponseCodes,omitempty"`


	// Q850ResponseCodes
	Q850ResponseCodes *[]int64 `json:"q850ResponseCodes,omitempty"`


	// Conference - Indicates whether the segment was a conference
	Conference *bool `json:"conference,omitempty"`


	// GroupId
	GroupId *string `json:"groupId,omitempty"`


	// Subject
	Subject *string `json:"subject,omitempty"`


	// AudioMuted
	AudioMuted *bool `json:"audioMuted,omitempty"`


	// VideoMuted
	VideoMuted *bool `json:"videoMuted,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationsegment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
