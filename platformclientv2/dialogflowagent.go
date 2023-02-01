package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialogflowagent
type Dialogflowagent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Project - The project this Dialogflow agent belongs to
	Project *Dialogflowproject `json:"project,omitempty"`

	// Languages - The supported languages of the Dialogflow agent
	Languages *[]string `json:"languages,omitempty"`

	// Intents - An array of Intents associated with this agent
	Intents *[]Dialogflowintent `json:"intents,omitempty"`

	// Environments - Available environments for this agent
	Environments *[]string `json:"environments,omitempty"`

	// Integration - The Integration this Dialogflow agent was referenced from.
	Integration *Domainentityref `json:"integration,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialogflowagent) SetField(field string, fieldValue interface{}) {
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

func (o Dialogflowagent) MarshalJSON() ([]byte, error) {
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
	type Alias Dialogflowagent
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Project *Dialogflowproject `json:"project,omitempty"`
		
		Languages *[]string `json:"languages,omitempty"`
		
		Intents *[]Dialogflowintent `json:"intents,omitempty"`
		
		Environments *[]string `json:"environments,omitempty"`
		
		Integration *Domainentityref `json:"integration,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Project: o.Project,
		
		Languages: o.Languages,
		
		Intents: o.Intents,
		
		Environments: o.Environments,
		
		Integration: o.Integration,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Dialogflowagent) UnmarshalJSON(b []byte) error {
	var DialogflowagentMap map[string]interface{}
	err := json.Unmarshal(b, &DialogflowagentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialogflowagentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialogflowagentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Project, ok := DialogflowagentMap["project"].(map[string]interface{}); ok {
		ProjectString, _ := json.Marshal(Project)
		json.Unmarshal(ProjectString, &o.Project)
	}
	
	if Languages, ok := DialogflowagentMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if Intents, ok := DialogflowagentMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if Environments, ok := DialogflowagentMap["environments"].([]interface{}); ok {
		EnvironmentsString, _ := json.Marshal(Environments)
		json.Unmarshal(EnvironmentsString, &o.Environments)
	}
	
	if Integration, ok := DialogflowagentMap["integration"].(map[string]interface{}); ok {
		IntegrationString, _ := json.Marshal(Integration)
		json.Unmarshal(IntegrationString, &o.Integration)
	}
	
	if SelfUri, ok := DialogflowagentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dialogflowagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
