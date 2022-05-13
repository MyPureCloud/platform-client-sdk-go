package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeconnectioninfo
type Edgeconnectioninfo struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// InterfaceName - Interface used for the connection on the edge
	InterfaceName *string `json:"interfaceName,omitempty"`


	// InterfaceIpAddress - IP address of the interface
	InterfaceIpAddress *string `json:"interfaceIpAddress,omitempty"`


	// ConnectionErrors - Connection errors
	ConnectionErrors *[]string `json:"connectionErrors,omitempty"`


	// Site
	Site *Addressableentityref `json:"site,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Edgeconnectioninfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgeconnectioninfo
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		InterfaceName *string `json:"interfaceName,omitempty"`
		
		InterfaceIpAddress *string `json:"interfaceIpAddress,omitempty"`
		
		ConnectionErrors *[]string `json:"connectionErrors,omitempty"`
		
		Site *Addressableentityref `json:"site,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		InterfaceName: o.InterfaceName,
		
		InterfaceIpAddress: o.InterfaceIpAddress,
		
		ConnectionErrors: o.ConnectionErrors,
		
		Site: o.Site,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgeconnectioninfo) UnmarshalJSON(b []byte) error {
	var EdgeconnectioninfoMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeconnectioninfoMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgeconnectioninfoMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EdgeconnectioninfoMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if InterfaceName, ok := EdgeconnectioninfoMap["interfaceName"].(string); ok {
		o.InterfaceName = &InterfaceName
	}
    
	if InterfaceIpAddress, ok := EdgeconnectioninfoMap["interfaceIpAddress"].(string); ok {
		o.InterfaceIpAddress = &InterfaceIpAddress
	}
    
	if ConnectionErrors, ok := EdgeconnectioninfoMap["connectionErrors"].([]interface{}); ok {
		ConnectionErrorsString, _ := json.Marshal(ConnectionErrors)
		json.Unmarshal(ConnectionErrorsString, &o.ConnectionErrors)
	}
	
	if Site, ok := EdgeconnectioninfoMap["site"].(map[string]interface{}); ok {
		SiteString, _ := json.Marshal(Site)
		json.Unmarshal(SiteString, &o.Site)
	}
	
	if SelfUri, ok := EdgeconnectioninfoMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeconnectioninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
