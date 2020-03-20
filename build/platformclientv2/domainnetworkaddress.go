package platformclientv2
import (
	"encoding/json"
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
	Family *int32 `json:"family,omitempty"`

}

// String returns a JSON representation of the model
func (o *Domainnetworkaddress) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
