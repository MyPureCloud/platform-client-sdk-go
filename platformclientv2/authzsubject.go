package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Authzsubject
type Authzsubject struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Grants
	Grants *[]Authzgrant `json:"grants,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Authzsubject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Authzsubject
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Grants *[]Authzgrant `json:"grants,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Grants: o.Grants,
		
		Version: o.Version,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Authzsubject) UnmarshalJSON(b []byte) error {
	var AuthzsubjectMap map[string]interface{}
	err := json.Unmarshal(b, &AuthzsubjectMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AuthzsubjectMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := AuthzsubjectMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Grants, ok := AuthzsubjectMap["grants"].([]interface{}); ok {
		GrantsString, _ := json.Marshal(Grants)
		json.Unmarshal(GrantsString, &o.Grants)
	}
	
	if Version, ok := AuthzsubjectMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if SelfUri, ok := AuthzsubjectMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Authzsubject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
