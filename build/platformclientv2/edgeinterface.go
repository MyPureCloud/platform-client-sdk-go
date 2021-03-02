package platformclientv2
import (
	"encoding/json"
)

// Edgeinterface
type Edgeinterface struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// MacAddress
	MacAddress *string `json:"macAddress,omitempty"`


	// IfName
	IfName *string `json:"ifName,omitempty"`


	// Endpoints
	Endpoints *[]Domainentityref `json:"endpoints,omitempty"`


	// LineTypes
	LineTypes *[]string `json:"lineTypes,omitempty"`


	// AddressFamilyId
	AddressFamilyId *string `json:"addressFamilyId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgeinterface) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
