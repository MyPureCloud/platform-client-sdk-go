package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Streetaddress
type Streetaddress struct { 
	// Country - 2 Letter Country code, like US or GB
	Country *string `json:"country,omitempty"`


	// A1 - State or Province
	A1 *string `json:"A1,omitempty"`


	// A3 - City or township
	A3 *string `json:"A3,omitempty"`


	// RD - Number and street
	RD *string `json:"RD,omitempty"`


	// HNO - House Number
	HNO *string `json:"HNO,omitempty"`


	// LOC - extra location info like suite 300
	LOC *string `json:"LOC,omitempty"`


	// NAM - Name of the customer
	NAM *string `json:"NAM,omitempty"`


	// PC - Postal code
	PC *string `json:"PC,omitempty"`

}

// String returns a JSON representation of the model
func (o *Streetaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
