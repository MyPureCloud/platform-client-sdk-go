package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentitytypeitem
type Namedentitytypeitem struct { 
	// Value - A value for an named entity type definition.
	Value *string `json:"value,omitempty"`


	// Synonyms - Synonyms for the given named entity value.
	Synonyms *[]string `json:"synonyms,omitempty"`

}

func (u *Namedentitytypeitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Namedentitytypeitem

	

	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		
		Synonyms *[]string `json:"synonyms,omitempty"`
		*Alias
	}{ 
		Value: u.Value,
		
		Synonyms: u.Synonyms,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Namedentitytypeitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
