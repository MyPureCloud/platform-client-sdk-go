package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Domainpermissionpolicy) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
