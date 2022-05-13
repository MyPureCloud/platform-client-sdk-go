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

func (o *Domainpermission) MarshalJSON() ([]byte, error) {
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
		Domain: o.Domain,
		
		EntityType: o.EntityType,
		
		Action: o.Action,
		
		Label: o.Label,
		
		AllowsConditions: o.AllowsConditions,
		
		DivisionAware: o.DivisionAware,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainpermission) UnmarshalJSON(b []byte) error {
	var DomainpermissionMap map[string]interface{}
	err := json.Unmarshal(b, &DomainpermissionMap)
	if err != nil {
		return err
	}
	
	if Domain, ok := DomainpermissionMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if EntityType, ok := DomainpermissionMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
    
	if Action, ok := DomainpermissionMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if Label, ok := DomainpermissionMap["label"].(string); ok {
		o.Label = &Label
	}
    
	if AllowsConditions, ok := DomainpermissionMap["allowsConditions"].(bool); ok {
		o.AllowsConditions = &AllowsConditions
	}
    
	if DivisionAware, ok := DomainpermissionMap["divisionAware"].(bool); ok {
		o.DivisionAware = &DivisionAware
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Domainpermission) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
