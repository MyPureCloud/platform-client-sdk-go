package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanbidgroupcreate
type Workplanbidgroupcreate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the work plan bid group
	Name *string `json:"name,omitempty"`

	// ManagementUnitId - The management unit ID this bid group belongs to
	ManagementUnitId *string `json:"managementUnitId,omitempty"`

	// AgentIds - Agent IDs who participate in this bid group
	AgentIds *[]string `json:"agentIds,omitempty"`

	// WorkPlans - The list of work plans used in this bid group
	WorkPlans *[]Bidgroupworkplanrequest `json:"workPlans,omitempty"`

	// PlanningGroupIds - The planning group IDs selected in this bid group
	PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workplanbidgroupcreate) SetField(field string, fieldValue interface{}) {
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

func (o Workplanbidgroupcreate) MarshalJSON() ([]byte, error) {
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
	type Alias Workplanbidgroupcreate
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ManagementUnitId *string `json:"managementUnitId,omitempty"`
		
		AgentIds *[]string `json:"agentIds,omitempty"`
		
		WorkPlans *[]Bidgroupworkplanrequest `json:"workPlans,omitempty"`
		
		PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		ManagementUnitId: o.ManagementUnitId,
		
		AgentIds: o.AgentIds,
		
		WorkPlans: o.WorkPlans,
		
		PlanningGroupIds: o.PlanningGroupIds,
		Alias:    (Alias)(o),
	})
}

func (o *Workplanbidgroupcreate) UnmarshalJSON(b []byte) error {
	var WorkplanbidgroupcreateMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanbidgroupcreateMap)
	if err != nil {
		return err
	}
	
	if Name, ok := WorkplanbidgroupcreateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ManagementUnitId, ok := WorkplanbidgroupcreateMap["managementUnitId"].(string); ok {
		o.ManagementUnitId = &ManagementUnitId
	}
    
	if AgentIds, ok := WorkplanbidgroupcreateMap["agentIds"].([]interface{}); ok {
		AgentIdsString, _ := json.Marshal(AgentIds)
		json.Unmarshal(AgentIdsString, &o.AgentIds)
	}
	
	if WorkPlans, ok := WorkplanbidgroupcreateMap["workPlans"].([]interface{}); ok {
		WorkPlansString, _ := json.Marshal(WorkPlans)
		json.Unmarshal(WorkPlansString, &o.WorkPlans)
	}
	
	if PlanningGroupIds, ok := WorkplanbidgroupcreateMap["planningGroupIds"].([]interface{}); ok {
		PlanningGroupIdsString, _ := json.Marshal(PlanningGroupIds)
		json.Unmarshal(PlanningGroupIdsString, &o.PlanningGroupIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanbidgroupcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
