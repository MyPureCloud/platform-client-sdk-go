package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagent
type Wfmagent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// User - The user associated with this data
	User *Userreference `json:"user,omitempty"`

	// WorkPlan - The work plan associated with this agent, if applicable
	WorkPlan *Workplanreference `json:"workPlan,omitempty"`

	// WorkPlanRotation - The work plan rotation associated with this agent, if applicable
	WorkPlanRotation *Workplanrotationreference `json:"workPlanRotation,omitempty"`

	// AcceptDirectShiftTrades - Whether the agent accepts direct shift trade requests
	AcceptDirectShiftTrades *bool `json:"acceptDirectShiftTrades,omitempty"`

	// WorkPlanOverrides - The work plan overrides associated with this agent. Populate with expand=workPlanOverrides
	WorkPlanOverrides *[]Workplanoverride `json:"workPlanOverrides,omitempty"`

	// Queues - List of queues to which this agent is capable of handling
	Queues *[]Queuereference `json:"queues,omitempty"`

	// Languages - The list of languages this agent is capable of handling
	Languages *[]Languagereference `json:"languages,omitempty"`

	// Skills - The list of skills this agent is capable of handling
	Skills *[]Routingskillreference `json:"skills,omitempty"`

	// Schedulable - Whether the agent can be included in schedule generation
	Schedulable *bool `json:"schedulable,omitempty"`

	// Metadata - Metadata for this agent
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmagent) SetField(field string, fieldValue interface{}) {
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

func (o Wfmagent) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmagent
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		WorkPlan *Workplanreference `json:"workPlan,omitempty"`
		
		WorkPlanRotation *Workplanrotationreference `json:"workPlanRotation,omitempty"`
		
		AcceptDirectShiftTrades *bool `json:"acceptDirectShiftTrades,omitempty"`
		
		WorkPlanOverrides *[]Workplanoverride `json:"workPlanOverrides,omitempty"`
		
		Queues *[]Queuereference `json:"queues,omitempty"`
		
		Languages *[]Languagereference `json:"languages,omitempty"`
		
		Skills *[]Routingskillreference `json:"skills,omitempty"`
		
		Schedulable *bool `json:"schedulable,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		User: o.User,
		
		WorkPlan: o.WorkPlan,
		
		WorkPlanRotation: o.WorkPlanRotation,
		
		AcceptDirectShiftTrades: o.AcceptDirectShiftTrades,
		
		WorkPlanOverrides: o.WorkPlanOverrides,
		
		Queues: o.Queues,
		
		Languages: o.Languages,
		
		Skills: o.Skills,
		
		Schedulable: o.Schedulable,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmagent) UnmarshalJSON(b []byte) error {
	var WfmagentMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmagentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if User, ok := WfmagentMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if WorkPlan, ok := WfmagentMap["workPlan"].(map[string]interface{}); ok {
		WorkPlanString, _ := json.Marshal(WorkPlan)
		json.Unmarshal(WorkPlanString, &o.WorkPlan)
	}
	
	if WorkPlanRotation, ok := WfmagentMap["workPlanRotation"].(map[string]interface{}); ok {
		WorkPlanRotationString, _ := json.Marshal(WorkPlanRotation)
		json.Unmarshal(WorkPlanRotationString, &o.WorkPlanRotation)
	}
	
	if AcceptDirectShiftTrades, ok := WfmagentMap["acceptDirectShiftTrades"].(bool); ok {
		o.AcceptDirectShiftTrades = &AcceptDirectShiftTrades
	}
    
	if WorkPlanOverrides, ok := WfmagentMap["workPlanOverrides"].([]interface{}); ok {
		WorkPlanOverridesString, _ := json.Marshal(WorkPlanOverrides)
		json.Unmarshal(WorkPlanOverridesString, &o.WorkPlanOverrides)
	}
	
	if Queues, ok := WfmagentMap["queues"].([]interface{}); ok {
		QueuesString, _ := json.Marshal(Queues)
		json.Unmarshal(QueuesString, &o.Queues)
	}
	
	if Languages, ok := WfmagentMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if Skills, ok := WfmagentMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if Schedulable, ok := WfmagentMap["schedulable"].(bool); ok {
		o.Schedulable = &Schedulable
	}
    
	if Metadata, ok := WfmagentMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := WfmagentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
