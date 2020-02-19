package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Vendorconnectionrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
