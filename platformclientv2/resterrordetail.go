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

func (u *Resterrordetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Resterrordetail

	

	return json.Marshal(&struct { 
		VarError *string `json:"error,omitempty"`
		
		Details *string `json:"details,omitempty"`
		*Alias
	}{ 
		VarError: u.VarError,
		
		Details: u.Details,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Resterrordetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
