package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dependencyobject
type Dependencyobject struct { 
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


	// ConsumedResources
	ConsumedResources *[]Dependency `json:"consumedResources,omitempty"`


	// ConsumingResources
	ConsumingResources *[]Dependency `json:"consumingResources,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Dependencyobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dependencyobject
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		Updated *bool `json:"updated,omitempty"`
		
		StateUnknown *bool `json:"stateUnknown,omitempty"`
		
		ConsumedResources *[]Dependency `json:"consumedResources,omitempty"`
		
		ConsumingResources *[]Dependency `json:"consumingResources,omitempty"`
		
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
		
		ConsumedResources: o.ConsumedResources,
		
		ConsumingResources: o.ConsumingResources,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Dependencyobject) UnmarshalJSON(b []byte) error {
	var DependencyobjectMap map[string]interface{}
	err := json.Unmarshal(b, &DependencyobjectMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DependencyobjectMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DependencyobjectMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Version, ok := DependencyobjectMap["version"].(string); ok {
		o.Version = &Version
	}
	
	if VarType, ok := DependencyobjectMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Deleted, ok := DependencyobjectMap["deleted"].(bool); ok {
		o.Deleted = &Deleted
	}
	
	if Updated, ok := DependencyobjectMap["updated"].(bool); ok {
		o.Updated = &Updated
	}
	
	if StateUnknown, ok := DependencyobjectMap["stateUnknown"].(bool); ok {
		o.StateUnknown = &StateUnknown
	}
	
	if ConsumedResources, ok := DependencyobjectMap["consumedResources"].([]interface{}); ok {
		ConsumedResourcesString, _ := json.Marshal(ConsumedResources)
		json.Unmarshal(ConsumedResourcesString, &o.ConsumedResources)
	}
	
	if ConsumingResources, ok := DependencyobjectMap["consumingResources"].([]interface{}); ok {
		ConsumingResourcesString, _ := json.Marshal(ConsumingResources)
		json.Unmarshal(ConsumingResourcesString, &o.ConsumingResources)
	}
	
	if SelfUri, ok := DependencyobjectMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dependencyobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
