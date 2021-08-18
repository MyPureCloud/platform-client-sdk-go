package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Authzgrantrole
type Authzgrantrole struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Policies
	Policies *[]Authzgrantpolicy `json:"policies,omitempty"`


	// VarDefault
	VarDefault *bool `json:"default,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Authzgrantrole) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Authzgrantrole

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Policies *[]Authzgrantpolicy `json:"policies,omitempty"`
		
		VarDefault *bool `json:"default,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Policies: u.Policies,
		
		VarDefault: u.VarDefault,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Authzgrantrole) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
