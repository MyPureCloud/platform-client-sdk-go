package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capabilitiesactionlimitinfo
type Capabilitiesactionlimitinfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FileAttachmentsPerResponseMax - The maximum number of file attachments allowed per response
	FileAttachmentsPerResponseMax *Capabilitieslimitcondition `json:"fileAttachmentsPerResponseMax,omitempty"`

	// InvocationsPerFlowMax - The maximum number of times an action can be executed per flow instance
	InvocationsPerFlowMax *Capabilitieslimitcondition `json:"invocationsPerFlowMax,omitempty"`

	// LoopCountMax - The maximum number of iterations allowed per loop action instance
	LoopCountMax *Capabilitieslimitcondition `json:"loopCountMax,omitempty"`

	// NotificationTypeChatInvocationsPerFlowMax - The maximum number of times a Send Notification action can be executed per flow instance for chat notifications
	NotificationTypeChatInvocationsPerFlowMax *Capabilitieslimitcondition `json:"notificationTypeChatInvocationsPerFlowMax,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capabilitiesactionlimitinfo) SetField(field string, fieldValue interface{}) {
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

func (o Capabilitiesactionlimitinfo) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Capabilitiesactionlimitinfo
	
	return json.Marshal(&struct { 
		FileAttachmentsPerResponseMax *Capabilitieslimitcondition `json:"fileAttachmentsPerResponseMax,omitempty"`
		
		InvocationsPerFlowMax *Capabilitieslimitcondition `json:"invocationsPerFlowMax,omitempty"`
		
		LoopCountMax *Capabilitieslimitcondition `json:"loopCountMax,omitempty"`
		
		NotificationTypeChatInvocationsPerFlowMax *Capabilitieslimitcondition `json:"notificationTypeChatInvocationsPerFlowMax,omitempty"`
		Alias
	}{ 
		FileAttachmentsPerResponseMax: o.FileAttachmentsPerResponseMax,
		
		InvocationsPerFlowMax: o.InvocationsPerFlowMax,
		
		LoopCountMax: o.LoopCountMax,
		
		NotificationTypeChatInvocationsPerFlowMax: o.NotificationTypeChatInvocationsPerFlowMax,
		Alias:    (Alias)(o),
	})
}

func (o *Capabilitiesactionlimitinfo) UnmarshalJSON(b []byte) error {
	var CapabilitiesactionlimitinfoMap map[string]interface{}
	err := json.Unmarshal(b, &CapabilitiesactionlimitinfoMap)
	if err != nil {
		return err
	}
	
	if FileAttachmentsPerResponseMax, ok := CapabilitiesactionlimitinfoMap["fileAttachmentsPerResponseMax"].(map[string]interface{}); ok {
		FileAttachmentsPerResponseMaxString, _ := json.Marshal(FileAttachmentsPerResponseMax)
		json.Unmarshal(FileAttachmentsPerResponseMaxString, &o.FileAttachmentsPerResponseMax)
	}
	
	if InvocationsPerFlowMax, ok := CapabilitiesactionlimitinfoMap["invocationsPerFlowMax"].(map[string]interface{}); ok {
		InvocationsPerFlowMaxString, _ := json.Marshal(InvocationsPerFlowMax)
		json.Unmarshal(InvocationsPerFlowMaxString, &o.InvocationsPerFlowMax)
	}
	
	if LoopCountMax, ok := CapabilitiesactionlimitinfoMap["loopCountMax"].(map[string]interface{}); ok {
		LoopCountMaxString, _ := json.Marshal(LoopCountMax)
		json.Unmarshal(LoopCountMaxString, &o.LoopCountMax)
	}
	
	if NotificationTypeChatInvocationsPerFlowMax, ok := CapabilitiesactionlimitinfoMap["notificationTypeChatInvocationsPerFlowMax"].(map[string]interface{}); ok {
		NotificationTypeChatInvocationsPerFlowMaxString, _ := json.Marshal(NotificationTypeChatInvocationsPerFlowMax)
		json.Unmarshal(NotificationTypeChatInvocationsPerFlowMaxString, &o.NotificationTypeChatInvocationsPerFlowMax)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capabilitiesactionlimitinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
