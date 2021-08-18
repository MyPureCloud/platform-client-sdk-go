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

func (u *Resourcepermissionpolicy) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Domain: u.Domain,
		
		EntityName: u.EntityName,
		
		PolicyName: u.PolicyName,
		
		PolicyDescription: u.PolicyDescription,
		
		ActionSetKey: u.ActionSetKey,
		
		AllowConditions: u.AllowConditions,
		
		ResourceConditionNode: u.ResourceConditionNode,
		
		NamedResources: u.NamedResources,
		
		ResourceCondition: u.ResourceCondition,
		
		ActionSet: u.ActionSet,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Resourcepermissionpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
