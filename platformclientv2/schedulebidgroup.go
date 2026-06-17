package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulebidgroup
type Schedulebidgroup struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the schedule bid group
	Name *string `json:"name,omitempty"`

	// ManagementUnit - The management unit to which this bid group belongs
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`

	// Agents - The agents who participate in this bid group
	Agents *[]Userreference `json:"agents,omitempty"`

	// WorkPlans - The work plans used in this bid group
	WorkPlans *[]Workplanreference `json:"workPlans,omitempty"`

	// WorkPlanRotations - The work plan rotations used in this bid group
	WorkPlanRotations *[]Bidgroupworkplanrotationresponse `json:"workPlanRotations,omitempty"`

	// PlanningGroups - The planning groups selected in this bid group
	PlanningGroups *[]Planninggroupreference `json:"planningGroups,omitempty"`

	// DownloadUrl - The downloadUrl to fetch Schedule sets. It will be populated if the status of this bid is 'Optimized'
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// DownloadTemplate - Schedule sets always come through downloadUrl, the schema included here is just for documentation
	DownloadTemplate *Bidgroupscheduleset `json:"downloadTemplate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Schedulebidgroup) SetField(field string, fieldValue interface{}) {
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

func (o Schedulebidgroup) MarshalJSON() ([]byte, error) {
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
	type Alias Schedulebidgroup
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		Agents *[]Userreference `json:"agents,omitempty"`
		
		WorkPlans *[]Workplanreference `json:"workPlans,omitempty"`
		
		WorkPlanRotations *[]Bidgroupworkplanrotationresponse `json:"workPlanRotations,omitempty"`
		
		PlanningGroups *[]Planninggroupreference `json:"planningGroups,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		DownloadTemplate *Bidgroupscheduleset `json:"downloadTemplate,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		ManagementUnit: o.ManagementUnit,
		
		Agents: o.Agents,
		
		WorkPlans: o.WorkPlans,
		
		WorkPlanRotations: o.WorkPlanRotations,
		
		PlanningGroups: o.PlanningGroups,
		
		DownloadUrl: o.DownloadUrl,
		
		DownloadTemplate: o.DownloadTemplate,
		Alias:    (Alias)(o),
	})
}

func (o *Schedulebidgroup) UnmarshalJSON(b []byte) error {
	var SchedulebidgroupMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulebidgroupMap)
	if err != nil {
		return err
	}
	
	if Name, ok := SchedulebidgroupMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ManagementUnit, ok := SchedulebidgroupMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if Agents, ok := SchedulebidgroupMap["agents"].([]interface{}); ok {
		AgentsString, _ := json.Marshal(Agents)
		json.Unmarshal(AgentsString, &o.Agents)
	}
	
	if WorkPlans, ok := SchedulebidgroupMap["workPlans"].([]interface{}); ok {
		WorkPlansString, _ := json.Marshal(WorkPlans)
		json.Unmarshal(WorkPlansString, &o.WorkPlans)
	}
	
	if WorkPlanRotations, ok := SchedulebidgroupMap["workPlanRotations"].([]interface{}); ok {
		WorkPlanRotationsString, _ := json.Marshal(WorkPlanRotations)
		json.Unmarshal(WorkPlanRotationsString, &o.WorkPlanRotations)
	}
	
	if PlanningGroups, ok := SchedulebidgroupMap["planningGroups"].([]interface{}); ok {
		PlanningGroupsString, _ := json.Marshal(PlanningGroups)
		json.Unmarshal(PlanningGroupsString, &o.PlanningGroups)
	}
	
	if DownloadUrl, ok := SchedulebidgroupMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if DownloadTemplate, ok := SchedulebidgroupMap["downloadTemplate"].(map[string]interface{}); ok {
		DownloadTemplateString, _ := json.Marshal(DownloadTemplate)
		json.Unmarshal(DownloadTemplateString, &o.DownloadTemplate)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulebidgroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
