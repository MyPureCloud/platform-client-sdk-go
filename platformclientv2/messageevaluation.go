package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messageevaluation
type Messageevaluation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContactColumn - The name of the contact column that was wrapped up
	ContactColumn *string `json:"contactColumn,omitempty"`

	// ContactAddress - The address (phone or email) that was wrapped up
	ContactAddress *string `json:"contactAddress,omitempty"`

	// MessageType - The type of message sent
	MessageType *string `json:"messageType,omitempty"`

	// WrapupCodeId - The id of the wrap-up code
	WrapupCodeId *string `json:"wrapupCodeId,omitempty"`

	// Timestamp - The time that the wrap-up was applied. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messageevaluation) SetField(field string, fieldValue interface{}) {
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

func (o Messageevaluation) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Timestamp", }
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
	type Alias Messageevaluation
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	
	return json.Marshal(&struct { 
		ContactColumn *string `json:"contactColumn,omitempty"`
		
		ContactAddress *string `json:"contactAddress,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		WrapupCodeId *string `json:"wrapupCodeId,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		Alias
	}{ 
		ContactColumn: o.ContactColumn,
		
		ContactAddress: o.ContactAddress,
		
		MessageType: o.MessageType,
		
		WrapupCodeId: o.WrapupCodeId,
		
		Timestamp: Timestamp,
		Alias:    (Alias)(o),
	})
}

func (o *Messageevaluation) UnmarshalJSON(b []byte) error {
	var MessageevaluationMap map[string]interface{}
	err := json.Unmarshal(b, &MessageevaluationMap)
	if err != nil {
		return err
	}
	
	if ContactColumn, ok := MessageevaluationMap["contactColumn"].(string); ok {
		o.ContactColumn = &ContactColumn
	}
    
	if ContactAddress, ok := MessageevaluationMap["contactAddress"].(string); ok {
		o.ContactAddress = &ContactAddress
	}
    
	if MessageType, ok := MessageevaluationMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if WrapupCodeId, ok := MessageevaluationMap["wrapupCodeId"].(string); ok {
		o.WrapupCodeId = &WrapupCodeId
	}
    
	if timestampString, ok := MessageevaluationMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messageevaluation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
