package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainpermission
type Domainpermission struct { 
	// Domain
	Domain *string `json:"domain,omitempty"`


	// EntityType
	EntityType *string `json:"entityType,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`


	// Label
	Label *string `json:"label,omitempty"`


	// AllowsConditions
	AllowsConditions *bool `json:"allowsConditions,omitempty"`


	// DivisionAware
	DivisionAware *bool `json:"divisionAware,omitempty"`

}

func (u *Domainpermission) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainpermission

	

	return json.Marshal(&struct { 
		Domain *string `json:"domain,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		Label *string `json:"label,omitempty"`
		
		AllowsConditions *bool `json:"allowsConditions,omitempty"`
		
		DivisionAware *bool `json:"divisionAware,omitempty"`
		*Alias
	}{ 
		Domain: u.Domain,
		
		EntityType: u.EntityType,
		
		Action: u.Action,
		
		Label: u.Label,
		
		AllowsConditions: u.AllowsConditions,
		
		DivisionAware: u.DivisionAware,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domainpermission) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
