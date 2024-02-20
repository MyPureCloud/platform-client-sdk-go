package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationmetricstopicconversationmetricrecord
type Conversationmetricstopicconversationmetricrecord struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Metric - Metric name
	Metric *string `json:"metric,omitempty"`

	// MetricDate - The date and time of metric creation
	MetricDate *time.Time `json:"metricDate,omitempty"`

	// Value - Metric value
	Value *int `json:"value,omitempty"`

	// RecordId - Record identifier
	RecordId *string `json:"recordId,omitempty"`

	// ActiveRouting - Active routing method
	ActiveRouting *string `json:"activeRouting,omitempty"`

	// ActiveSkillIds - ID(s) of Skill(s) that are active on the conversation
	ActiveSkillIds *[]string `json:"activeSkillIds,omitempty"`

	// AddressFrom - The address that initiated an action
	AddressFrom *string `json:"addressFrom,omitempty"`

	// AddressTo - The address receiving an action
	AddressTo *string `json:"addressTo,omitempty"`

	// AgentAssistantId - Unique identifier of the active virtual agent assistant
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`

	// AgentBullseyeRing - Bullseye ring of the targeted agent
	AgentBullseyeRing *int `json:"agentBullseyeRing,omitempty"`

	// AgentOwned - Flag indicating an agent-owned callback
	AgentOwned *bool `json:"agentOwned,omitempty"`

	// Ani - Automatic Number Identification (caller's number)
	Ani *string `json:"ani,omitempty"`

	// AssignerId - ID of the user that manually assigned a conversation
	AssignerId *string `json:"assignerId,omitempty"`

	// Authenticated - Flag that indicates that the identity of the customer has been asserted as verified by the provider.
	Authenticated *bool `json:"authenticated,omitempty"`

	// ConversationId - Unique identifier for the conversation
	ConversationId *string `json:"conversationId,omitempty"`

	// ConversationInitiator - Indicates the participant purpose of the participant initiating a message conversation
	ConversationInitiator *string `json:"conversationInitiator,omitempty"`

	// ConvertedFrom - Session media type that was converted from in case of a media type conversion
	ConvertedFrom *string `json:"convertedFrom,omitempty"`

	// ConvertedTo - Session media type that was converted to in case of a media type conversion
	ConvertedTo *string `json:"convertedTo,omitempty"`

	// CustomerParticipation - Indicates a messaging conversation in which the customer participated by sending at least one message
	CustomerParticipation *bool `json:"customerParticipation,omitempty"`

	// DeliveryStatus - The email or SMS delivery status
	DeliveryStatus *string `json:"deliveryStatus,omitempty"`

	// DestinationAddresses - Destination address(es) of transfers or consults
	DestinationAddresses *[]string `json:"destinationAddresses,omitempty"`

	// Direction - The direction of the communication
	Direction *string `json:"direction,omitempty"`

	// DisconnectType - The session disconnect type
	DisconnectType *string `json:"disconnectType,omitempty"`

	// DivisionIds - Identifier(s) of division(s) associated with a conversation
	DivisionIds *[]string `json:"divisionIds,omitempty"`

	// Dnis - Dialed number identification service (number dialed by the calling party)
	Dnis *string `json:"dnis,omitempty"`

	// EdgeId - Unique identifier of the edge device
	EdgeId *string `json:"edgeId,omitempty"`

	// EligibleAgentCounts - Number of eligible agents for each predictive routing attempt
	EligibleAgentCounts *[]int `json:"eligibleAgentCounts,omitempty"`

	// ErrorCode - A code corresponding to the error that occurred
	ErrorCode *string `json:"errorCode,omitempty"`

	// ExtendedDeliveryStatus - Extended delivery status
	ExtendedDeliveryStatus *string `json:"extendedDeliveryStatus,omitempty"`

	// ExternalContactId - External contact identifier
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// ExternalMediaCount - Count of any media (images, files, etc) included on the external session
	ExternalMediaCount *int `json:"externalMediaCount,omitempty"`

	// ExternalOrganizationId - External organization identifier
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`

	// ExternalTag - External tag for the conversation
	ExternalTag *string `json:"externalTag,omitempty"`

	// FirstQueue - Marker that is set if the current queue is the first queue in a conversation
	FirstQueue *bool `json:"firstQueue,omitempty"`

	// FlaggedReason - Reason for which participant flagged conversation
	FlaggedReason *string `json:"flaggedReason,omitempty"`

	// FlowInType - Type of flow in that occurred when entering ACD.
	FlowInType *string `json:"flowInType,omitempty"`

	// FlowOutType - Type of flow out that occurred when emitting tFlowOut.
	FlowOutType *string `json:"flowOutType,omitempty"`

	// GroupId - Unique identifier for a PureCloud group
	GroupId *string `json:"groupId,omitempty"`

	// InteractionType - The interaction type (enterprise or contactCenter)
	InteractionType *string `json:"interactionType,omitempty"`

	// JourneyActionId - Identifier of the journey action.
	JourneyActionId *string `json:"journeyActionId,omitempty"`

	// JourneyActionMapId - Identifier of the journey action map that triggered the action.
	JourneyActionMapId *string `json:"journeyActionMapId,omitempty"`

	// JourneyActionMapVersion - Version of the journey action map that triggered the action.
	JourneyActionMapVersion *int `json:"journeyActionMapVersion,omitempty"`

	// JourneyCustomerId - Primary identifier of the journey customer in the source where the activities originate from.
	JourneyCustomerId *string `json:"journeyCustomerId,omitempty"`

	// JourneyCustomerIdType - Type of primary identifier of the journey customer (e.g. cookie).
	JourneyCustomerIdType *string `json:"journeyCustomerIdType,omitempty"`

	// JourneyCustomerSessionId - Unique identifier of the journey session.
	JourneyCustomerSessionId *string `json:"journeyCustomerSessionId,omitempty"`

	// JourneyCustomerSessionIdType - Type or category of journey sessions (e.g. web, ticket, delivery, atm).
	JourneyCustomerSessionIdType *string `json:"journeyCustomerSessionIdType,omitempty"`

	// KnowledgeBaseIds - The unique identifier(s) of the knowledge base(s) used
	KnowledgeBaseIds *[]string `json:"knowledgeBaseIds,omitempty"`

	// MediaCount - Count of any media (images, files, etc) included in this session
	MediaCount *int `json:"mediaCount,omitempty"`

	// MediaType - The session media type
	MediaType *string `json:"mediaType,omitempty"`

	// MessageType - Message type for messaging services. E.g.: sms, facebook, twitter, line
	MessageType *string `json:"messageType,omitempty"`

	// OriginatingDirection - The original direction of the conversation
	OriginatingDirection *string `json:"originatingDirection,omitempty"`

	// OutboundCampaignId - (Dialer) Unique identifier of the outbound campaign
	OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`

	// OutboundContactId - (Dialer) Unique identifier of the contact
	OutboundContactId *string `json:"outboundContactId,omitempty"`

	// OutboundContactListId - (Dialer) Unique identifier of the contact list that this contact belongs to
	OutboundContactListId *string `json:"outboundContactListId,omitempty"`

	// ParticipantName - A human readable name identifying the participant
	ParticipantName *string `json:"participantName,omitempty"`

	// PeerId - This identifies pairs of related sessions on a conversation. E.g. an external session’s peerId will be the session that the call originally connected to, e.g. if an IVR was dialed, the IVR session, which will also have the external session’s ID as its peer. After that point, any transfers of that session to other internal components (acd, agent, etc.) will all spawn new sessions whose peerIds point back to that original external session.
	PeerId *string `json:"peerId,omitempty"`

	// Provider - The source provider for the communication.
	Provider *string `json:"provider,omitempty"`

	// Purpose - The participant's purpose
	Purpose *string `json:"purpose,omitempty"`

	// QueueId - Queue identifier
	QueueId *string `json:"queueId,omitempty"`

	// Remote - Name, phone number, or email address of the remote party.
	Remote *string `json:"remote,omitempty"`

	// RemovedSkillIds - ID(s) of Skill(s) that have been removed by bullseye routing
	RemovedSkillIds *[]string `json:"removedSkillIds,omitempty"`

	// RequestedLanguageId - Unique identifier for the language requested for an interaction
	RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`

	// RequestedRoutingSkillIds - Unique identifier(s) for skill(s) requested for an interaction
	RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`

	// RequestedRoutings - Routing type(s) for requested/attempted routing methods.
	RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`

	// RoomId - Unique identifier for the room
	RoomId *string `json:"roomId,omitempty"`

	// RoutingPriority - Routing priority for the current interaction
	RoutingPriority *int `json:"routingPriority,omitempty"`

	// RoutingRing - Routing ring for bullseye or preferred agent routing
	RoutingRing *int `json:"routingRing,omitempty"`

	// RoutingRule - Routing rule for preferred, conditional and predictive routing type
	RoutingRule *string `json:"routingRule,omitempty"`

	// RoutingRuleType - Routing rule type
	RoutingRuleType *string `json:"routingRuleType,omitempty"`

	// SelectedAgentId - Selected agent ID
	SelectedAgentId *string `json:"selectedAgentId,omitempty"`

	// SelectedAgentRank - Selected agent GPR rank
	SelectedAgentRank *int `json:"selectedAgentRank,omitempty"`

	// SelfServed - Indicates whether all flow sessions were self serviced
	SelfServed *bool `json:"selfServed,omitempty"`

	// SessionDnis - Dialed number for the current session; this can be different from dnis, e.g. if the call was transferred
	SessionDnis *string `json:"sessionDnis,omitempty"`

	// SessionId - The unique identifier of this session
	SessionId *string `json:"sessionId,omitempty"`

	// StationId - Unique identifier for a phone
	StationId *string `json:"stationId,omitempty"`

	// TeamId - The team ID the user is a member of
	TeamId *string `json:"teamId,omitempty"`

	// UsedRouting - Complete routing method
	UsedRouting *string `json:"usedRouting,omitempty"`

	// UserId - Unique identifier for the user
	UserId *string `json:"userId,omitempty"`

	// WaitingInteractionCounts - Number of waiting interactions for each predictive routing attempt
	WaitingInteractionCounts *[]int `json:"waitingInteractionCounts,omitempty"`

	// WrapUpCode - Wrap up code
	WrapUpCode *string `json:"wrapUpCode,omitempty"`

	// ProposedAgents - Proposed agents
	ProposedAgents *[]Conversationmetricstopicconversationproposedagent `json:"proposedAgents,omitempty"`

	// ScoredAgents - Scored agents
	ScoredAgents *[]Conversationmetricstopicconversationscoredagent `json:"scoredAgents,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationmetricstopicconversationmetricrecord) SetField(field string, fieldValue interface{}) {
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

func (o Conversationmetricstopicconversationmetricrecord) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "MetricDate", }
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
	type Alias Conversationmetricstopicconversationmetricrecord
	
	MetricDate := new(string)
	if o.MetricDate != nil {
		
		*MetricDate = timeutil.Strftime(o.MetricDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		MetricDate = nil
	}
	
	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		MetricDate *string `json:"metricDate,omitempty"`
		
		Value *int `json:"value,omitempty"`
		
		RecordId *string `json:"recordId,omitempty"`
		
		ActiveRouting *string `json:"activeRouting,omitempty"`
		
		ActiveSkillIds *[]string `json:"activeSkillIds,omitempty"`
		
		AddressFrom *string `json:"addressFrom,omitempty"`
		
		AddressTo *string `json:"addressTo,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		
		AgentBullseyeRing *int `json:"agentBullseyeRing,omitempty"`
		
		AgentOwned *bool `json:"agentOwned,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		AssignerId *string `json:"assignerId,omitempty"`
		
		Authenticated *bool `json:"authenticated,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ConversationInitiator *string `json:"conversationInitiator,omitempty"`
		
		ConvertedFrom *string `json:"convertedFrom,omitempty"`
		
		ConvertedTo *string `json:"convertedTo,omitempty"`
		
		CustomerParticipation *bool `json:"customerParticipation,omitempty"`
		
		DeliveryStatus *string `json:"deliveryStatus,omitempty"`
		
		DestinationAddresses *[]string `json:"destinationAddresses,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		DivisionIds *[]string `json:"divisionIds,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		EdgeId *string `json:"edgeId,omitempty"`
		
		EligibleAgentCounts *[]int `json:"eligibleAgentCounts,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ExtendedDeliveryStatus *string `json:"extendedDeliveryStatus,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ExternalMediaCount *int `json:"externalMediaCount,omitempty"`
		
		ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`
		
		ExternalTag *string `json:"externalTag,omitempty"`
		
		FirstQueue *bool `json:"firstQueue,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		FlowInType *string `json:"flowInType,omitempty"`
		
		FlowOutType *string `json:"flowOutType,omitempty"`
		
		GroupId *string `json:"groupId,omitempty"`
		
		InteractionType *string `json:"interactionType,omitempty"`
		
		JourneyActionId *string `json:"journeyActionId,omitempty"`
		
		JourneyActionMapId *string `json:"journeyActionMapId,omitempty"`
		
		JourneyActionMapVersion *int `json:"journeyActionMapVersion,omitempty"`
		
		JourneyCustomerId *string `json:"journeyCustomerId,omitempty"`
		
		JourneyCustomerIdType *string `json:"journeyCustomerIdType,omitempty"`
		
		JourneyCustomerSessionId *string `json:"journeyCustomerSessionId,omitempty"`
		
		JourneyCustomerSessionIdType *string `json:"journeyCustomerSessionIdType,omitempty"`
		
		KnowledgeBaseIds *[]string `json:"knowledgeBaseIds,omitempty"`
		
		MediaCount *int `json:"mediaCount,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		OriginatingDirection *string `json:"originatingDirection,omitempty"`
		
		OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`
		
		OutboundContactId *string `json:"outboundContactId,omitempty"`
		
		OutboundContactListId *string `json:"outboundContactListId,omitempty"`
		
		ParticipantName *string `json:"participantName,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		Purpose *string `json:"purpose,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		Remote *string `json:"remote,omitempty"`
		
		RemovedSkillIds *[]string `json:"removedSkillIds,omitempty"`
		
		RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`
		
		RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`
		
		RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`
		
		RoomId *string `json:"roomId,omitempty"`
		
		RoutingPriority *int `json:"routingPriority,omitempty"`
		
		RoutingRing *int `json:"routingRing,omitempty"`
		
		RoutingRule *string `json:"routingRule,omitempty"`
		
		RoutingRuleType *string `json:"routingRuleType,omitempty"`
		
		SelectedAgentId *string `json:"selectedAgentId,omitempty"`
		
		SelectedAgentRank *int `json:"selectedAgentRank,omitempty"`
		
		SelfServed *bool `json:"selfServed,omitempty"`
		
		SessionDnis *string `json:"sessionDnis,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		StationId *string `json:"stationId,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		UsedRouting *string `json:"usedRouting,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		WaitingInteractionCounts *[]int `json:"waitingInteractionCounts,omitempty"`
		
		WrapUpCode *string `json:"wrapUpCode,omitempty"`
		
		ProposedAgents *[]Conversationmetricstopicconversationproposedagent `json:"proposedAgents,omitempty"`
		
		ScoredAgents *[]Conversationmetricstopicconversationscoredagent `json:"scoredAgents,omitempty"`
		Alias
	}{ 
		Metric: o.Metric,
		
		MetricDate: MetricDate,
		
		Value: o.Value,
		
		RecordId: o.RecordId,
		
		ActiveRouting: o.ActiveRouting,
		
		ActiveSkillIds: o.ActiveSkillIds,
		
		AddressFrom: o.AddressFrom,
		
		AddressTo: o.AddressTo,
		
		AgentAssistantId: o.AgentAssistantId,
		
		AgentBullseyeRing: o.AgentBullseyeRing,
		
		AgentOwned: o.AgentOwned,
		
		Ani: o.Ani,
		
		AssignerId: o.AssignerId,
		
		Authenticated: o.Authenticated,
		
		ConversationId: o.ConversationId,
		
		ConversationInitiator: o.ConversationInitiator,
		
		ConvertedFrom: o.ConvertedFrom,
		
		ConvertedTo: o.ConvertedTo,
		
		CustomerParticipation: o.CustomerParticipation,
		
		DeliveryStatus: o.DeliveryStatus,
		
		DestinationAddresses: o.DestinationAddresses,
		
		Direction: o.Direction,
		
		DisconnectType: o.DisconnectType,
		
		DivisionIds: o.DivisionIds,
		
		Dnis: o.Dnis,
		
		EdgeId: o.EdgeId,
		
		EligibleAgentCounts: o.EligibleAgentCounts,
		
		ErrorCode: o.ErrorCode,
		
		ExtendedDeliveryStatus: o.ExtendedDeliveryStatus,
		
		ExternalContactId: o.ExternalContactId,
		
		ExternalMediaCount: o.ExternalMediaCount,
		
		ExternalOrganizationId: o.ExternalOrganizationId,
		
		ExternalTag: o.ExternalTag,
		
		FirstQueue: o.FirstQueue,
		
		FlaggedReason: o.FlaggedReason,
		
		FlowInType: o.FlowInType,
		
		FlowOutType: o.FlowOutType,
		
		GroupId: o.GroupId,
		
		InteractionType: o.InteractionType,
		
		JourneyActionId: o.JourneyActionId,
		
		JourneyActionMapId: o.JourneyActionMapId,
		
		JourneyActionMapVersion: o.JourneyActionMapVersion,
		
		JourneyCustomerId: o.JourneyCustomerId,
		
		JourneyCustomerIdType: o.JourneyCustomerIdType,
		
		JourneyCustomerSessionId: o.JourneyCustomerSessionId,
		
		JourneyCustomerSessionIdType: o.JourneyCustomerSessionIdType,
		
		KnowledgeBaseIds: o.KnowledgeBaseIds,
		
		MediaCount: o.MediaCount,
		
		MediaType: o.MediaType,
		
		MessageType: o.MessageType,
		
		OriginatingDirection: o.OriginatingDirection,
		
		OutboundCampaignId: o.OutboundCampaignId,
		
		OutboundContactId: o.OutboundContactId,
		
		OutboundContactListId: o.OutboundContactListId,
		
		ParticipantName: o.ParticipantName,
		
		PeerId: o.PeerId,
		
		Provider: o.Provider,
		
		Purpose: o.Purpose,
		
		QueueId: o.QueueId,
		
		Remote: o.Remote,
		
		RemovedSkillIds: o.RemovedSkillIds,
		
		RequestedLanguageId: o.RequestedLanguageId,
		
		RequestedRoutingSkillIds: o.RequestedRoutingSkillIds,
		
		RequestedRoutings: o.RequestedRoutings,
		
		RoomId: o.RoomId,
		
		RoutingPriority: o.RoutingPriority,
		
		RoutingRing: o.RoutingRing,
		
		RoutingRule: o.RoutingRule,
		
		RoutingRuleType: o.RoutingRuleType,
		
		SelectedAgentId: o.SelectedAgentId,
		
		SelectedAgentRank: o.SelectedAgentRank,
		
		SelfServed: o.SelfServed,
		
		SessionDnis: o.SessionDnis,
		
		SessionId: o.SessionId,
		
		StationId: o.StationId,
		
		TeamId: o.TeamId,
		
		UsedRouting: o.UsedRouting,
		
		UserId: o.UserId,
		
		WaitingInteractionCounts: o.WaitingInteractionCounts,
		
		WrapUpCode: o.WrapUpCode,
		
		ProposedAgents: o.ProposedAgents,
		
		ScoredAgents: o.ScoredAgents,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationmetricstopicconversationmetricrecord) UnmarshalJSON(b []byte) error {
	var ConversationmetricstopicconversationmetricrecordMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationmetricstopicconversationmetricrecordMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := ConversationmetricstopicconversationmetricrecordMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if metricDateString, ok := ConversationmetricstopicconversationmetricrecordMap["metricDate"].(string); ok {
		MetricDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", metricDateString)
		o.MetricDate = &MetricDate
	}
	
	if Value, ok := ConversationmetricstopicconversationmetricrecordMap["value"].(float64); ok {
		ValueInt := int(Value)
		o.Value = &ValueInt
	}
	
	if RecordId, ok := ConversationmetricstopicconversationmetricrecordMap["recordId"].(string); ok {
		o.RecordId = &RecordId
	}
    
	if ActiveRouting, ok := ConversationmetricstopicconversationmetricrecordMap["activeRouting"].(string); ok {
		o.ActiveRouting = &ActiveRouting
	}
    
	if ActiveSkillIds, ok := ConversationmetricstopicconversationmetricrecordMap["activeSkillIds"].([]interface{}); ok {
		ActiveSkillIdsString, _ := json.Marshal(ActiveSkillIds)
		json.Unmarshal(ActiveSkillIdsString, &o.ActiveSkillIds)
	}
	
	if AddressFrom, ok := ConversationmetricstopicconversationmetricrecordMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if AddressTo, ok := ConversationmetricstopicconversationmetricrecordMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if AgentAssistantId, ok := ConversationmetricstopicconversationmetricrecordMap["agentAssistantId"].(string); ok {
		o.AgentAssistantId = &AgentAssistantId
	}
    
	if AgentBullseyeRing, ok := ConversationmetricstopicconversationmetricrecordMap["agentBullseyeRing"].(float64); ok {
		AgentBullseyeRingInt := int(AgentBullseyeRing)
		o.AgentBullseyeRing = &AgentBullseyeRingInt
	}
	
	if AgentOwned, ok := ConversationmetricstopicconversationmetricrecordMap["agentOwned"].(bool); ok {
		o.AgentOwned = &AgentOwned
	}
    
	if Ani, ok := ConversationmetricstopicconversationmetricrecordMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if AssignerId, ok := ConversationmetricstopicconversationmetricrecordMap["assignerId"].(string); ok {
		o.AssignerId = &AssignerId
	}
    
	if Authenticated, ok := ConversationmetricstopicconversationmetricrecordMap["authenticated"].(bool); ok {
		o.Authenticated = &Authenticated
	}
    
	if ConversationId, ok := ConversationmetricstopicconversationmetricrecordMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ConversationInitiator, ok := ConversationmetricstopicconversationmetricrecordMap["conversationInitiator"].(string); ok {
		o.ConversationInitiator = &ConversationInitiator
	}
    
	if ConvertedFrom, ok := ConversationmetricstopicconversationmetricrecordMap["convertedFrom"].(string); ok {
		o.ConvertedFrom = &ConvertedFrom
	}
    
	if ConvertedTo, ok := ConversationmetricstopicconversationmetricrecordMap["convertedTo"].(string); ok {
		o.ConvertedTo = &ConvertedTo
	}
    
	if CustomerParticipation, ok := ConversationmetricstopicconversationmetricrecordMap["customerParticipation"].(bool); ok {
		o.CustomerParticipation = &CustomerParticipation
	}
    
	if DeliveryStatus, ok := ConversationmetricstopicconversationmetricrecordMap["deliveryStatus"].(string); ok {
		o.DeliveryStatus = &DeliveryStatus
	}
    
	if DestinationAddresses, ok := ConversationmetricstopicconversationmetricrecordMap["destinationAddresses"].([]interface{}); ok {
		DestinationAddressesString, _ := json.Marshal(DestinationAddresses)
		json.Unmarshal(DestinationAddressesString, &o.DestinationAddresses)
	}
	
	if Direction, ok := ConversationmetricstopicconversationmetricrecordMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if DisconnectType, ok := ConversationmetricstopicconversationmetricrecordMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if DivisionIds, ok := ConversationmetricstopicconversationmetricrecordMap["divisionIds"].([]interface{}); ok {
		DivisionIdsString, _ := json.Marshal(DivisionIds)
		json.Unmarshal(DivisionIdsString, &o.DivisionIds)
	}
	
	if Dnis, ok := ConversationmetricstopicconversationmetricrecordMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if EdgeId, ok := ConversationmetricstopicconversationmetricrecordMap["edgeId"].(string); ok {
		o.EdgeId = &EdgeId
	}
    
	if EligibleAgentCounts, ok := ConversationmetricstopicconversationmetricrecordMap["eligibleAgentCounts"].([]interface{}); ok {
		EligibleAgentCountsString, _ := json.Marshal(EligibleAgentCounts)
		json.Unmarshal(EligibleAgentCountsString, &o.EligibleAgentCounts)
	}
	
	if ErrorCode, ok := ConversationmetricstopicconversationmetricrecordMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ExtendedDeliveryStatus, ok := ConversationmetricstopicconversationmetricrecordMap["extendedDeliveryStatus"].(string); ok {
		o.ExtendedDeliveryStatus = &ExtendedDeliveryStatus
	}
    
	if ExternalContactId, ok := ConversationmetricstopicconversationmetricrecordMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if ExternalMediaCount, ok := ConversationmetricstopicconversationmetricrecordMap["externalMediaCount"].(float64); ok {
		ExternalMediaCountInt := int(ExternalMediaCount)
		o.ExternalMediaCount = &ExternalMediaCountInt
	}
	
	if ExternalOrganizationId, ok := ConversationmetricstopicconversationmetricrecordMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
    
	if ExternalTag, ok := ConversationmetricstopicconversationmetricrecordMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
    
	if FirstQueue, ok := ConversationmetricstopicconversationmetricrecordMap["firstQueue"].(bool); ok {
		o.FirstQueue = &FirstQueue
	}
    
	if FlaggedReason, ok := ConversationmetricstopicconversationmetricrecordMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
    
	if FlowInType, ok := ConversationmetricstopicconversationmetricrecordMap["flowInType"].(string); ok {
		o.FlowInType = &FlowInType
	}
    
	if FlowOutType, ok := ConversationmetricstopicconversationmetricrecordMap["flowOutType"].(string); ok {
		o.FlowOutType = &FlowOutType
	}
    
	if GroupId, ok := ConversationmetricstopicconversationmetricrecordMap["groupId"].(string); ok {
		o.GroupId = &GroupId
	}
    
	if InteractionType, ok := ConversationmetricstopicconversationmetricrecordMap["interactionType"].(string); ok {
		o.InteractionType = &InteractionType
	}
    
	if JourneyActionId, ok := ConversationmetricstopicconversationmetricrecordMap["journeyActionId"].(string); ok {
		o.JourneyActionId = &JourneyActionId
	}
    
	if JourneyActionMapId, ok := ConversationmetricstopicconversationmetricrecordMap["journeyActionMapId"].(string); ok {
		o.JourneyActionMapId = &JourneyActionMapId
	}
    
	if JourneyActionMapVersion, ok := ConversationmetricstopicconversationmetricrecordMap["journeyActionMapVersion"].(float64); ok {
		JourneyActionMapVersionInt := int(JourneyActionMapVersion)
		o.JourneyActionMapVersion = &JourneyActionMapVersionInt
	}
	
	if JourneyCustomerId, ok := ConversationmetricstopicconversationmetricrecordMap["journeyCustomerId"].(string); ok {
		o.JourneyCustomerId = &JourneyCustomerId
	}
    
	if JourneyCustomerIdType, ok := ConversationmetricstopicconversationmetricrecordMap["journeyCustomerIdType"].(string); ok {
		o.JourneyCustomerIdType = &JourneyCustomerIdType
	}
    
	if JourneyCustomerSessionId, ok := ConversationmetricstopicconversationmetricrecordMap["journeyCustomerSessionId"].(string); ok {
		o.JourneyCustomerSessionId = &JourneyCustomerSessionId
	}
    
	if JourneyCustomerSessionIdType, ok := ConversationmetricstopicconversationmetricrecordMap["journeyCustomerSessionIdType"].(string); ok {
		o.JourneyCustomerSessionIdType = &JourneyCustomerSessionIdType
	}
    
	if KnowledgeBaseIds, ok := ConversationmetricstopicconversationmetricrecordMap["knowledgeBaseIds"].([]interface{}); ok {
		KnowledgeBaseIdsString, _ := json.Marshal(KnowledgeBaseIds)
		json.Unmarshal(KnowledgeBaseIdsString, &o.KnowledgeBaseIds)
	}
	
	if MediaCount, ok := ConversationmetricstopicconversationmetricrecordMap["mediaCount"].(float64); ok {
		MediaCountInt := int(MediaCount)
		o.MediaCount = &MediaCountInt
	}
	
	if MediaType, ok := ConversationmetricstopicconversationmetricrecordMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if MessageType, ok := ConversationmetricstopicconversationmetricrecordMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if OriginatingDirection, ok := ConversationmetricstopicconversationmetricrecordMap["originatingDirection"].(string); ok {
		o.OriginatingDirection = &OriginatingDirection
	}
    
	if OutboundCampaignId, ok := ConversationmetricstopicconversationmetricrecordMap["outboundCampaignId"].(string); ok {
		o.OutboundCampaignId = &OutboundCampaignId
	}
    
	if OutboundContactId, ok := ConversationmetricstopicconversationmetricrecordMap["outboundContactId"].(string); ok {
		o.OutboundContactId = &OutboundContactId
	}
    
	if OutboundContactListId, ok := ConversationmetricstopicconversationmetricrecordMap["outboundContactListId"].(string); ok {
		o.OutboundContactListId = &OutboundContactListId
	}
    
	if ParticipantName, ok := ConversationmetricstopicconversationmetricrecordMap["participantName"].(string); ok {
		o.ParticipantName = &ParticipantName
	}
    
	if PeerId, ok := ConversationmetricstopicconversationmetricrecordMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if Provider, ok := ConversationmetricstopicconversationmetricrecordMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Purpose, ok := ConversationmetricstopicconversationmetricrecordMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if QueueId, ok := ConversationmetricstopicconversationmetricrecordMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if Remote, ok := ConversationmetricstopicconversationmetricrecordMap["remote"].(string); ok {
		o.Remote = &Remote
	}
    
	if RemovedSkillIds, ok := ConversationmetricstopicconversationmetricrecordMap["removedSkillIds"].([]interface{}); ok {
		RemovedSkillIdsString, _ := json.Marshal(RemovedSkillIds)
		json.Unmarshal(RemovedSkillIdsString, &o.RemovedSkillIds)
	}
	
	if RequestedLanguageId, ok := ConversationmetricstopicconversationmetricrecordMap["requestedLanguageId"].(string); ok {
		o.RequestedLanguageId = &RequestedLanguageId
	}
    
	if RequestedRoutingSkillIds, ok := ConversationmetricstopicconversationmetricrecordMap["requestedRoutingSkillIds"].([]interface{}); ok {
		RequestedRoutingSkillIdsString, _ := json.Marshal(RequestedRoutingSkillIds)
		json.Unmarshal(RequestedRoutingSkillIdsString, &o.RequestedRoutingSkillIds)
	}
	
	if RequestedRoutings, ok := ConversationmetricstopicconversationmetricrecordMap["requestedRoutings"].([]interface{}); ok {
		RequestedRoutingsString, _ := json.Marshal(RequestedRoutings)
		json.Unmarshal(RequestedRoutingsString, &o.RequestedRoutings)
	}
	
	if RoomId, ok := ConversationmetricstopicconversationmetricrecordMap["roomId"].(string); ok {
		o.RoomId = &RoomId
	}
    
	if RoutingPriority, ok := ConversationmetricstopicconversationmetricrecordMap["routingPriority"].(float64); ok {
		RoutingPriorityInt := int(RoutingPriority)
		o.RoutingPriority = &RoutingPriorityInt
	}
	
	if RoutingRing, ok := ConversationmetricstopicconversationmetricrecordMap["routingRing"].(float64); ok {
		RoutingRingInt := int(RoutingRing)
		o.RoutingRing = &RoutingRingInt
	}
	
	if RoutingRule, ok := ConversationmetricstopicconversationmetricrecordMap["routingRule"].(string); ok {
		o.RoutingRule = &RoutingRule
	}
    
	if RoutingRuleType, ok := ConversationmetricstopicconversationmetricrecordMap["routingRuleType"].(string); ok {
		o.RoutingRuleType = &RoutingRuleType
	}
    
	if SelectedAgentId, ok := ConversationmetricstopicconversationmetricrecordMap["selectedAgentId"].(string); ok {
		o.SelectedAgentId = &SelectedAgentId
	}
    
	if SelectedAgentRank, ok := ConversationmetricstopicconversationmetricrecordMap["selectedAgentRank"].(float64); ok {
		SelectedAgentRankInt := int(SelectedAgentRank)
		o.SelectedAgentRank = &SelectedAgentRankInt
	}
	
	if SelfServed, ok := ConversationmetricstopicconversationmetricrecordMap["selfServed"].(bool); ok {
		o.SelfServed = &SelfServed
	}
    
	if SessionDnis, ok := ConversationmetricstopicconversationmetricrecordMap["sessionDnis"].(string); ok {
		o.SessionDnis = &SessionDnis
	}
    
	if SessionId, ok := ConversationmetricstopicconversationmetricrecordMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if StationId, ok := ConversationmetricstopicconversationmetricrecordMap["stationId"].(string); ok {
		o.StationId = &StationId
	}
    
	if TeamId, ok := ConversationmetricstopicconversationmetricrecordMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
    
	if UsedRouting, ok := ConversationmetricstopicconversationmetricrecordMap["usedRouting"].(string); ok {
		o.UsedRouting = &UsedRouting
	}
    
	if UserId, ok := ConversationmetricstopicconversationmetricrecordMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if WaitingInteractionCounts, ok := ConversationmetricstopicconversationmetricrecordMap["waitingInteractionCounts"].([]interface{}); ok {
		WaitingInteractionCountsString, _ := json.Marshal(WaitingInteractionCounts)
		json.Unmarshal(WaitingInteractionCountsString, &o.WaitingInteractionCounts)
	}
	
	if WrapUpCode, ok := ConversationmetricstopicconversationmetricrecordMap["wrapUpCode"].(string); ok {
		o.WrapUpCode = &WrapUpCode
	}
    
	if ProposedAgents, ok := ConversationmetricstopicconversationmetricrecordMap["proposedAgents"].([]interface{}); ok {
		ProposedAgentsString, _ := json.Marshal(ProposedAgents)
		json.Unmarshal(ProposedAgentsString, &o.ProposedAgents)
	}
	
	if ScoredAgents, ok := ConversationmetricstopicconversationmetricrecordMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationmetricstopicconversationmetricrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
