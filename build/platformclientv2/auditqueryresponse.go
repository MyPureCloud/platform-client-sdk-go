package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryresponse
type Auditqueryresponse struct { }

// String returns a JSON representation of the model
func (o *Auditqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
