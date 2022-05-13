package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dependency
type Dependency struct { 
	// Id - The dependency identifier
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Deleted
	Deleted *bool `json:"deleted,omitempty"`


	// Updated
	Updated *bool `json:"updated,omitempty"`


	// StateUnknown
	StateUnknown *bool `json:"stateUnknown,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Dependency) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dependency
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		Updated *bool `json:"updated,omitempty"`
		
		StateUnknown *bool `json:"stateUnknown,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Version: o.Version,
		
		VarType: o.VarType,
		
		Deleted: o.Deleted,
		
		Updated: o.Updated,
		
		StateUnknown: o.StateUnknown,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Dependency) UnmarshalJSON(b []byte) error {
	var DependencyMap map[string]interface{}
	err := json.Unmarshal(b, &DependencyMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DependencyMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DependencyMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Version, ok := DependencyMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if VarType, ok := DependencyMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Deleted, ok := DependencyMap["deleted"].(bool); ok {
		o.Deleted = &Deleted
	}
    
	if Updated, ok := DependencyMap["updated"].(bool); ok {
		o.Updated = &Updated
	}
    
	if StateUnknown, ok := DependencyMap["stateUnknown"].(bool); ok {
		o.StateUnknown = &StateUnknown
	}
    
	if SelfUri, ok := DependencyMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dependency) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
