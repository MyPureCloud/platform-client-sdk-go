package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Screenmonitoring
type Screenmonitoring struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`

	// InitialState - The initial connection state of this communication.
	InitialState *string `json:"initialState,omitempty"`

	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`

	// Segments - The time line of the participant's Screen Monitoring media, divided into activity segments.
	Segments *[]Segment `json:"segments,omitempty"`

	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// Provider - The source provider of Screen Monitoring media.
	Provider *string `json:"provider,omitempty"`

	// TargetUser - The user participant who is being screen monitored.
	TargetUser *Addressableentityref `json:"targetUser,omitempty"`

	// ScreenMonitoring - Screen Monitoring identifier unique to the sourceUserId-targetUserId pair.
	ScreenMonitoring *Addressableentityref `json:"screenMonitoring,omitempty"`

	// MonitoringType - The screen monitoring type.
	MonitoringType *string `json:"monitoringType,omitempty"`

	// Count - Number of Screen Monitoring sessions the targetUserId is involved in.
	Count *int `json:"count,omitempty"`

	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Screenmonitoring) SetField(field string, fieldValue interface{}) {
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

func (o Screenmonitoring) MarshalJSON() ([]byte, error) {
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
	type Alias Screenmonitoring
	
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
		
		InitialState *string `json:"initialState,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Segments *[]Segment `json:"segments,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		TargetUser *Addressableentityref `json:"targetUser,omitempty"`
		
		ScreenMonitoring *Addressableentityref `json:"screenMonitoring,omitempty"`
		
		MonitoringType *string `json:"monitoringType,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		InitialState: o.InitialState,
		
		State: o.State,
		
		Segments: o.Segments,
		
		DisconnectType: o.DisconnectType,
		
		Provider: o.Provider,
		
		TargetUser: o.TargetUser,
		
		ScreenMonitoring: o.ScreenMonitoring,
		
		MonitoringType: o.MonitoringType,
		
		Count: o.Count,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		Alias:    (Alias)(o),
	})
}

func (o *Screenmonitoring) UnmarshalJSON(b []byte) error {
	var ScreenmonitoringMap map[string]interface{}
	err := json.Unmarshal(b, &ScreenmonitoringMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ScreenmonitoringMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if InitialState, ok := ScreenmonitoringMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if State, ok := ScreenmonitoringMap["state"].(string); ok {
		o.State = &State
	}
    
	if Segments, ok := ScreenmonitoringMap["segments"].([]interface{}); ok {
		SegmentsString, _ := json.Marshal(Segments)
		json.Unmarshal(SegmentsString, &o.Segments)
	}
	
	if DisconnectType, ok := ScreenmonitoringMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if Provider, ok := ScreenmonitoringMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if TargetUser, ok := ScreenmonitoringMap["targetUser"].(map[string]interface{}); ok {
		TargetUserString, _ := json.Marshal(TargetUser)
		json.Unmarshal(TargetUserString, &o.TargetUser)
	}
	
	if ScreenMonitoring, ok := ScreenmonitoringMap["screenMonitoring"].(map[string]interface{}); ok {
		ScreenMonitoringString, _ := json.Marshal(ScreenMonitoring)
		json.Unmarshal(ScreenMonitoringString, &o.ScreenMonitoring)
	}
	
	if MonitoringType, ok := ScreenmonitoringMap["monitoringType"].(string); ok {
		o.MonitoringType = &MonitoringType
	}
    
	if Count, ok := ScreenmonitoringMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if connectedTimeString, ok := ScreenmonitoringMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := ScreenmonitoringMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Screenmonitoring) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
