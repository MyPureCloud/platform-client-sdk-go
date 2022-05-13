package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Connectededge
type Connectededge struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// InterfaceName - Edge interface name used for the connection
	InterfaceName *string `json:"interfaceName,omitempty"`


	// InterfaceIpAddress - Edge interface IP address
	InterfaceIpAddress *string `json:"interfaceIpAddress,omitempty"`


	// EdgeConnectionList
	EdgeConnectionList *[]Edgeconnectioninfo `json:"edgeConnectionList,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Connectededge) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Connectededge
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		InterfaceName *string `json:"interfaceName,omitempty"`
		
		InterfaceIpAddress *string `json:"interfaceIpAddress,omitempty"`
		
		EdgeConnectionList *[]Edgeconnectioninfo `json:"edgeConnectionList,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		InterfaceName: o.InterfaceName,
		
		InterfaceIpAddress: o.InterfaceIpAddress,
		
		EdgeConnectionList: o.EdgeConnectionList,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Connectededge) UnmarshalJSON(b []byte) error {
	var ConnectededgeMap map[string]interface{}
	err := json.Unmarshal(b, &ConnectededgeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConnectededgeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ConnectededgeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if InterfaceName, ok := ConnectededgeMap["interfaceName"].(string); ok {
		o.InterfaceName = &InterfaceName
	}
    
	if InterfaceIpAddress, ok := ConnectededgeMap["interfaceIpAddress"].(string); ok {
		o.InterfaceIpAddress = &InterfaceIpAddress
	}
    
	if EdgeConnectionList, ok := ConnectededgeMap["edgeConnectionList"].([]interface{}); ok {
		EdgeConnectionListString, _ := json.Marshal(EdgeConnectionList)
		json.Unmarshal(EdgeConnectionListString, &o.EdgeConnectionList)
	}
	
	if SelfUri, ok := ConnectededgeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Connectededge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
