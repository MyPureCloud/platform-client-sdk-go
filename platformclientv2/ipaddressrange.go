package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ipaddressrange
type Ipaddressrange struct { 
	// Cidr
	Cidr *string `json:"cidr,omitempty"`


	// Service
	Service *string `json:"service,omitempty"`


	// Region
	Region *string `json:"region,omitempty"`

}

func (o *Ipaddressrange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ipaddressrange
	
	return json.Marshal(&struct { 
		Cidr *string `json:"cidr,omitempty"`
		
		Service *string `json:"service,omitempty"`
		
		Region *string `json:"region,omitempty"`
		*Alias
	}{ 
		Cidr: o.Cidr,
		
		Service: o.Service,
		
		Region: o.Region,
		Alias:    (*Alias)(o),
	})
}

func (o *Ipaddressrange) UnmarshalJSON(b []byte) error {
	var IpaddressrangeMap map[string]interface{}
	err := json.Unmarshal(b, &IpaddressrangeMap)
	if err != nil {
		return err
	}
	
	if Cidr, ok := IpaddressrangeMap["cidr"].(string); ok {
		o.Cidr = &Cidr
	}
    
	if Service, ok := IpaddressrangeMap["service"].(string); ok {
		o.Service = &Service
	}
    
	if Region, ok := IpaddressrangeMap["region"].(string); ok {
		o.Region = &Region
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Ipaddressrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
