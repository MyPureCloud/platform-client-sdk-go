package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Clientapp - Details for a ClientApp
type Clientapp struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the integration, used to distinguish this integration from others of the same type.
	Name *string `json:"name,omitempty"`

	// IntegrationType - Type of the integration
	IntegrationType *Integrationtype `json:"integrationType,omitempty"`

	// Notes - Notes about the integration.
	Notes *string `json:"notes,omitempty"`

	// IntendedState - Configured state of the integration.
	IntendedState *string `json:"intendedState,omitempty"`

	// Config - Configuration information for the integration.
	Config *Clientappconfigurationinfo `json:"config,omitempty"`

	// ReportedState - Last reported status of the integration.
	ReportedState *Integrationstatusinfo `json:"reportedState,omitempty"`

	// Attributes - Read-only attributes for the integration.
	Attributes *map[string]string `json:"attributes,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Clientapp) SetField(field string, fieldValue interface{}) {
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

func (o Clientapp) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Clientapp
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		IntegrationType *Integrationtype `json:"integrationType,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		IntendedState *string `json:"intendedState,omitempty"`
		
		Config *Clientappconfigurationinfo `json:"config,omitempty"`
		
		ReportedState *Integrationstatusinfo `json:"reportedState,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		IntegrationType: o.IntegrationType,
		
		Notes: o.Notes,
		
		IntendedState: o.IntendedState,
		
		Config: o.Config,
		
		ReportedState: o.ReportedState,
		
		Attributes: o.Attributes,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Clientapp) UnmarshalJSON(b []byte) error {
	var ClientappMap map[string]interface{}
	err := json.Unmarshal(b, &ClientappMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ClientappMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ClientappMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if IntegrationType, ok := ClientappMap["integrationType"].(map[string]interface{}); ok {
		IntegrationTypeString, _ := json.Marshal(IntegrationType)
		json.Unmarshal(IntegrationTypeString, &o.IntegrationType)
	}
	
	if Notes, ok := ClientappMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if IntendedState, ok := ClientappMap["intendedState"].(string); ok {
		o.IntendedState = &IntendedState
	}
    
	if Config, ok := ClientappMap["config"].(map[string]interface{}); ok {
		ConfigString, _ := json.Marshal(Config)
		json.Unmarshal(ConfigString, &o.Config)
	}
	
	if ReportedState, ok := ClientappMap["reportedState"].(map[string]interface{}); ok {
		ReportedStateString, _ := json.Marshal(ReportedState)
		json.Unmarshal(ReportedStateString, &o.ReportedState)
	}
	
	if Attributes, ok := ClientappMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if SelfUri, ok := ClientappMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Clientapp) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
