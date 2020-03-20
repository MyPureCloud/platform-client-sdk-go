package platformclientv2
import (
	"encoding/json"
)

// Localencryptionconfigurationlisting
type Localencryptionconfigurationlisting struct { 
	// Total
	Total *int64 `json:"total,omitempty"`


	// Entities
	Entities *[]Localencryptionconfiguration `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Localencryptionconfigurationlisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
