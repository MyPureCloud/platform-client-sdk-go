package platformclientv2
import (
	"encoding/json"
)

// Getmetricdefinitionsresponse
type Getmetricdefinitionsresponse struct { 
	// Total
	Total *int `json:"total,omitempty"`


	// Entities
	Entities *[]Metricdefinition `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Getmetricdefinitionsresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
