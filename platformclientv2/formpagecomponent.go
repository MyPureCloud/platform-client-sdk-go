package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Formpagecomponent - A component within a form page
type Formpagecomponent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FormComponentType - Type of the component
	FormComponentType *string `json:"formComponentType,omitempty"`

	// ListPicker - List picker configuration
	ListPicker *Formlistpicker `json:"listPicker,omitempty"`

	// DatePicker - Date picker configuration
	DatePicker *Formdatepicker `json:"datePicker,omitempty"`

	// Input - Input field configuration
	Input *Input `json:"input,omitempty"`

	// WheelPicker - Wheel picker configuration
	WheelPicker *Wheelpicker `json:"wheelPicker,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Formpagecomponent) SetField(field string, fieldValue interface{}) {
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

func (o Formpagecomponent) MarshalJSON() ([]byte, error) {
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
	type Alias Formpagecomponent
	
	return json.Marshal(&struct { 
		FormComponentType *string `json:"formComponentType,omitempty"`
		
		ListPicker *Formlistpicker `json:"listPicker,omitempty"`
		
		DatePicker *Formdatepicker `json:"datePicker,omitempty"`
		
		Input *Input `json:"input,omitempty"`
		
		WheelPicker *Wheelpicker `json:"wheelPicker,omitempty"`
		Alias
	}{ 
		FormComponentType: o.FormComponentType,
		
		ListPicker: o.ListPicker,
		
		DatePicker: o.DatePicker,
		
		Input: o.Input,
		
		WheelPicker: o.WheelPicker,
		Alias:    (Alias)(o),
	})
}

func (o *Formpagecomponent) UnmarshalJSON(b []byte) error {
	var FormpagecomponentMap map[string]interface{}
	err := json.Unmarshal(b, &FormpagecomponentMap)
	if err != nil {
		return err
	}
	
	if FormComponentType, ok := FormpagecomponentMap["formComponentType"].(string); ok {
		o.FormComponentType = &FormComponentType
	}
    
	if ListPicker, ok := FormpagecomponentMap["listPicker"].(map[string]interface{}); ok {
		ListPickerString, _ := json.Marshal(ListPicker)
		json.Unmarshal(ListPickerString, &o.ListPicker)
	}
	
	if DatePicker, ok := FormpagecomponentMap["datePicker"].(map[string]interface{}); ok {
		DatePickerString, _ := json.Marshal(DatePicker)
		json.Unmarshal(DatePickerString, &o.DatePicker)
	}
	
	if Input, ok := FormpagecomponentMap["input"].(map[string]interface{}); ok {
		InputString, _ := json.Marshal(Input)
		json.Unmarshal(InputString, &o.Input)
	}
	
	if WheelPicker, ok := FormpagecomponentMap["wheelPicker"].(map[string]interface{}); ok {
		WheelPickerString, _ := json.Marshal(WheelPicker)
		json.Unmarshal(WheelPickerString, &o.WheelPicker)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Formpagecomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
