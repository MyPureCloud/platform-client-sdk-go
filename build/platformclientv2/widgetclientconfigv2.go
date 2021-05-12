package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Widgetclientconfigv2
type Widgetclientconfigv2 struct { }

// String returns a JSON representation of the model
func (o *Widgetclientconfigv2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
