package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopiccallbackmediaparticipant
type Queueconversationcallbackeventtopiccallbackmediaparticipant struct { 
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
	User *Queueconversationcallbackeventtopicurireference `json:"user,omitempty"`

	// Queue
	Queue *Queueconversationcallbackeventtopicurireference `json:"queue,omitempty"`

	// Team
	Team *Queueconversationcallbackeventtopicurireference `json:"team,omitempty"`

	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`

	// ErrorInfo
	ErrorInfo *Queueconversationcallbackeventtopicerrorbody `json:"errorInfo,omitempty"`

	// Script
	Script *Queueconversationcallbackeventtopicurireference `json:"script,omitempty"`

	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`

	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`

	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`

	// Provider
	Provider *string `json:"provider,omitempty"`

	// ExternalContact
	ExternalContact *Queueconversationcallbackeventtopicurireference `json:"externalContact,omitempty"`

	// ExternalContactInitialDivisionId
	ExternalContactInitialDivisionId *string `json:"externalContactInitialDivisionId,omitempty"`

	// ExternalOrganization
	ExternalOrganization *Queueconversationcallbackeventtopicurireference `json:"externalOrganization,omitempty"`

	// Wrapup
	Wrapup *Queueconversationcallbackeventtopicwrapup `json:"wrapup,omitempty"`

	// ConversationRoutingData
	ConversationRoutingData *Queueconversationcallbackeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`

	// Peer
	Peer *string `json:"peer,omitempty"`

	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`

	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`

	// JourneyContext
	JourneyContext *Queueconversationcallbackeventtopicjourneycontext `json:"journeyContext,omitempty"`

	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`

	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`

	// ResumeTime
	ResumeTime *time.Time `json:"resumeTime,omitempty"`

	// ParkTime
	ParkTime *time.Time `json:"parkTime,omitempty"`

	// MediaRoles
	MediaRoles *[]string `json:"mediaRoles,omitempty"`

	// QueueMediaSettings
	QueueMediaSettings *Queueconversationcallbackeventtopicqueuemediasettings `json:"queueMediaSettings,omitempty"`

	// OutboundPreview
	OutboundPreview *Queueconversationcallbackeventtopicdialerpreview `json:"outboundPreview,omitempty"`

	// Voicemail
	Voicemail *Queueconversationcallbackeventtopicvoicemail `json:"voicemail,omitempty"`

	// CallbackNumbers
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`

	// CallbackUserName
	CallbackUserName *string `json:"callbackUserName,omitempty"`

	// SkipEnabled
	SkipEnabled *bool `json:"skipEnabled,omitempty"`

	// ExternalCampaign
	ExternalCampaign *bool `json:"externalCampaign,omitempty"`

	// TimeoutSeconds
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`

	// CallbackScheduledTime
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`

	// AutomatedCallbackConfigId
	AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueconversationcallbackeventtopiccallbackmediaparticipant) SetField(field string, fieldValue interface{}) {
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

func (o Queueconversationcallbackeventtopiccallbackmediaparticipant) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartTime","ConnectedTime","EndTime","StartHoldTime","StartAcwTime","EndAcwTime","ResumeTime","ParkTime","CallbackScheduledTime", }
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
	type Alias Queueconversationcallbackeventtopiccallbackmediaparticipant
	
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
	
	ResumeTime := new(string)
	if o.ResumeTime != nil {
		
		*ResumeTime = timeutil.Strftime(o.ResumeTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ResumeTime = nil
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
		
		InitialState *string `json:"initialState,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		WrapupRequired *bool `json:"wrapupRequired,omitempty"`
		
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		User *Queueconversationcallbackeventtopicurireference `json:"user,omitempty"`
		
		Queue *Queueconversationcallbackeventtopicurireference `json:"queue,omitempty"`
		
		Team *Queueconversationcallbackeventtopicurireference `json:"team,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ErrorInfo *Queueconversationcallbackeventtopicerrorbody `json:"errorInfo,omitempty"`
		
		Script *Queueconversationcallbackeventtopicurireference `json:"script,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ExternalContact *Queueconversationcallbackeventtopicurireference `json:"externalContact,omitempty"`
		
		ExternalContactInitialDivisionId *string `json:"externalContactInitialDivisionId,omitempty"`
		
		ExternalOrganization *Queueconversationcallbackeventtopicurireference `json:"externalOrganization,omitempty"`
		
		Wrapup *Queueconversationcallbackeventtopicwrapup `json:"wrapup,omitempty"`
		
		ConversationRoutingData *Queueconversationcallbackeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		Peer *string `json:"peer,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		JourneyContext *Queueconversationcallbackeventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		ResumeTime *string `json:"resumeTime,omitempty"`
		
		ParkTime *string `json:"parkTime,omitempty"`
		
		MediaRoles *[]string `json:"mediaRoles,omitempty"`
		
		QueueMediaSettings *Queueconversationcallbackeventtopicqueuemediasettings `json:"queueMediaSettings,omitempty"`
		
		OutboundPreview *Queueconversationcallbackeventtopicdialerpreview `json:"outboundPreview,omitempty"`
		
		Voicemail *Queueconversationcallbackeventtopicvoicemail `json:"voicemail,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		SkipEnabled *bool `json:"skipEnabled,omitempty"`
		
		ExternalCampaign *bool `json:"externalCampaign,omitempty"`
		
		TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
		
		CallbackScheduledTime *string `json:"callbackScheduledTime,omitempty"`
		
		AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`
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
		
		ExternalContactInitialDivisionId: o.ExternalContactInitialDivisionId,
		
		ExternalOrganization: o.ExternalOrganization,
		
		Wrapup: o.Wrapup,
		
		ConversationRoutingData: o.ConversationRoutingData,
		
		Peer: o.Peer,
		
		ScreenRecordingState: o.ScreenRecordingState,
		
		FlaggedReason: o.FlaggedReason,
		
		JourneyContext: o.JourneyContext,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		
		ResumeTime: ResumeTime,
		
		ParkTime: ParkTime,
		
		MediaRoles: o.MediaRoles,
		
		QueueMediaSettings: o.QueueMediaSettings,
		
		OutboundPreview: o.OutboundPreview,
		
		Voicemail: o.Voicemail,
		
		CallbackNumbers: o.CallbackNumbers,
		
		CallbackUserName: o.CallbackUserName,
		
		SkipEnabled: o.SkipEnabled,
		
		ExternalCampaign: o.ExternalCampaign,
		
		TimeoutSeconds: o.TimeoutSeconds,
		
		CallbackScheduledTime: CallbackScheduledTime,
		
		AutomatedCallbackConfigId: o.AutomatedCallbackConfigId,
		Alias:    (Alias)(o),
	})
}

