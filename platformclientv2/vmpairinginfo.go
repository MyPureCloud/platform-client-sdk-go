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

func (o *Vmpairinginfo) MarshalJSON() ([]byte, error) {
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
		MetaData: o.MetaData,
		
		EdgeId: o.EdgeId,
		
		AuthToken: o.AuthToken,
		
		OrgId: o.OrgId,
		Alias:    (*Alias)(o),
	})
}

func (o *Vmpairinginfo) UnmarshalJSON(b []byte) error {
	var VmpairinginfoMap map[string]interface{}
	err := json.Unmarshal(b, &VmpairinginfoMap)
	if err != nil {
		return err
	}
	
	if MetaData, ok := VmpairinginfoMap["meta-data"].(map[string]interface{}); ok {
		MetaDataString, _ := json.Marshal(MetaData)
		json.Unmarshal(MetaDataString, &o.MetaData)
	}
	
	if EdgeId, ok := VmpairinginfoMap["edge-id"].(string); ok {
		o.EdgeId = &EdgeId
	}
	
	if AuthToken, ok := VmpairinginfoMap["auth-token"].(string); ok {
		o.AuthToken = &AuthToken
	}
	
	if OrgId, ok := VmpairinginfoMap["org-id"].(string); ok {
		o.OrgId = &OrgId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Vmpairinginfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
