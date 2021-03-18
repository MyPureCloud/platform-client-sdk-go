package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Void
type Void struct { }

// String returns a JSON representation of the model
func (o *Void) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
