package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricsnetworktypeip
type Trunkmetricsnetworktypeip struct { 
	// Address - Assigned IP Address for the interface
	Address *string `json:"address,omitempty"`


	// ErrorInfo - Information about the error.
	ErrorInfo *Trunkerrorinfo `json:"errorInfo,omitempty"`

}

func (o *Trunkmetricsnetworktypeip) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricsnetworktypeip
	
	return json.Marshal(&struct { 
		Address *string `json:"address,omitempty"`
		
		ErrorInfo *Trunkerrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		Address: o.Address,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkmetricsnetworktypeip) UnmarshalJSON(b []byte) error {
	var TrunkmetricsnetworktypeipMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkmetricsnetworktypeipMap)
	if err != nil {
		return err
	}
	
	if Address, ok := TrunkmetricsnetworktypeipMap["address"].(string); ok {
		o.Address = &Address
	}
	
	if ErrorInfo, ok := TrunkmetricsnetworktypeipMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkmetricsnetworktypeip) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
