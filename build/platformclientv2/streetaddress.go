package platformclientv2
import (
	"encoding/json"
)

// Streetaddress
type Streetaddress struct { 
	// Country - 2 Letter Country code, like US or GB
	Country *string `json:"country,omitempty"`


	// A1 - State or Province
	A1 *string `json:"A1,omitempty"`


	// A3 - City or township
	A3 *string `json:"A3,omitempty"`


	// RD
	RD *string `json:"RD,omitempty"`


	// HNO
	HNO *string `json:"HNO,omitempty"`


	// LOC
	LOC *string `json:"LOC,omitempty"`


	// NAM
	NAM *string `json:"NAM,omitempty"`


	// PC
	PC *string `json:"PC,omitempty"`

}

// String returns a JSON representation of the model
func (o *Streetaddress) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
