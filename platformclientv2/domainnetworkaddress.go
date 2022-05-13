package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainnetworkaddress
type Domainnetworkaddress struct { 
	// VarType - The type of address.
	VarType *string `json:"type,omitempty"`


	// Address - An IPv4 or IPv6 IP address. When specifying an address of type \"ip\", use CIDR format for the subnet mask.
	Address *string `json:"address,omitempty"`


	// Persistent - True if this address will persist on Edge restart.  Addresses assigned by DHCP will be returned as false.
	Persistent *bool `json:"persistent,omitempty"`


	// Family - The address family for this address.
	Family *int `json:"family,omitempty"`

}

func (o *Domainnetworkaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainnetworkaddress
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		Persistent *bool `json:"persistent,omitempty"`
		
		Family *int `json:"family,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Address: o.Address,
		
		Persistent: o.Persistent,
		
		Family: o.Family,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainnetworkaddress) UnmarshalJSON(b []byte) error {
	var DomainnetworkaddressMap map[string]interface{}
	err := json.Unmarshal(b, &DomainnetworkaddressMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DomainnetworkaddressMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Address, ok := DomainnetworkaddressMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if Persistent, ok := DomainnetworkaddressMap["persistent"].(bool); ok {
		o.Persistent = &Persistent
	}
    
	if Family, ok := DomainnetworkaddressMap["family"].(float64); ok {
		FamilyInt := int(Family)
		o.Family = &FamilyInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainnetworkaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
