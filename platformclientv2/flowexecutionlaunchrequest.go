package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowexecutionlaunchrequest - Parameters for launching a flow.
type Flowexecutionlaunchrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FlowId - ID of the flow to launch.
	FlowId *string `json:"flowId,omitempty"`

	// FlowVersion - The version of the flow to launch. Omit this value (or supply null/empty) to use the latest published version.
	FlowVersion *string `json:"flowVersion,omitempty"`

	// InputData - Input values to the flow. Valid values are defined by a flow's input JSON schema.
	InputData *map[string]interface{} `json:"inputData,omitempty"`

	// Name - A displayable name to assign to the new flow execution
	Name *string `json:"name,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowexecutionlaunchrequest) SetField(field string, fieldValue interface{}) {
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

func (o Flowexecutionlaunchrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Flowexecutionlaunchrequest
	
	return json.Marshal(&struct { 
		FlowId *string `json:"flowId,omitempty"`
		
		FlowVersion *string `json:"flowVersion,omitempty"`
		
		InputData *map[string]interface{} `json:"inputData,omitempty"`
		
		Name *string `json:"name,omitempty"`
		Alias
	}{ 
		FlowId: o.FlowId,
		
		FlowVersion: o.FlowVersion,
		
		InputData: o.InputData,
		
		Name: o.Name,
		Alias:    (Alias)(o),
	})
}

func (o *Flowexecutionlaunchrequest) UnmarshalJSON(b []byte) error {
	var FlowexecutionlaunchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &FlowexecutionlaunchrequestMap)
	if err != nil {
		return err
	}
	
	if FlowId, ok := FlowexecutionlaunchrequestMap["flowId"].(string); ok {
		o.FlowId = &FlowId
	}
    
	if FlowVersion, ok := FlowexecutionlaunchrequestMap["flowVersion"].(string); ok {
		o.FlowVersion = &FlowVersion
	}
    
	if InputData, ok := FlowexecutionlaunchrequestMap["inputData"].(map[string]interface{}); ok {
		InputDataString, _ := json.Marshal(InputData)
		json.Unmarshal(InputDataString, &o.InputData)
	}
	
	if Name, ok := FlowexecutionlaunchrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Flowexecutionlaunchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
