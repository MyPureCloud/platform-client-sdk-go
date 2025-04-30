package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2flowexecutiondataflowidtopicflowexecutionitem
type V2flowexecutiondataflowidtopicflowexecutionitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ObjectType - The type of executionItem that was executed.
	ObjectType *string `json:"objectType,omitempty"`

	// ObjectId - If applicable, the actionId, menuId or taskId for the executionItem.
	ObjectId *string `json:"objectId,omitempty"`

	// OutputPathId - If applicable, the identifier of the OutputPath that was taken.
	OutputPathId *string `json:"outputPathId,omitempty"`

	// ExecutionId - If applicable, the executionId for the executionItem.
	ExecutionId *string `json:"executionId,omitempty"`

	// StartDateTime - This is the starting time of the executionItem.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`

	// VarError - Event generated when a Flow's Execution History is received and logged.
	VarError *V2flowexecutiondataflowidtopicflowerrorwarninginfo `json:"error,omitempty"`

	// Warning - Event generated when a Flow's Execution History is received and logged.
	Warning *V2flowexecutiondataflowidtopicflowerrorwarninginfo `json:"warning,omitempty"`

	// LanguageTag - If applicable, the language tag associated set by the execution.
	LanguageTag *string `json:"languageTag,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2flowexecutiondataflowidtopicflowexecutionitem) SetField(field string, fieldValue interface{}) {
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

func (o V2flowexecutiondataflowidtopicflowexecutionitem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDateTime", }
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
	type Alias V2flowexecutiondataflowidtopicflowexecutionitem
	
	StartDateTime := new(string)
	if o.StartDateTime != nil {
		
		*StartDateTime = timeutil.Strftime(o.StartDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDateTime = nil
	}
	
	return json.Marshal(&struct { 
		ObjectType *string `json:"objectType,omitempty"`
		
		ObjectId *string `json:"objectId,omitempty"`
		
		OutputPathId *string `json:"outputPathId,omitempty"`
		
		ExecutionId *string `json:"executionId,omitempty"`
		
		StartDateTime *string `json:"startDateTime,omitempty"`
		
		VarError *V2flowexecutiondataflowidtopicflowerrorwarninginfo `json:"error,omitempty"`
		
		Warning *V2flowexecutiondataflowidtopicflowerrorwarninginfo `json:"warning,omitempty"`
		
		LanguageTag *string `json:"languageTag,omitempty"`
		Alias
	}{ 
		ObjectType: o.ObjectType,
		
		ObjectId: o.ObjectId,
		
		OutputPathId: o.OutputPathId,
		
		ExecutionId: o.ExecutionId,
		
		StartDateTime: StartDateTime,
		
		VarError: o.VarError,
		
		Warning: o.Warning,
		
		LanguageTag: o.LanguageTag,
		Alias:    (Alias)(o),
	})
}

func (o *V2flowexecutiondataflowidtopicflowexecutionitem) UnmarshalJSON(b []byte) error {
	var V2flowexecutiondataflowidtopicflowexecutionitemMap map[string]interface{}
	err := json.Unmarshal(b, &V2flowexecutiondataflowidtopicflowexecutionitemMap)
	if err != nil {
		return err
	}
	
	if ObjectType, ok := V2flowexecutiondataflowidtopicflowexecutionitemMap["objectType"].(string); ok {
		o.ObjectType = &ObjectType
	}
    
	if ObjectId, ok := V2flowexecutiondataflowidtopicflowexecutionitemMap["objectId"].(string); ok {
		o.ObjectId = &ObjectId
	}
    
	if OutputPathId, ok := V2flowexecutiondataflowidtopicflowexecutionitemMap["outputPathId"].(string); ok {
		o.OutputPathId = &OutputPathId
	}
    
	if ExecutionId, ok := V2flowexecutiondataflowidtopicflowexecutionitemMap["executionId"].(string); ok {
		o.ExecutionId = &ExecutionId
	}
    
	if startDateTimeString, ok := V2flowexecutiondataflowidtopicflowexecutionitemMap["startDateTime"].(string); ok {
		StartDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateTimeString)
		o.StartDateTime = &StartDateTime
	}
	
	if VarError, ok := V2flowexecutiondataflowidtopicflowexecutionitemMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if Warning, ok := V2flowexecutiondataflowidtopicflowexecutionitemMap["warning"].(map[string]interface{}); ok {
		WarningString, _ := json.Marshal(Warning)
		json.Unmarshal(WarningString, &o.Warning)
	}
	
	if LanguageTag, ok := V2flowexecutiondataflowidtopicflowexecutionitemMap["languageTag"].(string); ok {
		o.LanguageTag = &LanguageTag
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2flowexecutiondataflowidtopicflowexecutionitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
