package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Trunkmetricstopictrunkmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