func (o *Queueconversationcallbackeventtopiccallbackmediaparticipant) UnmarshalJSON(b []byte) error {
	var QueueconversationcallbackeventtopiccallbackmediaparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcallbackeventtopiccallbackmediaparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Address, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if startTimeString, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if connectedTimeString, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if endTimeString, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if startHoldTimeString, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if Purpose, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if State, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Direction, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if DisconnectType, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if Held, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if WrapupRequired, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
    
	if WrapupPrompt, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
    
	if User, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Queue, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Team, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if Attributes, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if ErrorInfo, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Script, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if WrapupTimeoutMs, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if WrapupSkipped, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
    
	if AlertingTimeoutMs, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if Provider, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ExternalContact, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if ExternalContactInitialDivisionId, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["externalContactInitialDivisionId"].(string); ok {
		o.ExternalContactInitialDivisionId = &ExternalContactInitialDivisionId
	}
    
	if ExternalOrganization, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if Wrapup, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if ConversationRoutingData, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if Peer, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["peer"].(string); ok {
		o.Peer = &Peer
	}
    
	if ScreenRecordingState, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["screenRecordingState"].(string); ok {
		o.ScreenRecordingState = &ScreenRecordingState
	}
    
	if FlaggedReason, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
    
	if JourneyContext, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if startAcwTimeString, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	
	if resumeTimeString, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["resumeTime"].(string); ok {
		ResumeTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", resumeTimeString)
		o.ResumeTime = &ResumeTime
	}
	
	if parkTimeString, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["parkTime"].(string); ok {
		ParkTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", parkTimeString)
		o.ParkTime = &ParkTime
	}
	
	if MediaRoles, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["mediaRoles"].([]interface{}); ok {
		MediaRolesString, _ := json.Marshal(MediaRoles)
		json.Unmarshal(MediaRolesString, &o.MediaRoles)
	}
	
	if QueueMediaSettings, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["queueMediaSettings"].(map[string]interface{}); ok {
		QueueMediaSettingsString, _ := json.Marshal(QueueMediaSettings)
		json.Unmarshal(QueueMediaSettingsString, &o.QueueMediaSettings)
	}
	
	if OutboundPreview, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["outboundPreview"].(map[string]interface{}); ok {
		OutboundPreviewString, _ := json.Marshal(OutboundPreview)
		json.Unmarshal(OutboundPreviewString, &o.OutboundPreview)
	}
	
	if Voicemail, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["voicemail"].(map[string]interface{}); ok {
		VoicemailString, _ := json.Marshal(Voicemail)
		json.Unmarshal(VoicemailString, &o.Voicemail)
	}
	
	if CallbackNumbers, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if CallbackUserName, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
    
	if SkipEnabled, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["skipEnabled"].(bool); ok {
		o.SkipEnabled = &SkipEnabled
	}
    
	if ExternalCampaign, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["externalCampaign"].(bool); ok {
		o.ExternalCampaign = &ExternalCampaign
	}
    
	if TimeoutSeconds, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["timeoutSeconds"].(float64); ok {
		TimeoutSecondsInt := int(TimeoutSeconds)
		o.TimeoutSeconds = &TimeoutSecondsInt
	}
	
	if callbackScheduledTimeString, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["callbackScheduledTime"].(string); ok {
		CallbackScheduledTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", callbackScheduledTimeString)
		o.CallbackScheduledTime = &CallbackScheduledTime
	}
	
	if AutomatedCallbackConfigId, ok := QueueconversationcallbackeventtopiccallbackmediaparticipantMap["automatedCallbackConfigId"].(string); ok {
		o.AutomatedCallbackConfigId = &AutomatedCallbackConfigId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopiccallbackmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
