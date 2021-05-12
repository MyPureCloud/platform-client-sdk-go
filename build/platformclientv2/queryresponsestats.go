package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queryresponsestats
type Queryresponsestats struct { 
	// Count - The count for this metric
	Count *int `json:"count,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queryresponsestats) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
