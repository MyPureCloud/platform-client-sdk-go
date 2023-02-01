package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicsocialmediaparticipant
type Conversationsocialexpressioneventtopicsocialmediaparticipant struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Address
	Address *string `json:"address,omitempty"`

	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`

	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// EndTime
	EndTime *time.Time `json:"endTime,omitempty"`

	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`

	// Purpose
	Purpose *string `json:"purpose,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// InitialState
	InitialState *string `json:"initialState,omitempty"`

	// Direction
	Direction *string `json:"direction,omitempty"`

	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`

	// Held
	Held *bool `json:"held,omitempty"`

	// WrapupRequired
	WrapupRequired *bool `json:"wrapupRequired,omitempty"`

	// WrapupPrompt
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`

	// User
	User *Conversationsocialexpressioneventtopicurireference `json:"user,omitempty"`

	// Queue
	Queue *Conversationsocialexpressioneventtopicurireference `json:"queue,omitempty"`

	// Team
	Team *Conversationsocialexpressioneventtopicurireference `json:"team,omitempty"`

	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`

	// ErrorInfo
	ErrorInfo *Conversationsocialexpressioneventtopicerrorbody `json:"errorInfo,omitempty"`

	// Script
	Script *Conversationsocialexpressioneventtopicurireference `json:"script,omitempty"`

	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`

	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`

	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`

	// Provider
	Provider *string `json:"provider,omitempty"`

	// ExternalContact
	ExternalContact *Conversationsocialexpressioneventtopicurireference `json:"externalContact,omitempty"`

	// ExternalOrganization
	ExternalOrganization *Conversationsocialexpressioneventtopicurireference `json:"externalOrganization,omitempty"`

	// Wrapup
	Wrapup *Conversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`

	// ConversationRoutingData
	ConversationRoutingData *Conversationsocialexpressioneventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`

	// Peer
	Peer *string `json:"peer,omitempty"`

	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`

	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`

	// JourneyContext
	JourneyContext *Conversationsocialexpressioneventtopicjourneycontext `json:"journeyContext,omitempty"`

	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`

	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`

	// MediaRoles
	MediaRoles *[]string `json:"mediaRoles,omitempty"`

	// SocialMediaId
	SocialMediaId *string `json:"socialMediaId,omitempty"`

	// SocialMediaHub
	SocialMediaHub *string `json:"socialMediaHub,omitempty"`

	// SocialUserName
	SocialUserName *string `json:"socialUserName,omitempty"`

	// PreviewText
	PreviewText *string `json:"previewText,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationsocialexpressioneventtopicsocialmediaparticipant) SetField(field string, fieldValue interface{}) {
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

func (o Conversationsocialexpressioneventtopicsocialmediaparticipant) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartTime","ConnectedTime","EndTime","StartHoldTime","StartAcwTime","EndAcwTime", }
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
	type Alias Conversationsocialexpressioneventtopicsocialmediaparticipant
	
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
		
		InitialState *string `json:"initialState,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		WrapupRequired *bool `json:"wrapupRequired,omitempty"`
		
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		User *Conversationsocialexpressioneventtopicurireference `json:"user,omitempty"`
		
		Queue *Conversationsocialexpressioneventtopicurireference `json:"queue,omitempty"`
		
		Team *Conversationsocialexpressioneventtopicurireference `json:"team,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ErrorInfo *Conversationsocialexpressioneventtopicerrorbody `json:"errorInfo,omitempty"`
		
		Script *Conversationsocialexpressioneventtopicurireference `json:"script,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ExternalContact *Conversationsocialexpressioneventtopicurireference `json:"externalContact,omitempty"`
		
		ExternalOrganization *Conversationsocialexpressioneventtopicurireference `json:"externalOrganization,omitempty"`
		
		Wrapup *Conversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`
		
		ConversationRoutingData *Conversationsocialexpressioneventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		Peer *string `json:"peer,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		JourneyContext *Conversationsocialexpressioneventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		MediaRoles *[]string `json:"mediaRoles,omitempty"`
		
		SocialMediaId *string `json:"socialMediaId,omitempty"`
		
		SocialMediaHub *string `json:"socialMediaHub,omitempty"`
		
		SocialUserName *string `json:"socialUserName,omitempty"`
		
		PreviewText *string `json:"previewText,omitempty"`
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
		
		InitialState: o.InitialState,
		
		Direction: o.Direction,
		
		DisconnectType: o.DisconnectType,
		
		Held: o.Held,
		
		WrapupRequired: o.WrapupRequired,
		
		WrapupPrompt: o.WrapupPrompt,
		
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
		
		ConversationRoutingData: o.ConversationRoutingData,
		
		Peer: o.Peer,
		
		ScreenRecordingState: o.ScreenRecordingState,
		
		FlaggedReason: o.FlaggedReason,
		
		JourneyContext: o.JourneyContext,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		
		MediaRoles: o.MediaRoles,
		
		SocialMediaId: o.SocialMediaId,
		
		SocialMediaHub: o.SocialMediaHub,
		
		SocialUserName: o.SocialUserName,
		
		PreviewText: o.PreviewText,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationsocialexpressioneventtopicsocialmediaparticipant) UnmarshalJSON(b []byte) error {
	var ConversationsocialexpressioneventtopicsocialmediaparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsocialexpressioneventtopicsocialmediaparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Address, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if startTimeString, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if connectedTimeString, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if endTimeString, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if startHoldTimeString, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if Purpose, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if State, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Direction, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if DisconnectType, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if Held, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if WrapupRequired, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
    
	if WrapupPrompt, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
    
	if User, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Queue, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Team, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if Attributes, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if ErrorInfo, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Script, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if WrapupTimeoutMs, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if WrapupSkipped, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
    
	if AlertingTimeoutMs, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if Provider, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ExternalContact, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if ExternalOrganization, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if Wrapup, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if ConversationRoutingData, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if Peer, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["peer"].(string); ok {
		o.Peer = &Peer
	}
    
	if ScreenRecordingState, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["screenRecordingState"].(string); ok {
		o.ScreenRecordingState = &ScreenRecordingState
	}
    
	if FlaggedReason, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
    
	if JourneyContext, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if startAcwTimeString, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	
	if MediaRoles, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["mediaRoles"].([]interface{}); ok {
		MediaRolesString, _ := json.Marshal(MediaRoles)
		json.Unmarshal(MediaRolesString, &o.MediaRoles)
	}
	
	if SocialMediaId, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["socialMediaId"].(string); ok {
		o.SocialMediaId = &SocialMediaId
	}
    
	if SocialMediaHub, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["socialMediaHub"].(string); ok {
		o.SocialMediaHub = &SocialMediaHub
	}
    
	if SocialUserName, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["socialUserName"].(string); ok {
		o.SocialUserName = &SocialUserName
	}
    
	if PreviewText, ok := ConversationsocialexpressioneventtopicsocialmediaparticipantMap["previewText"].(string); ok {
		o.PreviewText = &PreviewText
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicsocialmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
