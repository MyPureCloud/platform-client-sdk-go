package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Resterrordetail
type Resterrordetail struct { 
	// VarError - name of the error
	VarError *string `json:"error,omitempty"`


	// Details - additional information regarding the error
	Details *string `json:"details,omitempty"`

}

func (o *Resterrordetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Resterrordetail
	
	return json.Marshal(&struct { 
		VarError *string `json:"error,omitempty"`
		
		Details *string `json:"details,omitempty"`
		*Alias
	}{ 
		VarError: o.VarError,
		
		Details: o.Details,
		Alias:    (*Alias)(o),
	})
}

func (o *Resterrordetail) UnmarshalJSON(b []byte) error {
	var ResterrordetailMap map[string]interface{}
	err := json.Unmarshal(b, &ResterrordetailMap)
	if err != nil {
		return err
	}
	
	if VarError, ok := ResterrordetailMap["error"].(string); ok {
		o.VarError = &VarError
	}
	
	if Details, ok := ResterrordetailMap["details"].(string); ok {
		o.Details = &Details
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Resterrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
