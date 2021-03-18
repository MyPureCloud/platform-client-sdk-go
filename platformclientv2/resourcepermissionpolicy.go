package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Resourcepermissionpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
