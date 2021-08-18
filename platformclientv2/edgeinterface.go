package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Edgeinterface) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgeinterface

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		MacAddress *string `json:"macAddress,omitempty"`
		
		IfName *string `json:"ifName,omitempty"`
		
		Endpoints *[]Domainentityref `json:"endpoints,omitempty"`
		
		LineTypes *[]string `json:"lineTypes,omitempty"`
		
		AddressFamilyId *string `json:"addressFamilyId,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		IpAddress: u.IpAddress,
		
		Name: u.Name,
		
		MacAddress: u.MacAddress,
		
		IfName: u.IfName,
		
		Endpoints: u.Endpoints,
		
		LineTypes: u.LineTypes,
		
		AddressFamilyId: u.AddressFamilyId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgeinterface) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
