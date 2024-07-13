package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Callbackmediaparticipant
type Callbackmediaparticipant struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The unique participant ID.
	Id *string `json:"id,omitempty"`

	// Name - The display friendly name of the participant.
	Name *string `json:"name,omitempty"`

	// Address - The participant address.
	Address *string `json:"address,omitempty"`

	// StartTime - The time when this participant first joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`

	// ConnectedTime - The time when this participant went connected for this media (eg: video connected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// EndTime - The time when this participant went disconnected for this media (eg: video disconnected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`

	// StartHoldTime - The time when this participant's hold started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`

	// Purpose - The participant's purpose.  Values can be: 'agent', 'user', 'customer', 'external', 'acd', 'ivr
	Purpose *string `json:"purpose,omitempty"`

	// State - The participant's state.  Values can be: 'alerting', 'connected', 'disconnected', 'dialing', 'contacting
	State *string `json:"state,omitempty"`

	// Direction - The participant's direction.  Values can be: 'inbound' or 'outbound'
	Direction *string `json:"direction,omitempty"`

	// DisconnectType - The reason the participant was disconnected from the conversation.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// Held - Value is true when the participant is on hold.
	Held *bool `json:"held,omitempty"`

	// WrapupRequired - Value is true when the participant requires wrap-up.
	WrapupRequired *bool `json:"wrapupRequired,omitempty"`

	// WrapupPrompt - The wrap-up prompt indicating the type of wrap-up to be performed.
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`

	// MediaRoles - List of roles this participant's media has had on the conversation, ie monitor, coach, etc
	MediaRoles *[]string `json:"mediaRoles,omitempty"`

	// User - The PureCloud user for this participant.
	User *Domainentityref `json:"user,omitempty"`

	// Queue - The PureCloud queue for this participant.
	Queue *Domainentityref `json:"queue,omitempty"`

	// Team - The PureCloud team for this participant.
	Team *Domainentityref `json:"team,omitempty"`

	// Attributes - A list of ad-hoc attributes for the participant.
	Attributes *map[string]string `json:"attributes,omitempty"`

	// ErrorInfo - If the conversation ends in error, contains additional error details.
	ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`

	// Script - The Engage script that should be used by this participant.
	Script *Domainentityref `json:"script,omitempty"`

	// WrapupTimeoutMs - The amount of time the participant has to complete wrap-up.
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`

	// WrapupSkipped - Value is true when the participant has skipped wrap-up.
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`

	// AlertingTimeoutMs - Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`

	// Provider - The source provider for the communication.
	Provider *string `json:"provider,omitempty"`

	// ExternalContact - If this participant represents an external contact, then this will be the reference for the external contact.
	ExternalContact *Domainentityref `json:"externalContact,omitempty"`

	// ExternalOrganization - If this participant represents an external org, then this will be the reference for the external org.
	ExternalOrganization *Domainentityref `json:"externalOrganization,omitempty"`

	// Wrapup - Wrapup for this participant, if it has been applied.
	Wrapup *Wrapup `json:"wrapup,omitempty"`

	// Peer - The peer communication corresponding to a matching leg for this communication.
	Peer *string `json:"peer,omitempty"`

	// FlaggedReason - The reason specifying why participant flagged the conversation.
	FlaggedReason *string `json:"flaggedReason,omitempty"`

	// JourneyContext - Journey System data/context that is applicable to this communication.  When used for historical purposes, the context should be immutable.  When null, there is no applicable Journey System context.
	JourneyContext *Journeycontext `json:"journeyContext,omitempty"`

	// ConversationRoutingData - Information on how a communication should be routed to an agent.
	ConversationRoutingData *Conversationroutingdata `json:"conversationRoutingData,omitempty"`

	// StartAcwTime - The timestamp when this participant started after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`

	// EndAcwTime - The timestamp when this participant ended after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`

	// ParkTime - The time when this participant's communication was last parked.  Does not reset on resume. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ParkTime *time.Time `json:"parkTime,omitempty"`

	// OutboundPreview - The outbound preview associated with this callback.
	OutboundPreview *Dialerpreview `json:"outboundPreview,omitempty"`

	// Voicemail - The voicemail associated with this callback.
	Voicemail *Voicemail `json:"voicemail,omitempty"`

	// CallbackNumbers - The list of phone number to use for this callback.
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`

	// CallbackUserName - The name of the callback target.
	CallbackUserName *string `json:"callbackUserName,omitempty"`

	// ExternalCampaign - True if the call for the callback uses external dialing.
	ExternalCampaign *bool `json:"externalCampaign,omitempty"`

	// SkipEnabled - If true, the callback can be skipped.
	SkipEnabled *bool `json:"skipEnabled,omitempty"`

	// TimeoutSeconds - Duration in seconds before the callback will be auto-dialed.
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`

	// AutomatedCallbackConfigId - The id of the config for automatically placing the callback (and handling the disposition). If absent, the callback will not be placed automatically but routed to an agent as per normal.
	AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`

	// CallbackScheduledTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Callbackmediaparticipant) SetField(field string, fieldValue interface{}) {
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

func (o Callbackmediaparticipant) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartTime","ConnectedTime","EndTime","StartHoldTime","StartAcwTime","EndAcwTime","ParkTime","CallbackScheduledTime", }
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
	type Alias Callbackmediaparticipant
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	StartHoldTime := new(string)
	if o.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(o.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	StartAcwTime := new(string)
	if o.StartAcwTime != nil {
		
		*StartAcwTime = timeutil.Strftime(o.StartAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartAcwTime = nil
	}
	
	EndAcwTime := new(string)
	if o.EndAcwTime != nil {
		
		*EndAcwTime = timeutil.Strftime(o.EndAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndAcwTime = nil
	}
	
	ParkTime := new(string)
	if o.ParkTime != nil {
		
		*ParkTime = timeutil.Strftime(o.ParkTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ParkTime = nil
	}
	
	CallbackScheduledTime := new(string)
	if o.CallbackScheduledTime != nil {
		
		*CallbackScheduledTime = timeutil.Strftime(o.CallbackScheduledTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallbackScheduledTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		Purpose *string `json:"purpose,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		WrapupRequired *bool `json:"wrapupRequired,omitempty"`
		
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		MediaRoles *[]string `json:"mediaRoles,omitempty"`
		
		User *Domainentityref `json:"user,omitempty"`
		
		Queue *Domainentityref `json:"queue,omitempty"`
		
		Team *Domainentityref `json:"team,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`
		
		Script *Domainentityref `json:"script,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ExternalContact *Domainentityref `json:"externalContact,omitempty"`
		
		ExternalOrganization *Domainentityref `json:"externalOrganization,omitempty"`
		
		Wrapup *Wrapup `json:"wrapup,omitempty"`
		
		Peer *string `json:"peer,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		JourneyContext *Journeycontext `json:"journeyContext,omitempty"`
		
		ConversationRoutingData *Conversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		ParkTime *string `json:"parkTime,omitempty"`
		
		OutboundPreview *Dialerpreview `json:"outboundPreview,omitempty"`
		
		Voicemail *Voicemail `json:"voicemail,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		ExternalCampaign *bool `json:"externalCampaign,omitempty"`
		
		SkipEnabled *bool `json:"skipEnabled,omitempty"`
		
		TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
		
		AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`
		
		CallbackScheduledTime *string `json:"callbackScheduledTime,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Address: o.Address,
		
		StartTime: StartTime,
		
		ConnectedTime: ConnectedTime,
		
		EndTime: EndTime,
		
		StartHoldTime: StartHoldTime,
		
		Purpose: o.Purpose,
		
		State: o.State,
		
		Direction: o.Direction,
		
		DisconnectType: o.DisconnectType,
		
		Held: o.Held,
		
		WrapupRequired: o.WrapupRequired,
		
		WrapupPrompt: o.WrapupPrompt,
		
		MediaRoles: o.MediaRoles,
		
		User: o.User,
		
		Queue: o.Queue,
		
		Team: o.Team,
		
		Attributes: o.Attributes,
		
		ErrorInfo: o.ErrorInfo,
		
		Script: o.Script,
		
		WrapupTimeoutMs: o.WrapupTimeoutMs,
		
		WrapupSkipped: o.WrapupSkipped,
		
		AlertingTimeoutMs: o.AlertingTimeoutMs,
		
		Provider: o.Provider,
		
		ExternalContact: o.ExternalContact,
		
		ExternalOrganization: o.ExternalOrganization,
		
		Wrapup: o.Wrapup,
		
		Peer: o.Peer,
		
		FlaggedReason: o.FlaggedReason,
		
		JourneyContext: o.JourneyContext,
		
		ConversationRoutingData: o.ConversationRoutingData,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		
		ParkTime: ParkTime,
		
		OutboundPreview: o.OutboundPreview,
		
		Voicemail: o.Voicemail,
		
		CallbackNumbers: o.CallbackNumbers,
		
		CallbackUserName: o.CallbackUserName,
		
		ExternalCampaign: o.ExternalCampaign,
		
		SkipEnabled: o.SkipEnabled,
		
		TimeoutSeconds: o.TimeoutSeconds,
		
		AutomatedCallbackConfigId: o.AutomatedCallbackConfigId,
		
		CallbackScheduledTime: CallbackScheduledTime,
		Alias:    (Alias)(o),
	})
}

func (o *Callbackmediaparticipant) UnmarshalJSON(b []byte) error {
	var CallbackmediaparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &CallbackmediaparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CallbackmediaparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CallbackmediaparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Address, ok := CallbackmediaparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if startTimeString, ok := CallbackmediaparticipantMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if connectedTimeString, ok := CallbackmediaparticipantMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if endTimeString, ok := CallbackmediaparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if startHoldTimeString, ok := CallbackmediaparticipantMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if Purpose, ok := CallbackmediaparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if State, ok := CallbackmediaparticipantMap["state"].(string); ok {
		o.State = &State
	}
    
	if Direction, ok := CallbackmediaparticipantMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if DisconnectType, ok := CallbackmediaparticipantMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if Held, ok := CallbackmediaparticipantMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if WrapupRequired, ok := CallbackmediaparticipantMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
    
	if WrapupPrompt, ok := CallbackmediaparticipantMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
    
	if MediaRoles, ok := CallbackmediaparticipantMap["mediaRoles"].([]interface{}); ok {
		MediaRolesString, _ := json.Marshal(MediaRoles)
		json.Unmarshal(MediaRolesString, &o.MediaRoles)
	}
	
	if User, ok := CallbackmediaparticipantMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Queue, ok := CallbackmediaparticipantMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Team, ok := CallbackmediaparticipantMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if Attributes, ok := CallbackmediaparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if ErrorInfo, ok := CallbackmediaparticipantMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Script, ok := CallbackmediaparticipantMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if WrapupTimeoutMs, ok := CallbackmediaparticipantMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if WrapupSkipped, ok := CallbackmediaparticipantMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
    
	if AlertingTimeoutMs, ok := CallbackmediaparticipantMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if Provider, ok := CallbackmediaparticipantMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ExternalContact, ok := CallbackmediaparticipantMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if ExternalOrganization, ok := CallbackmediaparticipantMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if Wrapup, ok := CallbackmediaparticipantMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if Peer, ok := CallbackmediaparticipantMap["peer"].(string); ok {
		o.Peer = &Peer
	}
    
	if FlaggedReason, ok := CallbackmediaparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
    
	if JourneyContext, ok := CallbackmediaparticipantMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if ConversationRoutingData, ok := CallbackmediaparticipantMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if startAcwTimeString, ok := CallbackmediaparticipantMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := CallbackmediaparticipantMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	
	if parkTimeString, ok := CallbackmediaparticipantMap["parkTime"].(string); ok {
		ParkTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", parkTimeString)
		o.ParkTime = &ParkTime
	}
	
	if OutboundPreview, ok := CallbackmediaparticipantMap["outboundPreview"].(map[string]interface{}); ok {
		OutboundPreviewString, _ := json.Marshal(OutboundPreview)
		json.Unmarshal(OutboundPreviewString, &o.OutboundPreview)
	}
	
	if Voicemail, ok := CallbackmediaparticipantMap["voicemail"].(map[string]interface{}); ok {
		VoicemailString, _ := json.Marshal(Voicemail)
		json.Unmarshal(VoicemailString, &o.Voicemail)
	}
	
	if CallbackNumbers, ok := CallbackmediaparticipantMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if CallbackUserName, ok := CallbackmediaparticipantMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
    
	if ExternalCampaign, ok := CallbackmediaparticipantMap["externalCampaign"].(bool); ok {
		o.ExternalCampaign = &ExternalCampaign
	}
    
	if SkipEnabled, ok := CallbackmediaparticipantMap["skipEnabled"].(bool); ok {
		o.SkipEnabled = &SkipEnabled
	}
    
	if TimeoutSeconds, ok := CallbackmediaparticipantMap["timeoutSeconds"].(float64); ok {
		TimeoutSecondsInt := int(TimeoutSeconds)
		o.TimeoutSeconds = &TimeoutSecondsInt
	}
	
	if AutomatedCallbackConfigId, ok := CallbackmediaparticipantMap["automatedCallbackConfigId"].(string); ok {
		o.AutomatedCallbackConfigId = &AutomatedCallbackConfigId
	}
    
	if callbackScheduledTimeString, ok := CallbackmediaparticipantMap["callbackScheduledTime"].(string); ok {
		CallbackScheduledTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", callbackScheduledTimeString)
		o.CallbackScheduledTime = &CallbackScheduledTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callbackmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
