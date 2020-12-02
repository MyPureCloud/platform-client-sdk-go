package platformclientv2
import (
	"encoding/json"
)

// Oauthscopelisting
type Oauthscopelisting struct { 
	// Total
	Total *int64 `json:"total,omitempty"`


	// Entities
	Entities *[]Oauthscope `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Oauthscopelisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
