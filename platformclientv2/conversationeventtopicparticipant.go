package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicparticipant
type Conversationeventtopicparticipant struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A globally unique identifier for this conversation.
	Id *string `json:"id,omitempty"`

	// ConnectedTime - The timestamp when this participant was connected to the conversation in the provider clock.
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// EndTime - The timestamp when this participant disconnected from the conversation in the provider clock.
	EndTime *time.Time `json:"endTime,omitempty"`

	// UserId - If this participant represents a user, then this will be the globally unique identifier for the user.
	UserId *string `json:"userId,omitempty"`

	// ExternalContactId - If this participant represents an external contact, then this will be the globally unique identifier for the external contact.
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// ExternalOrganizationId - If this participant represents an external org, then this will be the globally unique identifier for the external org.
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`

	// Name - A human readable name identifying the participant.
	Name *string `json:"name,omitempty"`

	// QueueId - If present, the queue id that the communication channel came in on.
	QueueId *string `json:"queueId,omitempty"`

	// GroupId - If present, the group id that the participant represents.
	GroupId *string `json:"groupId,omitempty"`

	// TeamId - The team id that this participant is a member of when added to the conversation.
	TeamId *string `json:"teamId,omitempty"`

	// Purpose - A well known string that specifies the purpose or type of this participant.
	Purpose *string `json:"purpose,omitempty"`

	// ConsultParticipantId - If this participant is part of a consult transfer, then this will be the participant id of the participant being transferred.
	ConsultParticipantId *string `json:"consultParticipantId,omitempty"`

	// Address - The address for the this participant. For a phone call this will be the ANI.
	Address *string `json:"address,omitempty"`

	// WrapupRequired - True iff this participant is required to enter wrapup for this conversation.
	WrapupRequired *bool `json:"wrapupRequired,omitempty"`

	// WrapupExpected - True when a participant is expected to enter a wrapup code once the call connects.
	WrapupExpected *bool `json:"wrapupExpected,omitempty"`

	// WrapupPrompt - This field controls how the UI prompts the agent for a wrapup.
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`

	// WrapupTimeoutMs - Specifies how long a timed ACW session will last.
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`

	// Wrapup
	Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`

	// StartAcwTime - The timestamp when this participant started after-call work.
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`

	// EndAcwTime - The timestamp when this participant ended after-call work.
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`

	// ConversationRoutingData
	ConversationRoutingData *Conversationeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`

	// AlertingTimeoutMs - Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`

	// MonitoredParticipantId - If this participant is a monitor, then this will be the id of the participant that is being monitored.
	MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`

	// CoachedParticipantId - If this participant is a coach, then this will be the id of the participant that is being coached.
	CoachedParticipantId *string `json:"coachedParticipantId,omitempty"`

	// BargedParticipantId - If this participant created a barge in conference, then this will be the id of the participant that is barged in.
	BargedParticipantId *string `json:"bargedParticipantId,omitempty"`

	// MediaRoles - List of roles this participant's media has had on the conversation, ie monitor, coach, etc.
	MediaRoles *[]string `json:"mediaRoles,omitempty"`

	// ScreenRecordingState - The current screen recording state for this participant.
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`

	// FlaggedReason - If this participant has flagged the conversation, the reason code given.
	FlaggedReason *string `json:"flaggedReason,omitempty"`

	// Attributes - Additional participant attributes
	Attributes *map[string]string `json:"attributes,omitempty"`

	// Calls
	Calls *[]Conversationeventtopiccall `json:"calls,omitempty"`

	// Callbacks
	Callbacks *[]Conversationeventtopiccallback `json:"callbacks,omitempty"`

	// Chats
	Chats *[]Conversationeventtopicchat `json:"chats,omitempty"`

	// Cobrowsesessions
	Cobrowsesessions *[]Conversationeventtopiccobrowse `json:"cobrowsesessions,omitempty"`

	// Emails
	Emails *[]Conversationeventtopicemail `json:"emails,omitempty"`

	// Messages
	Messages *[]Conversationeventtopicmessage `json:"messages,omitempty"`

	// InternalMessages
	InternalMessages *[]Conversationeventtopicinternalmessage `json:"internalMessages,omitempty"`

	// Screenshares
	Screenshares *[]Conversationeventtopicscreenshare `json:"screenshares,omitempty"`

	// SocialExpressions
	SocialExpressions *[]Conversationeventtopicsocialexpression `json:"socialExpressions,omitempty"`

	// Videos
	Videos *[]Conversationeventtopicvideo `json:"videos,omitempty"`

	// Workflow
	Workflow *Conversationeventtopicworkflow `json:"workflow,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationeventtopicparticipant) SetField(field string, fieldValue interface{}) {
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

func (o Conversationeventtopicparticipant) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ConnectedTime","EndTime","StartAcwTime","EndAcwTime", }
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
	type Alias Conversationeventtopicparticipant
	
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
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		GroupId *string `json:"groupId,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		Purpose *string `json:"purpose,omitempty"`
		
		ConsultParticipantId *string `json:"consultParticipantId,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		WrapupRequired *bool `json:"wrapupRequired,omitempty"`
		
		WrapupExpected *bool `json:"wrapupExpected,omitempty"`
		
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		ConversationRoutingData *Conversationeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`
		
		CoachedParticipantId *string `json:"coachedParticipantId,omitempty"`
		
		BargedParticipantId *string `json:"bargedParticipantId,omitempty"`
		
		MediaRoles *[]string `json:"mediaRoles,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		Calls *[]Conversationeventtopiccall `json:"calls,omitempty"`
		
		Callbacks *[]Conversationeventtopiccallback `json:"callbacks,omitempty"`
		
		Chats *[]Conversationeventtopicchat `json:"chats,omitempty"`
		
		Cobrowsesessions *[]Conversationeventtopiccobrowse `json:"cobrowsesessions,omitempty"`
		
		Emails *[]Conversationeventtopicemail `json:"emails,omitempty"`
		
		Messages *[]Conversationeventtopicmessage `json:"messages,omitempty"`
		
		InternalMessages *[]Conversationeventtopicinternalmessage `json:"internalMessages,omitempty"`
		
		Screenshares *[]Conversationeventtopicscreenshare `json:"screenshares,omitempty"`
		
		SocialExpressions *[]Conversationeventtopicsocialexpression `json:"socialExpressions,omitempty"`
		
		Videos *[]Conversationeventtopicvideo `json:"videos,omitempty"`
		
		Workflow *Conversationeventtopicworkflow `json:"workflow,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ConnectedTime: ConnectedTime,
		
		EndTime: EndTime,
		
		UserId: o.UserId,
		
		ExternalContactId: o.ExternalContactId,
		
		ExternalOrganizationId: o.ExternalOrganizationId,
		
		Name: o.Name,
		
		QueueId: o.QueueId,
		
		GroupId: o.GroupId,
		
		TeamId: o.TeamId,
		
		Purpose: o.Purpose,
		
		ConsultParticipantId: o.ConsultParticipantId,
		
		Address: o.Address,
		
		WrapupRequired: o.WrapupRequired,
		
		WrapupExpected: o.WrapupExpected,
		
		WrapupPrompt: o.WrapupPrompt,
		
		WrapupTimeoutMs: o.WrapupTimeoutMs,
		
		Wrapup: o.Wrapup,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		
		ConversationRoutingData: o.ConversationRoutingData,
		
		AlertingTimeoutMs: o.AlertingTimeoutMs,
		
		MonitoredParticipantId: o.MonitoredParticipantId,
		
		CoachedParticipantId: o.CoachedParticipantId,
		
		BargedParticipantId: o.BargedParticipantId,
		
		MediaRoles: o.MediaRoles,
		
		ScreenRecordingState: o.ScreenRecordingState,
		
		FlaggedReason: o.FlaggedReason,
		
		Attributes: o.Attributes,
		
		Calls: o.Calls,
		
		Callbacks: o.Callbacks,
		
		Chats: o.Chats,
		
		Cobrowsesessions: o.Cobrowsesessions,
		
		Emails: o.Emails,
		
		Messages: o.Messages,
		
		InternalMessages: o.InternalMessages,
		
		Screenshares: o.Screenshares,
		
		SocialExpressions: o.SocialExpressions,
		
		Videos: o.Videos,
		
		Workflow: o.Workflow,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationeventtopicparticipant) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationeventtopicparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if connectedTimeString, ok := ConversationeventtopicparticipantMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if endTimeString, ok := ConversationeventtopicparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if UserId, ok := ConversationeventtopicparticipantMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if ExternalContactId, ok := ConversationeventtopicparticipantMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if ExternalOrganizationId, ok := ConversationeventtopicparticipantMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
    
	if Name, ok := ConversationeventtopicparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if QueueId, ok := ConversationeventtopicparticipantMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if GroupId, ok := ConversationeventtopicparticipantMap["groupId"].(string); ok {
		o.GroupId = &GroupId
	}
    
	if TeamId, ok := ConversationeventtopicparticipantMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
    
	if Purpose, ok := ConversationeventtopicparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if ConsultParticipantId, ok := ConversationeventtopicparticipantMap["consultParticipantId"].(string); ok {
		o.ConsultParticipantId = &ConsultParticipantId
	}
    
	if Address, ok := ConversationeventtopicparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if WrapupRequired, ok := ConversationeventtopicparticipantMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
    
	if WrapupExpected, ok := ConversationeventtopicparticipantMap["wrapupExpected"].(bool); ok {
		o.WrapupExpected = &WrapupExpected
	}
    
	if WrapupPrompt, ok := ConversationeventtopicparticipantMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
    
	if WrapupTimeoutMs, ok := ConversationeventtopicparticipantMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if Wrapup, ok := ConversationeventtopicparticipantMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if startAcwTimeString, ok := ConversationeventtopicparticipantMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := ConversationeventtopicparticipantMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	
	if ConversationRoutingData, ok := ConversationeventtopicparticipantMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if AlertingTimeoutMs, ok := ConversationeventtopicparticipantMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if MonitoredParticipantId, ok := ConversationeventtopicparticipantMap["monitoredParticipantId"].(string); ok {
		o.MonitoredParticipantId = &MonitoredParticipantId
	}
    
	if CoachedParticipantId, ok := ConversationeventtopicparticipantMap["coachedParticipantId"].(string); ok {
		o.CoachedParticipantId = &CoachedParticipantId
	}
    
	if BargedParticipantId, ok := ConversationeventtopicparticipantMap["bargedParticipantId"].(string); ok {
		o.BargedParticipantId = &BargedParticipantId
	}
    
	if MediaRoles, ok := ConversationeventtopicparticipantMap["mediaRoles"].([]interface{}); ok {
		MediaRolesString, _ := json.Marshal(MediaRoles)
		json.Unmarshal(MediaRolesString, &o.MediaRoles)
	}
	
	if ScreenRecordingState, ok := ConversationeventtopicparticipantMap["screenRecordingState"].(string); ok {
		o.ScreenRecordingState = &ScreenRecordingState
	}
    
	if FlaggedReason, ok := ConversationeventtopicparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
    
	if Attributes, ok := ConversationeventtopicparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Calls, ok := ConversationeventtopicparticipantMap["calls"].([]interface{}); ok {
		CallsString, _ := json.Marshal(Calls)
		json.Unmarshal(CallsString, &o.Calls)
	}
	
	if Callbacks, ok := ConversationeventtopicparticipantMap["callbacks"].([]interface{}); ok {
		CallbacksString, _ := json.Marshal(Callbacks)
		json.Unmarshal(CallbacksString, &o.Callbacks)
	}
	
	if Chats, ok := ConversationeventtopicparticipantMap["chats"].([]interface{}); ok {
		ChatsString, _ := json.Marshal(Chats)
		json.Unmarshal(ChatsString, &o.Chats)
	}
	
	if Cobrowsesessions, ok := ConversationeventtopicparticipantMap["cobrowsesessions"].([]interface{}); ok {
		CobrowsesessionsString, _ := json.Marshal(Cobrowsesessions)
		json.Unmarshal(CobrowsesessionsString, &o.Cobrowsesessions)
	}
	
	if Emails, ok := ConversationeventtopicparticipantMap["emails"].([]interface{}); ok {
		EmailsString, _ := json.Marshal(Emails)
		json.Unmarshal(EmailsString, &o.Emails)
	}
	
	if Messages, ok := ConversationeventtopicparticipantMap["messages"].([]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	
	if InternalMessages, ok := ConversationeventtopicparticipantMap["internalMessages"].([]interface{}); ok {
		InternalMessagesString, _ := json.Marshal(InternalMessages)
		json.Unmarshal(InternalMessagesString, &o.InternalMessages)
	}
	
	if Screenshares, ok := ConversationeventtopicparticipantMap["screenshares"].([]interface{}); ok {
		ScreensharesString, _ := json.Marshal(Screenshares)
		json.Unmarshal(ScreensharesString, &o.Screenshares)
	}
	
	if SocialExpressions, ok := ConversationeventtopicparticipantMap["socialExpressions"].([]interface{}); ok {
		SocialExpressionsString, _ := json.Marshal(SocialExpressions)
		json.Unmarshal(SocialExpressionsString, &o.SocialExpressions)
	}
	
	if Videos, ok := ConversationeventtopicparticipantMap["videos"].([]interface{}); ok {
		VideosString, _ := json.Marshal(Videos)
		json.Unmarshal(VideosString, &o.Videos)
	}
	
	if Workflow, ok := ConversationeventtopicparticipantMap["workflow"].(map[string]interface{}); ok {
		WorkflowString, _ := json.Marshal(Workflow)
		json.Unmarshal(WorkflowString, &o.Workflow)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
