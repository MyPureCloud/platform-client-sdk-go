package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Presenceeventuserpresence
type Presenceeventuserpresence struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Source
	Source *string `json:"source,omitempty"`

	// PresenceDefinition
	PresenceDefinition *Presenceeventorganizationpresence `json:"presenceDefinition,omitempty"`

	// Primary
	Primary *bool `json:"primary,omitempty"`

	// Message
	Message *string `json:"message,omitempty"`

	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Presenceeventuserpresence) SetField(field string, fieldValue interface{}) {
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

func (o Presenceeventuserpresence) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ModifiedDate", }
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
	type Alias Presenceeventuserpresence
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Source *string `json:"source,omitempty"`
		
		PresenceDefinition *Presenceeventorganizationpresence `json:"presenceDefinition,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		Alias
	}{ 
		Source: o.Source,
		
		PresenceDefinition: o.PresenceDefinition,
		
		Primary: o.Primary,
		
		Message: o.Message,
		
		ModifiedDate: ModifiedDate,
		Alias:    (Alias)(o),
	})
}

func (o *Presenceeventuserpresence) UnmarshalJSON(b []byte) error {
	var PresenceeventuserpresenceMap map[string]interface{}
	err := json.Unmarshal(b, &PresenceeventuserpresenceMap)
	if err != nil {
		return err
	}
	
	if Source, ok := PresenceeventuserpresenceMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if PresenceDefinition, ok := PresenceeventuserpresenceMap["presenceDefinition"].(map[string]interface{}); ok {
		PresenceDefinitionString, _ := json.Marshal(PresenceDefinition)
		json.Unmarshal(PresenceDefinitionString, &o.PresenceDefinition)
	}
	
	if Primary, ok := PresenceeventuserpresenceMap["primary"].(bool); ok {
		o.Primary = &Primary
	}
    
	if Message, ok := PresenceeventuserpresenceMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if modifiedDateString, ok := PresenceeventuserpresenceMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Presenceeventuserpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
