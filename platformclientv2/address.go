package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Address) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Address
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		NameRaw *string `json:"nameRaw,omitempty"`
		
		AddressNormalized *string `json:"addressNormalized,omitempty"`
		
		AddressRaw *string `json:"addressRaw,omitempty"`
		
		AddressDisplayable *string `json:"addressDisplayable,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		NameRaw: o.NameRaw,
		
		AddressNormalized: o.AddressNormalized,
		
		AddressRaw: o.AddressRaw,
		
		AddressDisplayable: o.AddressDisplayable,
		Alias:    (*Alias)(o),
	})
}

func (o *Address) UnmarshalJSON(b []byte) error {
	var AddressMap map[string]interface{}
	err := json.Unmarshal(b, &AddressMap)
	if err != nil {
		return err
	}
	
	if Name, ok := AddressMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if NameRaw, ok := AddressMap["nameRaw"].(string); ok {
		o.NameRaw = &NameRaw
	}
	
	if AddressNormalized, ok := AddressMap["addressNormalized"].(string); ok {
		o.AddressNormalized = &AddressNormalized
	}
	
	if AddressRaw, ok := AddressMap["addressRaw"].(string); ok {
		o.AddressRaw = &AddressRaw
	}
	
	if AddressDisplayable, ok := AddressMap["addressDisplayable"].(string); ok {
		o.AddressDisplayable = &AddressDisplayable
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Address) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
