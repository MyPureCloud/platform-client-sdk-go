package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Queryrequestpredicate
type Queryrequestpredicate struct { 
	// Dimension - The dimension to be filtered
	Dimension *string `json:"dimension,omitempty"`


	// Value - The value to filter by
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queryrequestpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
