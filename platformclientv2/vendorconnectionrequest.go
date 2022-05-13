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

func (o *Vendorconnectionrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Vendorconnectionrequest
	
	return json.Marshal(&struct { 
		Publisher *string `json:"publisher,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Publisher: o.Publisher,
		
		VarType: o.VarType,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Vendorconnectionrequest) UnmarshalJSON(b []byte) error {
	var VendorconnectionrequestMap map[string]interface{}
	err := json.Unmarshal(b, &VendorconnectionrequestMap)
	if err != nil {
		return err
	}
	
	if Publisher, ok := VendorconnectionrequestMap["publisher"].(string); ok {
		o.Publisher = &Publisher
	}
    
	if VarType, ok := VendorconnectionrequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Name, ok := VendorconnectionrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Vendorconnectionrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
