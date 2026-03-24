package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Caseassociationquery
type Caseassociationquery struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PageSize - The number of entities to return in the response.
	PageSize *int `json:"pageSize,omitempty"`

	// After - The cursor that points to the end of the set of entities that has been returned.
	After *string `json:"after,omitempty"`

	// WorkitemId - The Workitem ID to query by.
	WorkitemId *string `json:"workitemId,omitempty"`

	// ConversationId - The conversation ID to query by.
	ConversationId *string `json:"conversationId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Caseassociationquery) SetField(field string, fieldValue interface{}) {
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

func (o Caseassociationquery) MarshalJSON() ([]byte, error) {
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
	type Alias Caseassociationquery
	
	return json.Marshal(&struct { 
		PageSize *int `json:"pageSize,omitempty"`
		
		After *string `json:"after,omitempty"`
		
		WorkitemId *string `json:"workitemId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		Alias
	}{ 
		PageSize: o.PageSize,
		
		After: o.After,
		
		WorkitemId: o.WorkitemId,
		
		ConversationId: o.ConversationId,
		Alias:    (Alias)(o),
	})
}

func (o *Caseassociationquery) UnmarshalJSON(b []byte) error {
	var CaseassociationqueryMap map[string]interface{}
	err := json.Unmarshal(b, &CaseassociationqueryMap)
	if err != nil {
		return err
	}
	
	if PageSize, ok := CaseassociationqueryMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if After, ok := CaseassociationqueryMap["after"].(string); ok {
		o.After = &After
	}
    
	if WorkitemId, ok := CaseassociationqueryMap["workitemId"].(string); ok {
		o.WorkitemId = &WorkitemId
	}
    
	if ConversationId, ok := CaseassociationqueryMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Caseassociationquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
