package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Intentdefinition
type Intentdefinition struct { 
	// Name - The name of the intent.
	Name *string `json:"name,omitempty"`


	// EntityTypeBindings - The bindings for the named entity types used in this intent.This field is mutually exclusive with entityNameReferences and entities
	EntityTypeBindings *[]Namedentitytypebinding `json:"entityTypeBindings,omitempty"`


	// EntityNameReferences - The references for the named entity used in this intent.This field is mutually exclusive with entityTypeBindings
	EntityNameReferences *[]string `json:"entityNameReferences,omitempty"`


	// Utterances - The utterances that act as training phrases for the intent.
	Utterances *[]Nluutterance `json:"utterances,omitempty"`

}

func (u *Intentdefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Intentdefinition

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		EntityTypeBindings *[]Namedentitytypebinding `json:"entityTypeBindings,omitempty"`
		
		EntityNameReferences *[]string `json:"entityNameReferences,omitempty"`
		
		Utterances *[]Nluutterance `json:"utterances,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		EntityTypeBindings: u.EntityTypeBindings,
		
		EntityNameReferences: u.EntityNameReferences,
		
		Utterances: u.Utterances,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Intentdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
