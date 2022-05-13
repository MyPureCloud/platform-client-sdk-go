package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Fieldconfig
type Fieldconfig struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// EntityType
	EntityType *string `json:"entityType,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Sections
	Sections *[]Section `json:"sections,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// SchemaVersion
	SchemaVersion *string `json:"schemaVersion,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Fieldconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Fieldconfig
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Sections *[]Section `json:"sections,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		SchemaVersion *string `json:"schemaVersion,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		EntityType: o.EntityType,
		
		State: o.State,
		
		Sections: o.Sections,
		
		Version: o.Version,
		
		SchemaVersion: o.SchemaVersion,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Fieldconfig) UnmarshalJSON(b []byte) error {
	var FieldconfigMap map[string]interface{}
	err := json.Unmarshal(b, &FieldconfigMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FieldconfigMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FieldconfigMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if EntityType, ok := FieldconfigMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
    
	if State, ok := FieldconfigMap["state"].(string); ok {
		o.State = &State
	}
    
	if Sections, ok := FieldconfigMap["sections"].([]interface{}); ok {
		SectionsString, _ := json.Marshal(Sections)
		json.Unmarshal(SectionsString, &o.Sections)
	}
	
	if Version, ok := FieldconfigMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if SchemaVersion, ok := FieldconfigMap["schemaVersion"].(string); ok {
		o.SchemaVersion = &SchemaVersion
	}
    
	if SelfUri, ok := FieldconfigMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Fieldconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
