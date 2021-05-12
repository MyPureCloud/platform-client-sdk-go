package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Domainpermissioncollection) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
