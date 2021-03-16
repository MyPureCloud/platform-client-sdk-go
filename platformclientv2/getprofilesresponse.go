package platformclientv2
import (
	"encoding/json"
)

// Getprofilesresponse
type Getprofilesresponse struct { 
	// Total
	Total *int `json:"total,omitempty"`


	// Entities
	Entities *[]Performanceprofile `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Getprofilesresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
