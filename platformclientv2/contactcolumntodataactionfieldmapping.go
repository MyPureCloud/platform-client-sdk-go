package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contactcolumntodataactionfieldmapping
type Contactcolumntodataactionfieldmapping struct { }

// String returns a JSON representation of the model
func (o *Contactcolumntodataactionfieldmapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
