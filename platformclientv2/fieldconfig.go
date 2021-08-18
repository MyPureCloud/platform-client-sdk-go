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

func (u *Fieldconfig) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		EntityType: u.EntityType,
		
		State: u.State,
		
		Sections: u.Sections,
		
		Version: u.Version,
		
		SchemaVersion: u.SchemaVersion,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Fieldconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
