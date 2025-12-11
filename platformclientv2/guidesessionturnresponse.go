package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Guidesessionturnresponse - Response for a guide session turn
type Guidesessionturnresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Response - The response content for this turn.
	Response *Guidesessionturnresponsedata `json:"response,omitempty"`

	// Status - The status of the turn.
	Status *string `json:"status,omitempty"`

	// Result - The result of the turn.
	Result *string `json:"result,omitempty"`

	// OutputVariables - The output variables for this turn.
	OutputVariables *[]Guidesessionvariable `json:"outputVariables,omitempty"`

	// InvocationId - Invocation ID for this turn.
	InvocationId *string `json:"invocationId,omitempty"`

	// Invocations - The invocations for this turn.
	Invocations *[]Guidesessionturninvocationresponse `json:"invocations,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Guidesessionturnresponse) SetField(field string, fieldValue interface{}) {
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

func (o Guidesessionturnresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Guidesessionturnresponse
	
	return json.Marshal(&struct { 
		Response *Guidesessionturnresponsedata `json:"response,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Result *string `json:"result,omitempty"`
		
		OutputVariables *[]Guidesessionvariable `json:"outputVariables,omitempty"`
		
		InvocationId *string `json:"invocationId,omitempty"`
		
		Invocations *[]Guidesessionturninvocationresponse `json:"invocations,omitempty"`
		Alias
	}{ 
		Response: o.Response,
		
		Status: o.Status,
		
		Result: o.Result,
		
		OutputVariables: o.OutputVariables,
		
		InvocationId: o.InvocationId,
		
		Invocations: o.Invocations,
		Alias:    (Alias)(o),
	})
}

func (o *Guidesessionturnresponse) UnmarshalJSON(b []byte) error {
	var GuidesessionturnresponseMap map[string]interface{}
	err := json.Unmarshal(b, &GuidesessionturnresponseMap)
	if err != nil {
		return err
	}
	
	if Response, ok := GuidesessionturnresponseMap["response"].(map[string]interface{}); ok {
		ResponseString, _ := json.Marshal(Response)
		json.Unmarshal(ResponseString, &o.Response)
	}
	
	if Status, ok := GuidesessionturnresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Result, ok := GuidesessionturnresponseMap["result"].(string); ok {
		o.Result = &Result
	}
    
	if OutputVariables, ok := GuidesessionturnresponseMap["outputVariables"].([]interface{}); ok {
		OutputVariablesString, _ := json.Marshal(OutputVariables)
		json.Unmarshal(OutputVariablesString, &o.OutputVariables)
	}
	
	if InvocationId, ok := GuidesessionturnresponseMap["invocationId"].(string); ok {
		o.InvocationId = &InvocationId
	}
    
	if Invocations, ok := GuidesessionturnresponseMap["invocations"].([]interface{}); ok {
		InvocationsString, _ := json.Marshal(Invocations)
		json.Unmarshal(InvocationsString, &o.Invocations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Guidesessionturnresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
