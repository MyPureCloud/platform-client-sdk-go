package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentintegrationsresponse
type Agentintegrationsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Agent - The user associated with the integrations
	Agent *Userreference `json:"agent,omitempty"`

	// SelectedIntegration - The integration selected for the agent. If not set, no integration will be used for the agent
	SelectedIntegration *Wfmintegrationreference `json:"selectedIntegration,omitempty"`

	// UserSelected - Whether the integration association has been manually selected
	UserSelected *bool `json:"userSelected,omitempty"`

	// AssociatedIntegrations - The list of integrations associated with the agent
	AssociatedIntegrations *[]Agentintegrationassociationresponse `json:"associatedIntegrations,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentintegrationsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Agentintegrationsresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Agentintegrationsresponse
	
	return json.Marshal(&struct { 
		Agent *Userreference `json:"agent,omitempty"`
		
		SelectedIntegration *Wfmintegrationreference `json:"selectedIntegration,omitempty"`
		
		UserSelected *bool `json:"userSelected,omitempty"`
		
		AssociatedIntegrations *[]Agentintegrationassociationresponse `json:"associatedIntegrations,omitempty"`
		Alias
	}{ 
		Agent: o.Agent,
		
		SelectedIntegration: o.SelectedIntegration,
		
		UserSelected: o.UserSelected,
		
		AssociatedIntegrations: o.AssociatedIntegrations,
		Alias:    (Alias)(o),
	})
}

func (o *Agentintegrationsresponse) UnmarshalJSON(b []byte) error {
	var AgentintegrationsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AgentintegrationsresponseMap)
	if err != nil {
		return err
	}
	
	if Agent, ok := AgentintegrationsresponseMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if SelectedIntegration, ok := AgentintegrationsresponseMap["selectedIntegration"].(map[string]interface{}); ok {
		SelectedIntegrationString, _ := json.Marshal(SelectedIntegration)
		json.Unmarshal(SelectedIntegrationString, &o.SelectedIntegration)
	}
	
	if UserSelected, ok := AgentintegrationsresponseMap["userSelected"].(bool); ok {
		o.UserSelected = &UserSelected
	}
    
	if AssociatedIntegrations, ok := AgentintegrationsresponseMap["associatedIntegrations"].([]interface{}); ok {
		AssociatedIntegrationsString, _ := json.Marshal(AssociatedIntegrations)
		json.Unmarshal(AssociatedIntegrationsString, &o.AssociatedIntegrations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentintegrationsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
