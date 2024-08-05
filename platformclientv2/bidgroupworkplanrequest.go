package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Bidgroupworkplanrequest
type Bidgroupworkplanrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// WorkPlanId - The ID of the work plan used in the bid group
	WorkPlanId *string `json:"workPlanId,omitempty"`

	// OverrideAgentCount - The modified agent count for this work plan
	OverrideAgentCount *int `json:"overrideAgentCount,omitempty"`

	// SuggestedAgentCount - The number of agents needed for this work plan to produce the optimal schedule
	SuggestedAgentCount *int `json:"suggestedAgentCount,omitempty"`

	// AgentCountRange - The range of agent slot count per work plan. The suggested slot count must be in agent count range
	AgentCountRange *Agentcountrange `json:"agentCountRange,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Bidgroupworkplanrequest) SetField(field string, fieldValue interface{}) {
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

func (o Bidgroupworkplanrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Bidgroupworkplanrequest
	
	return json.Marshal(&struct { 
		WorkPlanId *string `json:"workPlanId,omitempty"`
		
		OverrideAgentCount *int `json:"overrideAgentCount,omitempty"`
		
		SuggestedAgentCount *int `json:"suggestedAgentCount,omitempty"`
		
		AgentCountRange *Agentcountrange `json:"agentCountRange,omitempty"`
		Alias
	}{ 
		WorkPlanId: o.WorkPlanId,
		
		OverrideAgentCount: o.OverrideAgentCount,
		
		SuggestedAgentCount: o.SuggestedAgentCount,
		
		AgentCountRange: o.AgentCountRange,
		Alias:    (Alias)(o),
	})
}

func (o *Bidgroupworkplanrequest) UnmarshalJSON(b []byte) error {
	var BidgroupworkplanrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BidgroupworkplanrequestMap)
	if err != nil {
		return err
	}
	
	if WorkPlanId, ok := BidgroupworkplanrequestMap["workPlanId"].(string); ok {
		o.WorkPlanId = &WorkPlanId
	}
    
	if OverrideAgentCount, ok := BidgroupworkplanrequestMap["overrideAgentCount"].(float64); ok {
		OverrideAgentCountInt := int(OverrideAgentCount)
		o.OverrideAgentCount = &OverrideAgentCountInt
	}
	
	if SuggestedAgentCount, ok := BidgroupworkplanrequestMap["suggestedAgentCount"].(float64); ok {
		SuggestedAgentCountInt := int(SuggestedAgentCount)
		o.SuggestedAgentCount = &SuggestedAgentCountInt
	}
	
	if AgentCountRange, ok := BidgroupworkplanrequestMap["agentCountRange"].(map[string]interface{}); ok {
		AgentCountRangeString, _ := json.Marshal(AgentCountRange)
		json.Unmarshal(AgentCountRangeString, &o.AgentCountRange)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bidgroupworkplanrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
