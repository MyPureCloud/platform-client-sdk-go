package platformclientv2
import (
	"encoding/json"
)

// Namedentitytypeitem
type Namedentitytypeitem struct { 
	// Value - A value for an named entity type definition.
	Value *string `json:"value,omitempty"`


	// Synonyms - Synonyms for the given named entity value.
	Synonyms *[]string `json:"synonyms,omitempty"`

}

// String returns a JSON representation of the model
func (o *Namedentitytypeitem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
