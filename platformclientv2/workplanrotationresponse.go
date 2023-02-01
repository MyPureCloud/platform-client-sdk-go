package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanrotationresponse
type Workplanrotationresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Enabled - Whether the work plan rotation is enabled for scheduling
	Enabled *bool `json:"enabled,omitempty"`

	// DateRange - The date range to which this work plan rotation applies
	DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`

	// Pattern - Pattern with ordered list of work plans that rotate on a weekly basis
	Pattern *Workplanpatternresponse `json:"pattern,omitempty"`

	// AgentCount - Number of agents in this work plan rotation
	AgentCount *int `json:"agentCount,omitempty"`

	// Agents - Agents in this work plan rotation. Populate with expand=agents for GET WorkPlanRotationsList (defaults to empty list)
	Agents *[]Workplanrotationagentresponse `json:"agents,omitempty"`

	// Metadata - Version metadata for this work plan rotation
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workplanrotationresponse) SetField(field string, fieldValue interface{}) {
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

func (o Workplanrotationresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Workplanrotationresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		DateRange *Daterangewithoptionalend `json:"dateRange,omitempty"`
		
		Pattern *Workplanpatternresponse `json:"pattern,omitempty"`
		
		AgentCount *int `json:"agentCount,omitempty"`
		
		Agents *[]Workplanrotationagentresponse `json:"agents,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Enabled: o.Enabled,
		
		DateRange: o.DateRange,
		
		Pattern: o.Pattern,
		
		AgentCount: o.AgentCount,
		
		Agents: o.Agents,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Workplanrotationresponse) UnmarshalJSON(b []byte) error {
	var WorkplanrotationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanrotationresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkplanrotationresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorkplanrotationresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Enabled, ok := WorkplanrotationresponseMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if DateRange, ok := WorkplanrotationresponseMap["dateRange"].(map[string]interface{}); ok {
		DateRangeString, _ := json.Marshal(DateRange)
		json.Unmarshal(DateRangeString, &o.DateRange)
	}
	
	if Pattern, ok := WorkplanrotationresponseMap["pattern"].(map[string]interface{}); ok {
		PatternString, _ := json.Marshal(Pattern)
		json.Unmarshal(PatternString, &o.Pattern)
	}
	
	if AgentCount, ok := WorkplanrotationresponseMap["agentCount"].(float64); ok {
		AgentCountInt := int(AgentCount)
		o.AgentCount = &AgentCountInt
	}
	
	if Agents, ok := WorkplanrotationresponseMap["agents"].([]interface{}); ok {
		AgentsString, _ := json.Marshal(Agents)
		json.Unmarshal(AgentsString, &o.Agents)
	}
	
	if Metadata, ok := WorkplanrotationresponseMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := WorkplanrotationresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanrotationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
