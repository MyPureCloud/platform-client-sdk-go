package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicbuschedulenotification
type Wfmbuscheduletopicbuschedulenotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Status
	Status *string `json:"status,omitempty"`

	// OperationId
	OperationId *string `json:"operationId,omitempty"`

	// EventType
	EventType *string `json:"eventType,omitempty"`

	// Result
	Result *Wfmbuscheduletopicbuschedulemetadata `json:"result,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmbuscheduletopicbuschedulenotification) SetField(field string, fieldValue interface{}) {
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

func (o Wfmbuscheduletopicbuschedulenotification) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmbuscheduletopicbuschedulenotification
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		Result *Wfmbuscheduletopicbuschedulemetadata `json:"result,omitempty"`
		Alias
	}{ 
		Status: o.Status,
		
		OperationId: o.OperationId,
		
		EventType: o.EventType,
		
		Result: o.Result,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmbuscheduletopicbuschedulenotification) UnmarshalJSON(b []byte) error {
	var WfmbuscheduletopicbuschedulenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduletopicbuschedulenotificationMap)
	if err != nil {
		return err
	}
	
	if Status, ok := WfmbuscheduletopicbuschedulenotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if OperationId, ok := WfmbuscheduletopicbuschedulenotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if EventType, ok := WfmbuscheduletopicbuschedulenotificationMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if Result, ok := WfmbuscheduletopicbuschedulenotificationMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbuschedulenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
