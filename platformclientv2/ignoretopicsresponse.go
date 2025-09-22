package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Ignoretopicsresponse
type Ignoretopicsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TotalTopics - Total number of topics in current org
	TotalTopics *int `json:"totalTopics,omitempty"`

	// AddedTopics - Number of topics added in current request
	AddedTopics *int `json:"addedTopics,omitempty"`

	// UpdatedTopics - Number of topics updated in current request
	UpdatedTopics *int `json:"updatedTopics,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Ignoretopicsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Ignoretopicsresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Ignoretopicsresponse
	
	return json.Marshal(&struct { 
		TotalTopics *int `json:"totalTopics,omitempty"`
		
		AddedTopics *int `json:"addedTopics,omitempty"`
		
		UpdatedTopics *int `json:"updatedTopics,omitempty"`
		Alias
	}{ 
		TotalTopics: o.TotalTopics,
		
		AddedTopics: o.AddedTopics,
		
		UpdatedTopics: o.UpdatedTopics,
		Alias:    (Alias)(o),
	})
}

func (o *Ignoretopicsresponse) UnmarshalJSON(b []byte) error {
	var IgnoretopicsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &IgnoretopicsresponseMap)
	if err != nil {
		return err
	}
	
	if TotalTopics, ok := IgnoretopicsresponseMap["totalTopics"].(float64); ok {
		TotalTopicsInt := int(TotalTopics)
		o.TotalTopics = &TotalTopicsInt
	}
	
	if AddedTopics, ok := IgnoretopicsresponseMap["addedTopics"].(float64); ok {
		AddedTopicsInt := int(AddedTopics)
		o.AddedTopics = &AddedTopicsInt
	}
	
	if UpdatedTopics, ok := IgnoretopicsresponseMap["updatedTopics"].(float64); ok {
		UpdatedTopicsInt := int(UpdatedTopics)
		o.UpdatedTopics = &UpdatedTopicsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ignoretopicsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
