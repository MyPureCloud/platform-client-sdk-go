package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Authzgrantrole
type Authzgrantrole struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Policies
	Policies *[]Authzgrantpolicy `json:"policies,omitempty"`


	// VarDefault
	VarDefault *bool `json:"default,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Authzgrantrole) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Authzgrantrole
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Policies *[]Authzgrantpolicy `json:"policies,omitempty"`
		
		VarDefault *bool `json:"default,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Policies: o.Policies,
		
		VarDefault: o.VarDefault,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Authzgrantrole) UnmarshalJSON(b []byte) error {
	var AuthzgrantroleMap map[string]interface{}
	err := json.Unmarshal(b, &AuthzgrantroleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AuthzgrantroleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AuthzgrantroleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := AuthzgrantroleMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Policies, ok := AuthzgrantroleMap["policies"].([]interface{}); ok {
		PoliciesString, _ := json.Marshal(Policies)
		json.Unmarshal(PoliciesString, &o.Policies)
	}
	
	if VarDefault, ok := AuthzgrantroleMap["default"].(bool); ok {
		o.VarDefault = &VarDefault
	}
    
	if SelfUri, ok := AuthzgrantroleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Authzgrantrole) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
