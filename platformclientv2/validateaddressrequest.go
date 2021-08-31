package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Validateaddressrequest
type Validateaddressrequest struct { 
	// Address - Address schema
	Address *Streetaddress `json:"address,omitempty"`

}

func (o *Validateaddressrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validateaddressrequest
	
	return json.Marshal(&struct { 
		Address *Streetaddress `json:"address,omitempty"`
		*Alias
	}{ 
		Address: o.Address,
		Alias:    (*Alias)(o),
	})
}

func (o *Validateaddressrequest) UnmarshalJSON(b []byte) error {
	var ValidateaddressrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ValidateaddressrequestMap)
	if err != nil {
		return err
	}
	
	if Address, ok := ValidateaddressrequestMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Validateaddressrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
