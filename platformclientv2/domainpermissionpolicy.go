package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainpermissionpolicy
type Domainpermissionpolicy struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Domain
	Domain *string `json:"domain,omitempty"`

	// EntityName
	EntityName *string `json:"entityName,omitempty"`

	// PolicyName
	PolicyName *string `json:"policyName,omitempty"`

	// PolicyDescription
	PolicyDescription *string `json:"policyDescription,omitempty"`

	// ActionSet
	ActionSet *[]string `json:"actionSet,omitempty"`

	// NamedResources
	NamedResources *[]string `json:"namedResources,omitempty"`

	// AllowConditions
	AllowConditions *bool `json:"allowConditions,omitempty"`

	// ResourceConditionNode
	ResourceConditionNode *Domainresourceconditionnode `json:"resourceConditionNode,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Domainpermissionpolicy) SetField(field string, fieldValue interface{}) {
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

func (o Domainpermissionpolicy) MarshalJSON() ([]byte, error) {
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
	type Alias Domainpermissionpolicy
	
	return json.Marshal(&struct { 
		Domain *string `json:"domain,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		PolicyName *string `json:"policyName,omitempty"`
		
		PolicyDescription *string `json:"policyDescription,omitempty"`
		
		ActionSet *[]string `json:"actionSet,omitempty"`
		
		NamedResources *[]string `json:"namedResources,omitempty"`
		
		AllowConditions *bool `json:"allowConditions,omitempty"`
		
		ResourceConditionNode *Domainresourceconditionnode `json:"resourceConditionNode,omitempty"`
		Alias
	}{ 
		Domain: o.Domain,
		
		EntityName: o.EntityName,
		
		PolicyName: o.PolicyName,
		
		PolicyDescription: o.PolicyDescription,
		
		ActionSet: o.ActionSet,
		
		NamedResources: o.NamedResources,
		
		AllowConditions: o.AllowConditions,
		
		ResourceConditionNode: o.ResourceConditionNode,
		Alias:    (Alias)(o),
	})
}

func (o *Domainpermissionpolicy) UnmarshalJSON(b []byte) error {
	var DomainpermissionpolicyMap map[string]interface{}
	err := json.Unmarshal(b, &DomainpermissionpolicyMap)
	if err != nil {
		return err
	}
	
	if Domain, ok := DomainpermissionpolicyMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if EntityName, ok := DomainpermissionpolicyMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
    
	if PolicyName, ok := DomainpermissionpolicyMap["policyName"].(string); ok {
		o.PolicyName = &PolicyName
	}
    
	if PolicyDescription, ok := DomainpermissionpolicyMap["policyDescription"].(string); ok {
		o.PolicyDescription = &PolicyDescription
	}
    
	if ActionSet, ok := DomainpermissionpolicyMap["actionSet"].([]interface{}); ok {
		ActionSetString, _ := json.Marshal(ActionSet)
		json.Unmarshal(ActionSetString, &o.ActionSet)
	}
	
	if NamedResources, ok := DomainpermissionpolicyMap["namedResources"].([]interface{}); ok {
		NamedResourcesString, _ := json.Marshal(NamedResources)
		json.Unmarshal(NamedResourcesString, &o.NamedResources)
	}
	
	if AllowConditions, ok := DomainpermissionpolicyMap["allowConditions"].(bool); ok {
		o.AllowConditions = &AllowConditions
	}
    
	if ResourceConditionNode, ok := DomainpermissionpolicyMap["resourceConditionNode"].(map[string]interface{}); ok {
		ResourceConditionNodeString, _ := json.Marshal(ResourceConditionNode)
		json.Unmarshal(ResourceConditionNodeString, &o.ResourceConditionNode)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainpermissionpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
