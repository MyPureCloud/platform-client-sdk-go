package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Botentityvalue
type Botentityvalue struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`

	// VarType - The data type of the entity. Valid types include: String, Integer, Decimal, Boolean, Duration, Datetime, Currency, StringCollection, IntegerCollection, DecimalCollection, BooleanCollection, DurationCollection, DatetimeCollection, CurrencyCollection.
	VarType *string `json:"type,omitempty"`

	// Value - The string value of the entity for simple types. Required when using non-collection types. Format depends on the 'type' field: String: \"any text\"; Integer: whole number (e.g., \"42\"); Decimal: number with optional decimal point (e.g., \"42.5\"); Boolean: \"true\" or \"false\"; Duration: ISO-8601 duration format (e.g., \"PT1H30M\" for 1 hour and 30 minutes); Datetime: ISO-8601 datetime format (e.g., \"2023-04-15T14:30:00Z\"); Currency: JSON object with 'amount' and 'code' fields as an escaped JSON string (e.g., \"{\\\"amount\\\":10.50,\\\"code\\\":\\\"USD\\\"}\" - note the escaped quotes).
	Value *string `json:"value,omitempty"`

	// Values - The collection values for collection types. Required when using collection types. Each value must follow the format of its base type: StringCollection: array of strings; IntegerCollection: array of integer strings (e.g., [\"1\", \"2\", \"3\"]); DecimalCollection: array of decimal strings (e.g., [\"1.5\", \"2.0\", \"3.75\"]); BooleanCollection: array of boolean strings (e.g., [\"true\", \"false\"]); DurationCollection: array of ISO-8601 duration strings; DatetimeCollection: array of ISO-8601 datetime strings; CurrencyCollection: array of escaped JSON currency object strings (e.g., [\"{\\\"amount\\\":10.50,\\\"code\\\":\\\"USD\\\"}\", \"{\\\"amount\\\":25.00,\\\"code\\\":\\\"EUR\\\"}\"] - note the escaped quotes).
	Values *[]string `json:"values,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Botentityvalue) SetField(field string, fieldValue interface{}) {
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

func (o Botentityvalue) MarshalJSON() ([]byte, error) {
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
	type Alias Botentityvalue
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		VarType: o.VarType,
		
		Value: o.Value,
		
		Values: o.Values,
		Alias:    (Alias)(o),
	})
}

func (o *Botentityvalue) UnmarshalJSON(b []byte) error {
	var BotentityvalueMap map[string]interface{}
	err := json.Unmarshal(b, &BotentityvalueMap)
	if err != nil {
		return err
	}
	
	if Name, ok := BotentityvalueMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := BotentityvalueMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := BotentityvalueMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Values, ok := BotentityvalueMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Botentityvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
