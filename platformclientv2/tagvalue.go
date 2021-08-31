package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Tagvalue
type Tagvalue struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The workspace tag name.
	Name *string `json:"name,omitempty"`


	// InUse
	InUse *bool `json:"inUse,omitempty"`


	// Acl
	Acl *[]string `json:"acl,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Tagvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Tagvalue
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		InUse *bool `json:"inUse,omitempty"`
		
		Acl *[]string `json:"acl,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		InUse: o.InUse,
		
		Acl: o.Acl,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Tagvalue) UnmarshalJSON(b []byte) error {
	var TagvalueMap map[string]interface{}
	err := json.Unmarshal(b, &TagvalueMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TagvalueMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := TagvalueMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if InUse, ok := TagvalueMap["inUse"].(bool); ok {
		o.InUse = &InUse
	}
	
	if Acl, ok := TagvalueMap["acl"].([]interface{}); ok {
		AclString, _ := json.Marshal(Acl)
		json.Unmarshal(AclString, &o.Acl)
	}
	
	if SelfUri, ok := TagvalueMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Tagvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
