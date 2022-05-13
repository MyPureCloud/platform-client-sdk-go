package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Resourcepermissionpolicy
type Resourcepermissionpolicy struct { 
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

func (o *Resourcepermissionpolicy) MarshalJSON() ([]byte, error) {
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
		*Alias
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
		Alias:    (*Alias)(o),
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
