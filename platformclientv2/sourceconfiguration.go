package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Sourceconfiguration
type Sourceconfiguration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SourceId - Identifies the external platform that is the source of the conversation.
	SourceId *string `json:"sourceId,omitempty"`

	// InteractionId - The customer's unique external identifier associated with the conversation that comes from the external platform.
	InteractionId *string `json:"interactionId,omitempty"`

	// TagId - The customer's external identifier or tag associated with the conversation. If set, it will be used to tag the conversation.
	TagId *string `json:"tagId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Sourceconfiguration) SetField(field string, fieldValue interface{}) {
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

func (o Sourceconfiguration) MarshalJSON() ([]byte, error) {
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
	type Alias Sourceconfiguration
	
	return json.Marshal(&struct { 
		SourceId *string `json:"sourceId,omitempty"`
		
		InteractionId *string `json:"interactionId,omitempty"`
		
		TagId *string `json:"tagId,omitempty"`
		Alias
	}{ 
		SourceId: o.SourceId,
		
		InteractionId: o.InteractionId,
		
		TagId: o.TagId,
		Alias:    (Alias)(o),
	})
}

func (o *Sourceconfiguration) UnmarshalJSON(b []byte) error {
	var SourceconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &SourceconfigurationMap)
	if err != nil {
		return err
	}
	
	if SourceId, ok := SourceconfigurationMap["sourceId"].(string); ok {
		o.SourceId = &SourceId
	}
    
	if InteractionId, ok := SourceconfigurationMap["interactionId"].(string); ok {
		o.InteractionId = &InteractionId
	}
    
	if TagId, ok := SourceconfigurationMap["tagId"].(string); ok {
		o.TagId = &TagId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Sourceconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
