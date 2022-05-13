package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainpermissioncollection
type Domainpermissioncollection struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Domain
	Domain *string `json:"domain,omitempty"`


	// PermissionMap
	PermissionMap *map[string][]Domainpermission `json:"permissionMap,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Domainpermissioncollection) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainpermissioncollection
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		PermissionMap *map[string][]Domainpermission `json:"permissionMap,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Domain: o.Domain,
		
		PermissionMap: o.PermissionMap,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainpermissioncollection) UnmarshalJSON(b []byte) error {
	var DomainpermissioncollectionMap map[string]interface{}
	err := json.Unmarshal(b, &DomainpermissioncollectionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DomainpermissioncollectionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DomainpermissioncollectionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Domain, ok := DomainpermissioncollectionMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if PermissionMap, ok := DomainpermissioncollectionMap["permissionMap"].(map[string]interface{}); ok {
		PermissionMapString, _ := json.Marshal(PermissionMap)
		json.Unmarshal(PermissionMapString, &o.PermissionMap)
	}
	
	if SelfUri, ok := DomainpermissioncollectionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Domainpermissioncollection) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
