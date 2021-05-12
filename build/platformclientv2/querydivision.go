package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Querydivision
type Querydivision struct { }

// String returns a JSON representation of the model
func (o *Querydivision) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
