package platformclientv2
import (
	"encoding/json"
)

// Vmpairinginfo
type Vmpairinginfo struct { 
	// MetaData - This is to be used to complete the setup process of a locally deployed virtual edge device.
	MetaData *Metadata `json:"meta-data,omitempty"`


	// EdgeId
	EdgeId *string `json:"edge-id,omitempty"`


	// AuthToken
	AuthToken *string `json:"auth-token,omitempty"`


	// OrgId
	OrgId *string `json:"org-id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Vmpairinginfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
