package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricstopictrunkmetricsqos
type Trunkmetricstopictrunkmetricsqos struct { 
	// MismatchCount
	MismatchCount *int `json:"mismatchCount,omitempty"`

}

func (u *Trunkmetricstopictrunkmetricsqos) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricstopictrunkmetricsqos

	

	return json.Marshal(&struct { 
		MismatchCount *int `json:"mismatchCount,omitempty"`
		*Alias
	}{ 
		MismatchCount: u.MismatchCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkmetricstopictrunkmetricsqos) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
