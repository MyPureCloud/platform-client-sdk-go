package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowpathselement
type Flowpathselement struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ParentId - Unique identifier of the parent element. Will be null for the root element.
	ParentId *string `json:"parentId,omitempty"`

	// VarType - Type of the element.
	VarType *string `json:"type,omitempty"`

	// Count - Count of all journeys that include this element.
	Count *int `json:"count,omitempty"`

	// Flows - Details of flows involved in journeys that include this element.
	Flows *[]Flowpathsflowdetails `json:"flows,omitempty"`

	// FlowMilestone - The flow milestone, set if the element type is Milestone.
	FlowMilestone *Addressableentityref `json:"flowMilestone,omitempty"`

	// FlowOutcome - The flow outcome, set if the element type is Outcome or Milestone.
	FlowOutcome *Addressableentityref `json:"flowOutcome,omitempty"`

	// FlowOutcomeValue - The value of the flow outcome, if the element type is Outcome.
	FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowpathselement) SetField(field string, fieldValue interface{}) {
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

func (o Flowpathselement) MarshalJSON() ([]byte, error) {
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
	type Alias Flowpathselement
	
	return json.Marshal(&struct { 
		ParentId *string `json:"parentId,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		Flows *[]Flowpathsflowdetails `json:"flows,omitempty"`
		
		FlowMilestone *Addressableentityref `json:"flowMilestone,omitempty"`
		
		FlowOutcome *Addressableentityref `json:"flowOutcome,omitempty"`
		
		FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`
		Alias
	}{ 
		ParentId: o.ParentId,
		
		VarType: o.VarType,
		
		Count: o.Count,
		
		Flows: o.Flows,
		
		FlowMilestone: o.FlowMilestone,
		
		FlowOutcome: o.FlowOutcome,
		
		FlowOutcomeValue: o.FlowOutcomeValue,
		Alias:    (Alias)(o),
	})
}

func (o *Flowpathselement) UnmarshalJSON(b []byte) error {
	var FlowpathselementMap map[string]interface{}
	err := json.Unmarshal(b, &FlowpathselementMap)
	if err != nil {
		return err
	}
	
	if ParentId, ok := FlowpathselementMap["parentId"].(string); ok {
		o.ParentId = &ParentId
	}
    
	if VarType, ok := FlowpathselementMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Count, ok := FlowpathselementMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Flows, ok := FlowpathselementMap["flows"].([]interface{}); ok {
		FlowsString, _ := json.Marshal(Flows)
		json.Unmarshal(FlowsString, &o.Flows)
	}
	
	if FlowMilestone, ok := FlowpathselementMap["flowMilestone"].(map[string]interface{}); ok {
		FlowMilestoneString, _ := json.Marshal(FlowMilestone)
		json.Unmarshal(FlowMilestoneString, &o.FlowMilestone)
	}
	
	if FlowOutcome, ok := FlowpathselementMap["flowOutcome"].(map[string]interface{}); ok {
		FlowOutcomeString, _ := json.Marshal(FlowOutcome)
		json.Unmarshal(FlowOutcomeString, &o.FlowOutcome)
	}
	
	if FlowOutcomeValue, ok := FlowpathselementMap["flowOutcomeValue"].(string); ok {
		o.FlowOutcomeValue = &FlowOutcomeValue
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Flowpathselement) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
