package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingmessagereceipt
type Recordingmessagereceipt struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The id of the message receipt. Message receipts will have the same ID as the message they reference.
	Id *string `json:"id,omitempty"`

	// Status - The message receipt status
	Status *string `json:"status,omitempty"`

	// Reasons - List of reasons for a message receipt that indicates the message has failed. Only used with Failed status.
	Reasons *[]Recordingmessagereceiptreason `json:"reasons,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recordingmessagereceipt) SetField(field string, fieldValue interface{}) {
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

func (o Recordingmessagereceipt) MarshalJSON() ([]byte, error) {
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
	type Alias Recordingmessagereceipt
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Reasons *[]Recordingmessagereceiptreason `json:"reasons,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Status: o.Status,
		
		Reasons: o.Reasons,
		Alias:    (Alias)(o),
	})
}

func (o *Recordingmessagereceipt) UnmarshalJSON(b []byte) error {
	var RecordingmessagereceiptMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingmessagereceiptMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecordingmessagereceiptMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := RecordingmessagereceiptMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Reasons, ok := RecordingmessagereceiptMap["reasons"].([]interface{}); ok {
		ReasonsString, _ := json.Marshal(Reasons)
		json.Unmarshal(ReasonsString, &o.Reasons)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingmessagereceipt) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
