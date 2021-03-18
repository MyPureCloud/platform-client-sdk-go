package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Address
type Address struct { 
	// Name - This will be nameRaw if present, or a locality lookup of the address field otherwise.
	Name *string `json:"name,omitempty"`


	// NameRaw - The name as close to the bits on the wire as possible.
	NameRaw *string `json:"nameRaw,omitempty"`


	// AddressNormalized - The normalized address. This field is acquired from the Address Normalization Table.  The addressRaw could have gone through some transformations, such as only using the numeric portion, before being run through the Address Normalization Table.
	AddressNormalized *string `json:"addressNormalized,omitempty"`


	// AddressRaw - The address as close to the bits on the wire as possible.
	AddressRaw *string `json:"addressRaw,omitempty"`


	// AddressDisplayable - The displayable address. This field is acquired from the Address Normalization Table.  The addressRaw could have gone through some transformations, such as only using the numeric portion, before being run through the Address Normalization Table.
	AddressDisplayable *string `json:"addressDisplayable,omitempty"`

}

// String returns a JSON representation of the model
func (o *Address) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
