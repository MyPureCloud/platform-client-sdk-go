package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulebidgroupcreate
type Schedulebidgroupcreate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the schedule bid group
	Name *string `json:"name,omitempty"`

	// ManagementUnitId - The ID of the management unit to which this bid group belongs
	ManagementUnitId *string `json:"managementUnitId,omitempty"`

	// AgentIds - The IDs of the agents who participate in this bid group
	AgentIds *[]string `json:"agentIds,omitempty"`

	// WorkPlanIds - The IDs of the work plans used in this bid group
	WorkPlanIds *[]string `json:"workPlanIds,omitempty"`

	// WorkPlanRotations - The work plan rotations used in this bid group
	WorkPlanRotations *[]Bidgroupworkplanrotationrequest `json:"workPlanRotations,omitempty"`

	// PlanningGroupIds - The IDs of the planning groups selected in this bid group
	PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Schedulebidgroupcreate) SetField(field string, fieldValue interface{}) {
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

func (o Schedulebidgroupcreate) MarshalJSON() ([]byte, error) {
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
	type Alias Schedulebidgroupcreate
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ManagementUnitId *string `json:"managementUnitId,omitempty"`
		
		AgentIds *[]string `json:"agentIds,omitempty"`
		
		WorkPlanIds *[]string `json:"workPlanIds,omitempty"`
		
		WorkPlanRotations *[]Bidgroupworkplanrotationrequest `json:"workPlanRotations,omitempty"`
		
		PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		ManagementUnitId: o.ManagementUnitId,
		
		AgentIds: o.AgentIds,
		
		WorkPlanIds: o.WorkPlanIds,
		
		WorkPlanRotations: o.WorkPlanRotations,
		
		PlanningGroupIds: o.PlanningGroupIds,
		Alias:    (Alias)(o),
	})
}

func (o *Schedulebidgroupcreate) UnmarshalJSON(b []byte) error {
	var SchedulebidgroupcreateMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulebidgroupcreateMap)
	if err != nil {
		return err
	}
	
	if Name, ok := SchedulebidgroupcreateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ManagementUnitId, ok := SchedulebidgroupcreateMap["managementUnitId"].(string); ok {
		o.ManagementUnitId = &ManagementUnitId
	}
    
	if AgentIds, ok := SchedulebidgroupcreateMap["agentIds"].([]interface{}); ok {
		AgentIdsString, _ := json.Marshal(AgentIds)
		json.Unmarshal(AgentIdsString, &o.AgentIds)
	}
	
	if WorkPlanIds, ok := SchedulebidgroupcreateMap["workPlanIds"].([]interface{}); ok {
		WorkPlanIdsString, _ := json.Marshal(WorkPlanIds)
		json.Unmarshal(WorkPlanIdsString, &o.WorkPlanIds)
	}
	
	if WorkPlanRotations, ok := SchedulebidgroupcreateMap["workPlanRotations"].([]interface{}); ok {
		WorkPlanRotationsString, _ := json.Marshal(WorkPlanRotations)
		json.Unmarshal(WorkPlanRotationsString, &o.WorkPlanRotations)
	}
	
	if PlanningGroupIds, ok := SchedulebidgroupcreateMap["planningGroupIds"].([]interface{}); ok {
		PlanningGroupIdsString, _ := json.Marshal(PlanningGroupIds)
		json.Unmarshal(PlanningGroupIdsString, &o.PlanningGroupIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulebidgroupcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
