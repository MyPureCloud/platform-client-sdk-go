package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Edgeconnectioninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
