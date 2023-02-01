package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Resourcepermissionpolicy
type Resourcepermissionpolicy struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Domain
	Domain *string `json:"domain,omitempty"`

	// EntityName
	EntityName *string `json:"entityName,omitempty"`

	// PolicyName
	PolicyName *string `json:"policyName,omitempty"`

	// PolicyDescription
	PolicyDescription *string `json:"policyDescription,omitempty"`

	// ActionSetKey
	ActionSetKey *string `json:"actionSetKey,omitempty"`

	// AllowConditions
	AllowConditions *bool `json:"allowConditions,omitempty"`

	// ResourceConditionNode
	ResourceConditionNode *Resourceconditionnode `json:"resourceConditionNode,omitempty"`

	// NamedResources
	NamedResources *[]string `json:"namedResources,omitempty"`

	// ResourceCondition
	ResourceCondition *string `json:"resourceCondition,omitempty"`

	// ActionSet
	ActionSet *[]string `json:"actionSet,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Resourcepermissionpolicy) SetField(field string, fieldValue interface{}) {
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

func (o Resourcepermissionpolicy) MarshalJSON() ([]byte, error) {
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
	type Alias Resourcepermissionpolicy
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		
		PolicyName *string `json:"policyName,omitempty"`
		
		PolicyDescription *string `json:"policyDescription,omitempty"`
		
		ActionSetKey *string `json:"actionSetKey,omitempty"`
		
		AllowConditions *bool `json:"allowConditions,omitempty"`
		
		ResourceConditionNode *Resourceconditionnode `json:"resourceConditionNode,omitempty"`
		
		NamedResources *[]string `json:"namedResources,omitempty"`
		
		ResourceCondition *string `json:"resourceCondition,omitempty"`
		
		ActionSet *[]string `json:"actionSet,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Domain: o.Domain,
		
		EntityName: o.EntityName,
		
		PolicyName: o.PolicyName,
		
		PolicyDescription: o.PolicyDescription,
		
		ActionSetKey: o.ActionSetKey,
		
		AllowConditions: o.AllowConditions,
		
		ResourceConditionNode: o.ResourceConditionNode,
		
		NamedResources: o.NamedResources,
		
		ResourceCondition: o.ResourceCondition,
		
		ActionSet: o.ActionSet,
		Alias:    (Alias)(o),
	})
}

func (o *Resourcepermissionpolicy) UnmarshalJSON(b []byte) error {
	var ResourcepermissionpolicyMap map[string]interface{}
	err := json.Unmarshal(b, &ResourcepermissionpolicyMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ResourcepermissionpolicyMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Domain, ok := ResourcepermissionpolicyMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if EntityName, ok := ResourcepermissionpolicyMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
    
	if PolicyName, ok := ResourcepermissionpolicyMap["policyName"].(string); ok {
		o.PolicyName = &PolicyName
	}
    
	if PolicyDescription, ok := ResourcepermissionpolicyMap["policyDescription"].(string); ok {
		o.PolicyDescription = &PolicyDescription
	}
    
	if ActionSetKey, ok := ResourcepermissionpolicyMap["actionSetKey"].(string); ok {
		o.ActionSetKey = &ActionSetKey
	}
    
	if AllowConditions, ok := ResourcepermissionpolicyMap["allowConditions"].(bool); ok {
		o.AllowConditions = &AllowConditions
	}
    
	if ResourceConditionNode, ok := ResourcepermissionpolicyMap["resourceConditionNode"].(map[string]interface{}); ok {
		ResourceConditionNodeString, _ := json.Marshal(ResourceConditionNode)
		json.Unmarshal(ResourceConditionNodeString, &o.ResourceConditionNode)
	}
	
	if NamedResources, ok := ResourcepermissionpolicyMap["namedResources"].([]interface{}); ok {
		NamedResourcesString, _ := json.Marshal(NamedResources)
		json.Unmarshal(NamedResourcesString, &o.NamedResources)
	}
	
	if ResourceCondition, ok := ResourcepermissionpolicyMap["resourceCondition"].(string); ok {
		o.ResourceCondition = &ResourceCondition
	}
    
	if ActionSet, ok := ResourcepermissionpolicyMap["actionSet"].([]interface{}); ok {
		ActionSetString, _ := json.Marshal(ActionSet)
		json.Unmarshal(ActionSetString, &o.ActionSet)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Resourcepermissionpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
