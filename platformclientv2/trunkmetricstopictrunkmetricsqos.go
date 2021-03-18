package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricstopictrunkmetricsqos
type Trunkmetricstopictrunkmetricsqos struct { 
	// MismatchCount
	MismatchCount *int `json:"mismatchCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricstopictrunkmetricsqos) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
