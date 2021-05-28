package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Connectededge) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
