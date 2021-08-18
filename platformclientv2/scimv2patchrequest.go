package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2patchrequest - Defines a SCIM PATCH request. See section 3.5.2 \"Modifying with PATCH\" in RFC 7644 for details.
type Scimv2patchrequest struct { 
	// Schemas - The list of schemas used in the PATCH request.
	Schemas *[]string `json:"schemas,omitempty"`


	// Operations - The list of operations to perform for the PATCH request.
	Operations *[]Scimv2patchoperation `json:"Operations,omitempty"`

}

func (u *Scimv2patchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimv2patchrequest

	

	return json.Marshal(&struct { 
		Schemas *[]string `json:"schemas,omitempty"`
		
		Operations *[]Scimv2patchoperation `json:"Operations,omitempty"`
		*Alias
	}{ 
		Schemas: u.Schemas,
		
		Operations: u.Operations,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Scimv2patchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
