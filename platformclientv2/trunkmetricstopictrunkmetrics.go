package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricstopictrunkmetrics
type Trunkmetricstopictrunkmetrics struct { 
	// Calls
	Calls *Trunkmetricstopictrunkmetricscalls `json:"calls,omitempty"`


	// Qos
	Qos *Trunkmetricstopictrunkmetricsqos `json:"qos,omitempty"`


	// Trunk
	Trunk *Trunkmetricstopicurireference `json:"trunk,omitempty"`

}

func (u *Trunkmetricstopictrunkmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricstopictrunkmetrics

	

	return json.Marshal(&struct { 
		Calls *Trunkmetricstopictrunkmetricscalls `json:"calls,omitempty"`
		
		Qos *Trunkmetricstopictrunkmetricsqos `json:"qos,omitempty"`
		
		Trunk *Trunkmetricstopicurireference `json:"trunk,omitempty"`
		*Alias
	}{ 
		Calls: u.Calls,
		
		Qos: u.Qos,
		
		Trunk: u.Trunk,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkmetricstopictrunkmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
