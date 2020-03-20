package platformclientv2
import (
	"encoding/json"
)

// Queryfacetinfo
type Queryfacetinfo struct { 
	// Attributes
	Attributes *[]Facetkeyattribute `json:"attributes,omitempty"`


	// Facets
	Facets *[]Facetentry `json:"facets,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queryfacetinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
