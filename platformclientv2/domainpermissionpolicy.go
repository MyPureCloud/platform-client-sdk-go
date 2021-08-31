package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainpermissionpolicy
type Domainpermissionpolicy struct { 
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

func (o *Domainpermissionpolicy) MarshalJSON() ([]byte, error) {
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
		*Alias
	}{ 
		Domain: o.Domain,
		
		EntityName: o.EntityName,
		
		PolicyName: o.PolicyName,
		
		PolicyDescription: o.PolicyDescription,
		
		ActionSet: o.ActionSet,
		
		NamedResources: o.NamedResources,
		
		AllowConditions: o.AllowConditions,
		
		ResourceConditionNode: o.ResourceConditionNode,
		Alias:    (*Alias)(o),
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
