package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowsquerycriteriaresponse - The response for QueryCapabilities which contains the allowed criteria, flow types and action types for the organization.
type Flowsquerycriteriaresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Criteria - The is a list of allowed criteria to query on.
	Criteria *[]Querycriteria `json:"criteria,omitempty"`

	// FlowTypes - The is a list of flow types the organization has access to.
	FlowTypes *[]string `json:"flowTypes,omitempty"`

	// ActionTypes - The is a list of action types the organization has access to.
	ActionTypes *[]string `json:"actionTypes,omitempty"`

	// ErrorCodes - The is a list of potential error codes the organization may encounter.
	ErrorCodes *[]string `json:"errorCodes,omitempty"`

	// WarningCodes - The is a list of potential warning codes the organization may encounter.
	WarningCodes *[]string `json:"warningCodes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowsquerycriteriaresponse) SetField(field string, fieldValue interface{}) {
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

func (o Flowsquerycriteriaresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Flowsquerycriteriaresponse
	
	return json.Marshal(&struct { 
		Criteria *[]Querycriteria `json:"criteria,omitempty"`
		
		FlowTypes *[]string `json:"flowTypes,omitempty"`
		
		ActionTypes *[]string `json:"actionTypes,omitempty"`
		
		ErrorCodes *[]string `json:"errorCodes,omitempty"`
		
		WarningCodes *[]string `json:"warningCodes,omitempty"`
		Alias
	}{ 
		Criteria: o.Criteria,
		
		FlowTypes: o.FlowTypes,
		
		ActionTypes: o.ActionTypes,
		
		ErrorCodes: o.ErrorCodes,
		
		WarningCodes: o.WarningCodes,
		Alias:    (Alias)(o),
	})
}

func (o *Flowsquerycriteriaresponse) UnmarshalJSON(b []byte) error {
	var FlowsquerycriteriaresponseMap map[string]interface{}
	err := json.Unmarshal(b, &FlowsquerycriteriaresponseMap)
	if err != nil {
		return err
	}
	
	if Criteria, ok := FlowsquerycriteriaresponseMap["criteria"].([]interface{}); ok {
		CriteriaString, _ := json.Marshal(Criteria)
		json.Unmarshal(CriteriaString, &o.Criteria)
	}
	
	if FlowTypes, ok := FlowsquerycriteriaresponseMap["flowTypes"].([]interface{}); ok {
		FlowTypesString, _ := json.Marshal(FlowTypes)
		json.Unmarshal(FlowTypesString, &o.FlowTypes)
	}
	
	if ActionTypes, ok := FlowsquerycriteriaresponseMap["actionTypes"].([]interface{}); ok {
		ActionTypesString, _ := json.Marshal(ActionTypes)
		json.Unmarshal(ActionTypesString, &o.ActionTypes)
	}
	
	if ErrorCodes, ok := FlowsquerycriteriaresponseMap["errorCodes"].([]interface{}); ok {
		ErrorCodesString, _ := json.Marshal(ErrorCodes)
		json.Unmarshal(ErrorCodesString, &o.ErrorCodes)
	}
	
	if WarningCodes, ok := FlowsquerycriteriaresponseMap["warningCodes"].([]interface{}); ok {
		WarningCodesString, _ := json.Marshal(WarningCodes)
		json.Unmarshal(WarningCodesString, &o.WarningCodes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowsquerycriteriaresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
