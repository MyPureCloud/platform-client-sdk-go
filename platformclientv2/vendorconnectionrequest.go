package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Vendorconnectionrequest
type Vendorconnectionrequest struct { 
	// Publisher - Publisher of the integration or connector who registered the new connection. Typically, inin.
	Publisher *string `json:"publisher,omitempty"`


	// VarType - Integration or connector type that registered the new connection. Example, wfm-rta-integration
	VarType *string `json:"type,omitempty"`


	// Name - Name of the integration or connector instance that registered the new connection. Example, my-wfm
	Name *string `json:"name,omitempty"`

}

func (u *Vendorconnectionrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Vendorconnectionrequest

	

	return json.Marshal(&struct { 
		Publisher *string `json:"publisher,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Publisher: u.Publisher,
		
		VarType: u.VarType,
		
		Name: u.Name,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Vendorconnectionrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
