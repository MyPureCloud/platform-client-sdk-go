package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Cloneduser - Represents a cloned user in a trustor organization.
type Cloneduser struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Trustor - The ID of the trustor organization this clone exists in.
	Trustor *Domainentityref `json:"trustor,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Cloneduser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Cloneduser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Trustor *Domainentityref `json:"trustor,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Trustor: o.Trustor,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Cloneduser) UnmarshalJSON(b []byte) error {
	var CloneduserMap map[string]interface{}
	err := json.Unmarshal(b, &CloneduserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CloneduserMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := CloneduserMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Trustor, ok := CloneduserMap["trustor"].(map[string]interface{}); ok {
		TrustorString, _ := json.Marshal(Trustor)
		json.Unmarshal(TrustorString, &o.Trustor)
	}
	
	if SelfUri, ok := CloneduserMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Cloneduser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
