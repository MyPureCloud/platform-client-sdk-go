package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Domainpermission) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
