package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticssessionmetric
type Analyticssessionmetric struct { 
	// EmitDate - Metric emission date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EmitDate *time.Time `json:"emitDate,omitempty"`


	// Name - Unique name of this metric
	Name *string `json:"name,omitempty"`


	// Value - The metric value
	Value *int `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticssessionmetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
