package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationqualityv2topiccalibration
type Evaluationqualityv2topiccalibration struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationqualityv2topiccalibration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
