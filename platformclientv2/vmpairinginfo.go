package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Vmpairinginfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Vmpairinginfo

	

	return json.Marshal(&struct { 
		MetaData *Metadata `json:"meta-data,omitempty"`
		
		EdgeId *string `json:"edge-id,omitempty"`
		
		AuthToken *string `json:"auth-token,omitempty"`
		
		OrgId *string `json:"org-id,omitempty"`
		*Alias
	}{ 
		MetaData: u.MetaData,
		
		EdgeId: u.EdgeId,
		
		AuthToken: u.AuthToken,
		
		OrgId: u.OrgId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Vmpairinginfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
