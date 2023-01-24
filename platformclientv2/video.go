package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Video
type Video struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`

	// InitialState - The initial connection state of this communication.
	InitialState *string `json:"initialState,omitempty"`

	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`

	// Context - The room id context (xmpp jid) for the conference session.
	Context *string `json:"context,omitempty"`

	// AudioMuted - Indicates whether this participant has muted their outgoing audio.
	AudioMuted *bool `json:"audioMuted,omitempty"`

	// VideoMuted - Indicates whether this participant has muted/paused their outgoing video.
	VideoMuted *bool `json:"videoMuted,omitempty"`

	// SharingScreen - Indicates whether this participant is sharing their screen to the session.
	SharingScreen *bool `json:"sharingScreen,omitempty"`

	// PeerCount - The number of peer participants from the perspective of the participant in the conference.
	PeerCount *int `json:"peerCount,omitempty"`

	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAlertingTime *time.Time `json:"startAlertingTime,omitempty"`

	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`

	// Provider - The source provider for the video.
	Provider *string `json:"provider,omitempty"`

	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`

	// Msids - List of media stream ids
	Msids *[]string `json:"msids,omitempty"`

	// Self - Address and name data for a call endpoint.
	Self *Address `json:"self,omitempty"`

	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`

	// AfterCallWork - After-call work for the communication.
	AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`

	// AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Video) SetField(field string, fieldValue interface{}) {
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

func (o Video) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartAlertingTime","ConnectedTime","DisconnectedTime", }
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
	type Alias Video
	
	StartAlertingTime := new(string)
	if o.StartAlertingTime != nil {
		
		*StartAlertingTime = timeutil.Strftime(o.StartAlertingTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartAlertingTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	DisconnectedTime := new(string)
	if o.DisconnectedTime != nil {
		
		*DisconnectedTime = timeutil.Strftime(o.DisconnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DisconnectedTime = nil
	}
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		InitialState *string `json:"initialState,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Context *string `json:"context,omitempty"`
		
		AudioMuted *bool `json:"audioMuted,omitempty"`
		
		VideoMuted *bool `json:"videoMuted,omitempty"`
		
		SharingScreen *bool `json:"sharingScreen,omitempty"`
		
		PeerCount *int `json:"peerCount,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartAlertingTime *string `json:"startAlertingTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		Msids *[]string `json:"msids,omitempty"`
		
		Self *Address `json:"self,omitempty"`
		
		Wrapup *Wrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		Alias
	}{ 
		State: o.State,
		
		InitialState: o.InitialState,
		
		Id: o.Id,
		
		Context: o.Context,
		
		AudioMuted: o.AudioMuted,
		
		VideoMuted: o.VideoMuted,
		
		SharingScreen: o.SharingScreen,
		
		PeerCount: o.PeerCount,
		
		DisconnectType: o.DisconnectType,
		
		StartAlertingTime: StartAlertingTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		Provider: o.Provider,
		
		PeerId: o.PeerId,
		
		Msids: o.Msids,
		
		Self: o.Self,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		Alias:    (Alias)(o),
	})
}

func (o *Video) UnmarshalJSON(b []byte) error {
	var VideoMap map[string]interface{}
	err := json.Unmarshal(b, &VideoMap)
	if err != nil {
		return err
	}
	
	if State, ok := VideoMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := VideoMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Id, ok := VideoMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Context, ok := VideoMap["context"].(string); ok {
		o.Context = &Context
	}
    
	if AudioMuted, ok := VideoMap["audioMuted"].(bool); ok {
		o.AudioMuted = &AudioMuted
	}
    
	if VideoMuted, ok := VideoMap["videoMuted"].(bool); ok {
		o.VideoMuted = &VideoMuted
	}
    
	if SharingScreen, ok := VideoMap["sharingScreen"].(bool); ok {
		o.SharingScreen = &SharingScreen
	}
    
	if PeerCount, ok := VideoMap["peerCount"].(float64); ok {
		PeerCountInt := int(PeerCount)
		o.PeerCount = &PeerCountInt
	}
	
	if DisconnectType, ok := VideoMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if startAlertingTimeString, ok := VideoMap["startAlertingTime"].(string); ok {
		StartAlertingTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAlertingTimeString)
		o.StartAlertingTime = &StartAlertingTime
	}
	
	if connectedTimeString, ok := VideoMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := VideoMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if Provider, ok := VideoMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if PeerId, ok := VideoMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if Msids, ok := VideoMap["msids"].([]interface{}); ok {
		MsidsString, _ := json.Marshal(Msids)
		json.Unmarshal(MsidsString, &o.Msids)
	}
	
	if Self, ok := VideoMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}
	
	if Wrapup, ok := VideoMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := VideoMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := VideoMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Video) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
