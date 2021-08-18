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

func (u *Edgeconnectioninfo) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		InterfaceName: u.InterfaceName,
		
		InterfaceIpAddress: u.InterfaceIpAddress,
		
		ConnectionErrors: u.ConnectionErrors,
		
		Site: u.Site,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgeconnectioninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
