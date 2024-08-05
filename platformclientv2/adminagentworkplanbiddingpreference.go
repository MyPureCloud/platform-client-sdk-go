package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Adminagentworkplanbiddingpreference
type Adminagentworkplanbiddingpreference struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Agent - The agent to whom this work plan bidding preference applies
	Agent *Userreference `json:"agent,omitempty"`

	// Submitted - Whether the preference is submitted
	Submitted *bool `json:"submitted,omitempty"`

	// AssignedWorkPlan - The work plan assigned to the agent by the bid process
	AssignedWorkPlan *Workplanreference `json:"assignedWorkPlan,omitempty"`

	// OverriddenWorkPlan - The work plan that overrides the assigned work plan for the agent
	OverriddenWorkPlan *Workplanreference `json:"overriddenWorkPlan,omitempty"`

	// OverrideReason - The reason why the assigned work plan has been overridden. This must be null without an override work plan
	OverrideReason *string `json:"overrideReason,omitempty"`

	// Priorities - The agent priorities for the list of work plans. The index of the priorities should match with the list of work plans that belong to bid group. It contains null if priority is not set for the work plan
	Priorities *[]int `json:"priorities,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Adminagentworkplanbiddingpreference) SetField(field string, fieldValue interface{}) {
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

func (o Adminagentworkplanbiddingpreference) MarshalJSON() ([]byte, error) {
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
	type Alias Adminagentworkplanbiddingpreference
	
	return json.Marshal(&struct { 
		Agent *Userreference `json:"agent,omitempty"`
		
		Submitted *bool `json:"submitted,omitempty"`
		
		AssignedWorkPlan *Workplanreference `json:"assignedWorkPlan,omitempty"`
		
		OverriddenWorkPlan *Workplanreference `json:"overriddenWorkPlan,omitempty"`
		
		OverrideReason *string `json:"overrideReason,omitempty"`
		
		Priorities *[]int `json:"priorities,omitempty"`
		Alias
	}{ 
		Agent: o.Agent,
		
		Submitted: o.Submitted,
		
		AssignedWorkPlan: o.AssignedWorkPlan,
		
		OverriddenWorkPlan: o.OverriddenWorkPlan,
		
		OverrideReason: o.OverrideReason,
		
		Priorities: o.Priorities,
		Alias:    (Alias)(o),
	})
}

func (o *Adminagentworkplanbiddingpreference) UnmarshalJSON(b []byte) error {
	var AdminagentworkplanbiddingpreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &AdminagentworkplanbiddingpreferenceMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := AdminagentworkplanbiddingpreferenceMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Submitted, ok := AdminagentworkplanbiddingpreferenceMap["submitted"].(bool); ok {
		o.Submitted = &Submitted
	}
    
	if AssignedWorkPlan, ok := AdminagentworkplanbiddingpreferenceMap["assignedWorkPlan"].(map[string]interface{}); ok {
		AssignedWorkPlanString, _ := json.Marshal(AssignedWorkPlan)
		json.Unmarshal(AssignedWorkPlanString, &o.AssignedWorkPlan)
	}
	
	if OverriddenWorkPlan, ok := AdminagentworkplanbiddingpreferenceMap["overriddenWorkPlan"].(map[string]interface{}); ok {
		OverriddenWorkPlanString, _ := json.Marshal(OverriddenWorkPlan)
		json.Unmarshal(OverriddenWorkPlanString, &o.OverriddenWorkPlan)
	}
	
	if OverrideReason, ok := AdminagentworkplanbiddingpreferenceMap["overrideReason"].(string); ok {
		o.OverrideReason = &OverrideReason
	}
    
	if Priorities, ok := AdminagentworkplanbiddingpreferenceMap["priorities"].([]interface{}); ok {
		PrioritiesString, _ := json.Marshal(Priorities)
		json.Unmarshal(PrioritiesString, &o.Priorities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Adminagentworkplanbiddingpreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
