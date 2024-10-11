package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicinternalmessage
type Queueconversationsocialexpressioneventtopicinternalmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// InitialState
	InitialState *string `json:"initialState,omitempty"`

	// Provider - The source provider of the message.
	Provider *string `json:"provider,omitempty"`

	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`

	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// ConnectedTime - The timestamp when this communication was connected in the cloud clock.
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock.
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`

	// TargetUserId - The user ID for the participant on receiving side of the internal message conversation.
	TargetUserId *string `json:"targetUserId,omitempty"`

	// SourceUserId - The user ID for the participant on sending side of the internal message conversation.
	SourceUserId *string `json:"sourceUserId,omitempty"`

	// ToAddress - Address and name data for a call endpoint.
	ToAddress *Queueconversationsocialexpressioneventtopicaddress `json:"toAddress,omitempty"`

	// FromAddress - Address and name data for a call endpoint.
	FromAddress *Queueconversationsocialexpressioneventtopicaddress `json:"fromAddress,omitempty"`

	// Messages - The messages sent on this communication channel.
	Messages *[]Queueconversationsocialexpressioneventtopicinternalmessagedetails `json:"messages,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueconversationsocialexpressioneventtopicinternalmessage) SetField(field string, fieldValue interface{}) {
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

func (o Queueconversationsocialexpressioneventtopicinternalmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ConnectedTime","DisconnectedTime", }
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
	type Alias Queueconversationsocialexpressioneventtopicinternalmessage
	
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
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		InitialState *string `json:"initialState,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		TargetUserId *string `json:"targetUserId,omitempty"`
		
		SourceUserId *string `json:"sourceUserId,omitempty"`
		
		ToAddress *Queueconversationsocialexpressioneventtopicaddress `json:"toAddress,omitempty"`
		
		FromAddress *Queueconversationsocialexpressioneventtopicaddress `json:"fromAddress,omitempty"`
		
		Messages *[]Queueconversationsocialexpressioneventtopicinternalmessagedetails `json:"messages,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		InitialState: o.InitialState,
		
		Provider: o.Provider,
		
		PeerId: o.PeerId,
		
		DisconnectType: o.DisconnectType,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		TargetUserId: o.TargetUserId,
		
		SourceUserId: o.SourceUserId,
		
		ToAddress: o.ToAddress,
		
		FromAddress: o.FromAddress,
		
		Messages: o.Messages,
		Alias:    (Alias)(o),
	})
}

func (o *Queueconversationsocialexpressioneventtopicinternalmessage) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicinternalmessageMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicinternalmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Provider, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if PeerId, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if DisconnectType, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if connectedTimeString, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if TargetUserId, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["targetUserId"].(string); ok {
		o.TargetUserId = &TargetUserId
	}
    
	if SourceUserId, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["sourceUserId"].(string); ok {
		o.SourceUserId = &SourceUserId
	}
    
	if ToAddress, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["toAddress"].(map[string]interface{}); ok {
		ToAddressString, _ := json.Marshal(ToAddress)
		json.Unmarshal(ToAddressString, &o.ToAddress)
	}
	
	if FromAddress, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["fromAddress"].(map[string]interface{}); ok {
		FromAddressString, _ := json.Marshal(FromAddress)
		json.Unmarshal(FromAddressString, &o.FromAddress)
	}
	
	if Messages, ok := QueueconversationsocialexpressioneventtopicinternalmessageMap["messages"].([]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicinternalmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
