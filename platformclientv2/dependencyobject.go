package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dependencyobject
type Dependencyobject struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The dependency identifier
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Version
	Version *string `json:"version,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// Deleted
	Deleted *bool `json:"deleted,omitempty"`

	// Updated
	Updated *bool `json:"updated,omitempty"`

	// StateUnknown
	StateUnknown *bool `json:"stateUnknown,omitempty"`

	// ConsumedResources
	ConsumedResources *[]Dependency `json:"consumedResources,omitempty"`

	// ConsumingResources
	ConsumingResources *[]Dependency `json:"consumingResources,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dependencyobject) SetField(field string, fieldValue interface{}) {
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

func (o Dependencyobject) MarshalJSON() ([]byte, error) {
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
	type Alias Dependencyobject
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		Updated *bool `json:"updated,omitempty"`
		
		StateUnknown *bool `json:"stateUnknown,omitempty"`
		
		ConsumedResources *[]Dependency `json:"consumedResources,omitempty"`
		
		ConsumingResources *[]Dependency `json:"consumingResources,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Version: o.Version,
		
		VarType: o.VarType,
		
		Deleted: o.Deleted,
		
		Updated: o.Updated,
		
		StateUnknown: o.StateUnknown,
		
		ConsumedResources: o.ConsumedResources,
		
		ConsumingResources: o.ConsumingResources,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Dependencyobject) UnmarshalJSON(b []byte) error {
	var DependencyobjectMap map[string]interface{}
	err := json.Unmarshal(b, &DependencyobjectMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DependencyobjectMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DependencyobjectMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Version, ok := DependencyobjectMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if VarType, ok := DependencyobjectMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Deleted, ok := DependencyobjectMap["deleted"].(bool); ok {
		o.Deleted = &Deleted
	}
    
	if Updated, ok := DependencyobjectMap["updated"].(bool); ok {
		o.Updated = &Updated
	}
    
	if StateUnknown, ok := DependencyobjectMap["stateUnknown"].(bool); ok {
		o.StateUnknown = &StateUnknown
	}
    
	if ConsumedResources, ok := DependencyobjectMap["consumedResources"].([]interface{}); ok {
		ConsumedResourcesString, _ := json.Marshal(ConsumedResources)
		json.Unmarshal(ConsumedResourcesString, &o.ConsumedResources)
	}
	
	if ConsumingResources, ok := DependencyobjectMap["consumingResources"].([]interface{}); ok {
		ConsumingResourcesString, _ := json.Marshal(ConsumingResources)
		json.Unmarshal(ConsumingResourcesString, &o.ConsumingResources)
	}
	
	if SelfUri, ok := DependencyobjectMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dependencyobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
