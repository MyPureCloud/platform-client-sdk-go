package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Action
type Action struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// IntegrationId - The ID of the integration for which this action is associated
	IntegrationId *string `json:"integrationId,omitempty"`

	// Category - Category of Action
	Category *string `json:"category,omitempty"`

	// Contract - Action contract
	Contract *Actioncontract `json:"contract,omitempty"`

	// Version - Version of this action
	Version *int `json:"version,omitempty"`

	// Secure - Indication of whether or not the action is designed to accept sensitive data
	Secure *bool `json:"secure,omitempty"`

	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Action) SetField(field string, fieldValue interface{}) {
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

func (o Action) MarshalJSON() ([]byte, error) {
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
	type Alias Action
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		Contract *Actioncontract `json:"contract,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Secure *bool `json:"secure,omitempty"`
		
		Config *Actionconfig `json:"config,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		IntegrationId: o.IntegrationId,
		
		Category: o.Category,
		
		Contract: o.Contract,
		
		Version: o.Version,
		
		Secure: o.Secure,
		
		Config: o.Config,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Action) UnmarshalJSON(b []byte) error {
	var ActionMap map[string]interface{}
	err := json.Unmarshal(b, &ActionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ActionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if IntegrationId, ok := ActionMap["integrationId"].(string); ok {
		o.IntegrationId = &IntegrationId
	}
    
	if Category, ok := ActionMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if Contract, ok := ActionMap["contract"].(map[string]interface{}); ok {
		ContractString, _ := json.Marshal(Contract)
		json.Unmarshal(ContractString, &o.Contract)
	}
	
	if Version, ok := ActionMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Secure, ok := ActionMap["secure"].(bool); ok {
		o.Secure = &Secure
	}
    
	if Config, ok := ActionMap["config"].(map[string]interface{}); ok {
		ConfigString, _ := json.Marshal(Config)
		json.Unmarshal(ConfigString, &o.Config)
	}
	
	if SelfUri, ok := ActionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Action) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
