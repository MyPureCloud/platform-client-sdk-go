package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Licensedefinition
type Licensedefinition struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Permissions
	Permissions *Permissions `json:"permissions,omitempty"`


	// Prerequisites
	Prerequisites *[]Addressablelicensedefinition `json:"prerequisites,omitempty"`


	// Comprises
	Comprises *[]Licensedefinition `json:"comprises,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Licensedefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Licensedefinition

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Permissions *Permissions `json:"permissions,omitempty"`
		
		Prerequisites *[]Addressablelicensedefinition `json:"prerequisites,omitempty"`
		
		Comprises *[]Licensedefinition `json:"comprises,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Description: u.Description,
		
		Permissions: u.Permissions,
		
		Prerequisites: u.Prerequisites,
		
		Comprises: u.Comprises,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Licensedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
