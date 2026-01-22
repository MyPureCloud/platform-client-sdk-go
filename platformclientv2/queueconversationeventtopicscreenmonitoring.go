package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicscreenmonitoring
type Queueconversationeventtopicscreenmonitoring struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// InitialState
	InitialState *string `json:"initialState,omitempty"`

	// Provider - The source provider.
	Provider *string `json:"provider,omitempty"`

	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// ConnectedTime - The timestamp when this communication was connected in the cloud clock.
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock.
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`

	// TargetUserId - The user ID for the participant who is being screen monitored.
	TargetUserId *string `json:"targetUserId,omitempty"`

	// ScreenMonitoringId - Screen Monitoring identifier unique to the sourceUserId-targetUserId pair.
	ScreenMonitoringId *string `json:"screenMonitoringId,omitempty"`

	// MonitoringType - The screen monitoring type.
	MonitoringType *string `json:"monitoringType,omitempty"`

	// Count - Number of Screen Monitoring sessions the targetUserId is involved in.
	Count *int `json:"count,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueconversationeventtopicscreenmonitoring) SetField(field string, fieldValue interface{}) {
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

func (o Queueconversationeventtopicscreenmonitoring) MarshalJSON() ([]byte, error) {
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
	type Alias Queueconversationeventtopicscreenmonitoring
	
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
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		TargetUserId *string `json:"targetUserId,omitempty"`
		
		ScreenMonitoringId *string `json:"screenMonitoringId,omitempty"`
		
		MonitoringType *string `json:"monitoringType,omitempty"`
		
		Count *int `json:"count,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		InitialState: o.InitialState,
		
		Provider: o.Provider,
		
		DisconnectType: o.DisconnectType,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		TargetUserId: o.TargetUserId,
		
		ScreenMonitoringId: o.ScreenMonitoringId,
		
		MonitoringType: o.MonitoringType,
		
		Count: o.Count,
		Alias:    (Alias)(o),
	})
}

func (o *Queueconversationeventtopicscreenmonitoring) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopicscreenmonitoringMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopicscreenmonitoringMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationeventtopicscreenmonitoringMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := QueueconversationeventtopicscreenmonitoringMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := QueueconversationeventtopicscreenmonitoringMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Provider, ok := QueueconversationeventtopicscreenmonitoringMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if DisconnectType, ok := QueueconversationeventtopicscreenmonitoringMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if connectedTimeString, ok := QueueconversationeventtopicscreenmonitoringMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationeventtopicscreenmonitoringMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if TargetUserId, ok := QueueconversationeventtopicscreenmonitoringMap["targetUserId"].(string); ok {
		o.TargetUserId = &TargetUserId
	}
    
	if ScreenMonitoringId, ok := QueueconversationeventtopicscreenmonitoringMap["screenMonitoringId"].(string); ok {
		o.ScreenMonitoringId = &ScreenMonitoringId
	}
    
	if MonitoringType, ok := QueueconversationeventtopicscreenmonitoringMap["monitoringType"].(string); ok {
		o.MonitoringType = &MonitoringType
	}
    
	if Count, ok := QueueconversationeventtopicscreenmonitoringMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicscreenmonitoring) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
