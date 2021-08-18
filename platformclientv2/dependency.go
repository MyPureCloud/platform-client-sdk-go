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

func (u *Dependency) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		Version: u.Version,
		
		VarType: u.VarType,
		
		Deleted: u.Deleted,
		
		Updated: u.Updated,
		
		StateUnknown: u.StateUnknown,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dependency) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
