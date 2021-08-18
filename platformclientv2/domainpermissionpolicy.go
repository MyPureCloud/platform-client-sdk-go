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

func (u *Domainpermissionpolicy) MarshalJSON() ([]byte, error) {
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
		Domain: u.Domain,
		
		EntityName: u.EntityName,
		
		PolicyName: u.PolicyName,
		
		PolicyDescription: u.PolicyDescription,
		
		ActionSet: u.ActionSet,
		
		NamedResources: u.NamedResources,
		
		AllowConditions: u.AllowConditions,
		
		ResourceConditionNode: u.ResourceConditionNode,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domainpermissionpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
