package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsconversationsegment
type Analyticsconversationsegment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AudioMuted - Flag indicating if audio is muted or not (true/false)
	AudioMuted *bool `json:"audioMuted,omitempty"`

	// Conference - Indicates whether the segment was a conference
	Conference *bool `json:"conference,omitempty"`

	// DestinationConversationId - The unique identifier of a new conversation when a conversation is ended for a conference
	DestinationConversationId *string `json:"destinationConversationId,omitempty"`

	// DestinationSessionId - The unique identifier of a new session when a session is ended for a conference
	DestinationSessionId *string `json:"destinationSessionId,omitempty"`

	// DisconnectType - The session disconnect type
	DisconnectType *string `json:"disconnectType,omitempty"`

	// ErrorCode - A code corresponding to the error that occurred
	ErrorCode *string `json:"errorCode,omitempty"`

	// GroupId - Unique identifier for a PureCloud group
	GroupId *string `json:"groupId,omitempty"`

	// Q850ResponseCodes - Q.850 response code(s)
	Q850ResponseCodes *[]int `json:"q850ResponseCodes,omitempty"`

	// QueueId - Queue identifier
	QueueId *string `json:"queueId,omitempty"`

	// RequestedLanguageId - Unique identifier for the language requested for an interaction
	RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`

	// RequestedRoutingSkillIds - Unique identifier(s) for skill(s) requested for an interaction
	RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`

	// RequestedRoutingUserIds - Unique identifier(s) for agent(s) requested for an interaction
	RequestedRoutingUserIds *[]string `json:"requestedRoutingUserIds,omitempty"`

	// SegmentEnd - The end time of a segment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SegmentEnd *time.Time `json:"segmentEnd,omitempty"`

	// SegmentStart - The start time of a segment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SegmentStart *time.Time `json:"segmentStart,omitempty"`

	// SegmentType - The activity that takes place in the segment, such as hold or interact
	SegmentType *string `json:"segmentType,omitempty"`

	// SipResponseCodes - SIP response code(s)
	SipResponseCodes *[]int `json:"sipResponseCodes,omitempty"`

	// SourceConversationId - The unique identifier of the previous conversation when a new conversation is created for a conference
	SourceConversationId *string `json:"sourceConversationId,omitempty"`

	// SourceSessionId - The unique identifier of the previous session when a new session is created for a conference
	SourceSessionId *string `json:"sourceSessionId,omitempty"`

	// Subject - The subject for the initial email that started this conversation
	Subject *string `json:"subject,omitempty"`

	// VideoMuted - Flag indicating if video is muted/paused or not (true/false)
	VideoMuted *bool `json:"videoMuted,omitempty"`

	// WrapUpCode - Wrap up code
	WrapUpCode *string `json:"wrapUpCode,omitempty"`

	// WrapUpNote - Note entered by an agent during after-call work
	WrapUpNote *string `json:"wrapUpNote,omitempty"`

	// WrapUpTags - Tag(s) assigned during after-call work
	WrapUpTags *[]string `json:"wrapUpTags,omitempty"`

	// ScoredAgents - Scored agents
	ScoredAgents *[]Analyticsscoredagent `json:"scoredAgents,omitempty"`

	// Properties - Additional segment properties
	Properties *[]Analyticsproperty `json:"properties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsconversationsegment) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsconversationsegment) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "SegmentEnd","SegmentStart", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Analyticsconversationsegment
	
	SegmentEnd := new(string)
	if o.SegmentEnd != nil {
		
		*SegmentEnd = timeutil.Strftime(o.SegmentEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SegmentEnd = nil
	}
	
	SegmentStart := new(string)
	if o.SegmentStart != nil {
		
		*SegmentStart = timeutil.Strftime(o.SegmentStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SegmentStart = nil
	}
	
	return json.Marshal(&struct { 
		AudioMuted *bool `json:"audioMuted,omitempty"`
		
		Conference *bool `json:"conference,omitempty"`
		
		DestinationConversationId *string `json:"destinationConversationId,omitempty"`
		
		DestinationSessionId *string `json:"destinationSessionId,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		GroupId *string `json:"groupId,omitempty"`
		
		Q850ResponseCodes *[]int `json:"q850ResponseCodes,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`
		
		RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`
		
		RequestedRoutingUserIds *[]string `json:"requestedRoutingUserIds,omitempty"`
		
		SegmentEnd *string `json:"segmentEnd,omitempty"`
		
		SegmentStart *string `json:"segmentStart,omitempty"`
		
		SegmentType *string `json:"segmentType,omitempty"`
		
		SipResponseCodes *[]int `json:"sipResponseCodes,omitempty"`
		
		SourceConversationId *string `json:"sourceConversationId,omitempty"`
		
		SourceSessionId *string `json:"sourceSessionId,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		VideoMuted *bool `json:"videoMuted,omitempty"`
		
		WrapUpCode *string `json:"wrapUpCode,omitempty"`
		
		WrapUpNote *string `json:"wrapUpNote,omitempty"`
		
		WrapUpTags *[]string `json:"wrapUpTags,omitempty"`
		
		ScoredAgents *[]Analyticsscoredagent `json:"scoredAgents,omitempty"`
		
		Properties *[]Analyticsproperty `json:"properties,omitempty"`
		Alias
	}{ 
		AudioMuted: o.AudioMuted,
		
		Conference: o.Conference,
		
		DestinationConversationId: o.DestinationConversationId,
		
		DestinationSessionId: o.DestinationSessionId,
		
		DisconnectType: o.DisconnectType,
		
		ErrorCode: o.ErrorCode,
		
		GroupId: o.GroupId,
		
		Q850ResponseCodes: o.Q850ResponseCodes,
		
		QueueId: o.QueueId,
		
		RequestedLanguageId: o.RequestedLanguageId,
		
		RequestedRoutingSkillIds: o.RequestedRoutingSkillIds,
		
		RequestedRoutingUserIds: o.RequestedRoutingUserIds,
		
		SegmentEnd: SegmentEnd,
		
		SegmentStart: SegmentStart,
		
		SegmentType: o.SegmentType,
		
		SipResponseCodes: o.SipResponseCodes,
		
		SourceConversationId: o.SourceConversationId,
		
		SourceSessionId: o.SourceSessionId,
		
		Subject: o.Subject,
		
		VideoMuted: o.VideoMuted,
		
		WrapUpCode: o.WrapUpCode,
		
		WrapUpNote: o.WrapUpNote,
		
		WrapUpTags: o.WrapUpTags,
		
		ScoredAgents: o.ScoredAgents,
		
		Properties: o.Properties,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsconversationsegment) UnmarshalJSON(b []byte) error {
	var AnalyticsconversationsegmentMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsconversationsegmentMap)
	if err != nil {
		return err
	}
	
	if AudioMuted, ok := AnalyticsconversationsegmentMap["audioMuted"].(bool); ok {
		o.AudioMuted = &AudioMuted
	}
    
	if Conference, ok := AnalyticsconversationsegmentMap["conference"].(bool); ok {
		o.Conference = &Conference
	}
    
	if DestinationConversationId, ok := AnalyticsconversationsegmentMap["destinationConversationId"].(string); ok {
		o.DestinationConversationId = &DestinationConversationId
	}
    
	if DestinationSessionId, ok := AnalyticsconversationsegmentMap["destinationSessionId"].(string); ok {
		o.DestinationSessionId = &DestinationSessionId
	}
    
	if DisconnectType, ok := AnalyticsconversationsegmentMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if ErrorCode, ok := AnalyticsconversationsegmentMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if GroupId, ok := AnalyticsconversationsegmentMap["groupId"].(string); ok {
		o.GroupId = &GroupId
	}
    
	if Q850ResponseCodes, ok := AnalyticsconversationsegmentMap["q850ResponseCodes"].([]interface{}); ok {
		Q850ResponseCodesString, _ := json.Marshal(Q850ResponseCodes)
		json.Unmarshal(Q850ResponseCodesString, &o.Q850ResponseCodes)
	}
	
	if QueueId, ok := AnalyticsconversationsegmentMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if RequestedLanguageId, ok := AnalyticsconversationsegmentMap["requestedLanguageId"].(string); ok {
		o.RequestedLanguageId = &RequestedLanguageId
	}
    
	if RequestedRoutingSkillIds, ok := AnalyticsconversationsegmentMap["requestedRoutingSkillIds"].([]interface{}); ok {
		RequestedRoutingSkillIdsString, _ := json.Marshal(RequestedRoutingSkillIds)
		json.Unmarshal(RequestedRoutingSkillIdsString, &o.RequestedRoutingSkillIds)
	}
	
	if RequestedRoutingUserIds, ok := AnalyticsconversationsegmentMap["requestedRoutingUserIds"].([]interface{}); ok {
		RequestedRoutingUserIdsString, _ := json.Marshal(RequestedRoutingUserIds)
		json.Unmarshal(RequestedRoutingUserIdsString, &o.RequestedRoutingUserIds)
	}
	
	if segmentEndString, ok := AnalyticsconversationsegmentMap["segmentEnd"].(string); ok {
		SegmentEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", segmentEndString)
		o.SegmentEnd = &SegmentEnd
	}
	
	if segmentStartString, ok := AnalyticsconversationsegmentMap["segmentStart"].(string); ok {
		SegmentStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", segmentStartString)
		o.SegmentStart = &SegmentStart
	}
	
	if SegmentType, ok := AnalyticsconversationsegmentMap["segmentType"].(string); ok {
		o.SegmentType = &SegmentType
	}
    
	if SipResponseCodes, ok := AnalyticsconversationsegmentMap["sipResponseCodes"].([]interface{}); ok {
		SipResponseCodesString, _ := json.Marshal(SipResponseCodes)
		json.Unmarshal(SipResponseCodesString, &o.SipResponseCodes)
	}
	
	if SourceConversationId, ok := AnalyticsconversationsegmentMap["sourceConversationId"].(string); ok {
		o.SourceConversationId = &SourceConversationId
	}
    
	if SourceSessionId, ok := AnalyticsconversationsegmentMap["sourceSessionId"].(string); ok {
		o.SourceSessionId = &SourceSessionId
	}
    
	if Subject, ok := AnalyticsconversationsegmentMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if VideoMuted, ok := AnalyticsconversationsegmentMap["videoMuted"].(bool); ok {
		o.VideoMuted = &VideoMuted
	}
    
	if WrapUpCode, ok := AnalyticsconversationsegmentMap["wrapUpCode"].(string); ok {
		o.WrapUpCode = &WrapUpCode
	}
    
	if WrapUpNote, ok := AnalyticsconversationsegmentMap["wrapUpNote"].(string); ok {
		o.WrapUpNote = &WrapUpNote
	}
    
	if WrapUpTags, ok := AnalyticsconversationsegmentMap["wrapUpTags"].([]interface{}); ok {
		WrapUpTagsString, _ := json.Marshal(WrapUpTags)
		json.Unmarshal(WrapUpTagsString, &o.WrapUpTags)
	}
	
	if ScoredAgents, ok := AnalyticsconversationsegmentMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	
	if Properties, ok := AnalyticsconversationsegmentMap["properties"].([]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsconversationsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
