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

func (o *Edgeinterface) MarshalJSON() ([]byte, error) {
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
		VarType: o.VarType,
		
		IpAddress: o.IpAddress,
		
		Name: o.Name,
		
		MacAddress: o.MacAddress,
		
		IfName: o.IfName,
		
		Endpoints: o.Endpoints,
		
		LineTypes: o.LineTypes,
		
		AddressFamilyId: o.AddressFamilyId,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgeinterface) UnmarshalJSON(b []byte) error {
	var EdgeinterfaceMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeinterfaceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := EdgeinterfaceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if IpAddress, ok := EdgeinterfaceMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if Name, ok := EdgeinterfaceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if MacAddress, ok := EdgeinterfaceMap["macAddress"].(string); ok {
		o.MacAddress = &MacAddress
	}
    
	if IfName, ok := EdgeinterfaceMap["ifName"].(string); ok {
		o.IfName = &IfName
	}
    
	if Endpoints, ok := EdgeinterfaceMap["endpoints"].([]interface{}); ok {
		EndpointsString, _ := json.Marshal(Endpoints)
		json.Unmarshal(EndpointsString, &o.Endpoints)
	}
	
	if LineTypes, ok := EdgeinterfaceMap["lineTypes"].([]interface{}); ok {
		LineTypesString, _ := json.Marshal(LineTypes)
		json.Unmarshal(LineTypesString, &o.LineTypes)
	}
	
	if AddressFamilyId, ok := EdgeinterfaceMap["addressFamilyId"].(string); ok {
		o.AddressFamilyId = &AddressFamilyId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeinterface) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
