package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemsattributechangeinstant
type Workitemsattributechangeinstant struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// NewValue - New property value. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	NewValue *time.Time `json:"newValue,omitempty"`

	// OldValue - Old property value. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	OldValue *time.Time `json:"oldValue,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workitemsattributechangeinstant) SetField(field string, fieldValue interface{}) {
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

func (o Workitemsattributechangeinstant) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "NewValue","OldValue", }
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
	type Alias Workitemsattributechangeinstant
	
	NewValue := new(string)
	if o.NewValue != nil {
		
		*NewValue = timeutil.Strftime(o.NewValue, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		NewValue = nil
	}
	
	OldValue := new(string)
	if o.OldValue != nil {
		
		*OldValue = timeutil.Strftime(o.OldValue, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		OldValue = nil
	}
	
	return json.Marshal(&struct { 
		NewValue *string `json:"newValue,omitempty"`
		
		OldValue *string `json:"oldValue,omitempty"`
		Alias
	}{ 
		NewValue: NewValue,
		
		OldValue: OldValue,
		Alias:    (Alias)(o),
	})
}

func (o *Workitemsattributechangeinstant) UnmarshalJSON(b []byte) error {
	var WorkitemsattributechangeinstantMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemsattributechangeinstantMap)
	if err != nil {
		return err
	}
	
	if newValueString, ok := WorkitemsattributechangeinstantMap["newValue"].(string); ok {
		NewValue, _ := time.Parse("2006-01-02T15:04:05.999999Z", newValueString)
		o.NewValue = &NewValue
	}
	
	if oldValueString, ok := WorkitemsattributechangeinstantMap["oldValue"].(string); ok {
		OldValue, _ := time.Parse("2006-01-02T15:04:05.999999Z", oldValueString)
		o.OldValue = &OldValue
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemsattributechangeinstant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
