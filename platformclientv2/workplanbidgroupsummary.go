package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanbidgroupsummary
type Workplanbidgroupsummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name assigned to this bid group
	Name *string `json:"name,omitempty"`

	// ManagementUnit - The management unit this bid group belongs to
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`

	// AgentCount - The number of agents in this bid group
	AgentCount *int `json:"agentCount,omitempty"`

	// WorkPlanCount - The number of work plans in this bid group
	WorkPlanCount *int `json:"workPlanCount,omitempty"`

	// PlanningGroupCount - The number of planning groups in this bid group
	PlanningGroupCount *int `json:"planningGroupCount,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workplanbidgroupsummary) SetField(field string, fieldValue interface{}) {
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

func (o Workplanbidgroupsummary) MarshalJSON() ([]byte, error) {
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
	type Alias Workplanbidgroupsummary
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		AgentCount *int `json:"agentCount,omitempty"`
		
		WorkPlanCount *int `json:"workPlanCount,omitempty"`
		
		PlanningGroupCount *int `json:"planningGroupCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ManagementUnit: o.ManagementUnit,
		
		AgentCount: o.AgentCount,
		
		WorkPlanCount: o.WorkPlanCount,
		
		PlanningGroupCount: o.PlanningGroupCount,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Workplanbidgroupsummary) UnmarshalJSON(b []byte) error {
	var WorkplanbidgroupsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanbidgroupsummaryMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkplanbidgroupsummaryMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorkplanbidgroupsummaryMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ManagementUnit, ok := WorkplanbidgroupsummaryMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if AgentCount, ok := WorkplanbidgroupsummaryMap["agentCount"].(float64); ok {
		AgentCountInt := int(AgentCount)
		o.AgentCount = &AgentCountInt
	}
	
	if WorkPlanCount, ok := WorkplanbidgroupsummaryMap["workPlanCount"].(float64); ok {
		WorkPlanCountInt := int(WorkPlanCount)
		o.WorkPlanCount = &WorkPlanCountInt
	}
	
	if PlanningGroupCount, ok := WorkplanbidgroupsummaryMap["planningGroupCount"].(float64); ok {
		PlanningGroupCountInt := int(PlanningGroupCount)
		o.PlanningGroupCount = &PlanningGroupCountInt
	}
	
	if SelfUri, ok := WorkplanbidgroupsummaryMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanbidgroupsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
