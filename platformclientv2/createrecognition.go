package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createrecognition
type Createrecognition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RecipientId - The recipient of the recognition
	RecipientId *string `json:"recipientId,omitempty"`

	// VarType - The type of the recognition
	VarType *string `json:"type,omitempty"`

	// Title - The title of the recognition. Max length of 100 characters (optional)
	Title *string `json:"title,omitempty"`

	// Note - The note of the recognition. Max length of 800 characters (optional)
	Note *string `json:"note,omitempty"`

	// ContextType - The context type (optional)
	ContextType *string `json:"contextType,omitempty"`

	// ContextId - The context id (optional)
	ContextId *string `json:"contextId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createrecognition) SetField(field string, fieldValue interface{}) {
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

func (o Createrecognition) MarshalJSON() ([]byte, error) {
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
	type Alias Createrecognition
	
	return json.Marshal(&struct { 
		RecipientId *string `json:"recipientId,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Note *string `json:"note,omitempty"`
		
		ContextType *string `json:"contextType,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		Alias
	}{ 
		RecipientId: o.RecipientId,
		
		VarType: o.VarType,
		
		Title: o.Title,
		
		Note: o.Note,
		
		ContextType: o.ContextType,
		
		ContextId: o.ContextId,
		Alias:    (Alias)(o),
	})
}

func (o *Createrecognition) UnmarshalJSON(b []byte) error {
	var CreaterecognitionMap map[string]interface{}
	err := json.Unmarshal(b, &CreaterecognitionMap)
	if err != nil {
		return err
	}
	
	if RecipientId, ok := CreaterecognitionMap["recipientId"].(string); ok {
		o.RecipientId = &RecipientId
	}
    
	if VarType, ok := CreaterecognitionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Title, ok := CreaterecognitionMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Note, ok := CreaterecognitionMap["note"].(string); ok {
		o.Note = &Note
	}
    
	if ContextType, ok := CreaterecognitionMap["contextType"].(string); ok {
		o.ContextType = &ContextType
	}
    
	if ContextId, ok := CreaterecognitionMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createrecognition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
