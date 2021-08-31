package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Securityprofile
type Securityprofile struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Permissions
	Permissions *[]string `json:"permissions,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Securityprofile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Securityprofile
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Permissions *[]string `json:"permissions,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Permissions: o.Permissions,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Securityprofile) UnmarshalJSON(b []byte) error {
	var SecurityprofileMap map[string]interface{}
	err := json.Unmarshal(b, &SecurityprofileMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SecurityprofileMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := SecurityprofileMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Permissions, ok := SecurityprofileMap["permissions"].([]interface{}); ok {
		PermissionsString, _ := json.Marshal(Permissions)
		json.Unmarshal(PermissionsString, &o.Permissions)
	}
	
	if SelfUri, ok := SecurityprofileMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Securityprofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
