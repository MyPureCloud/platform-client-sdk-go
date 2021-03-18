package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Ignoredactivitycategories
type Ignoredactivitycategories struct { 
	// Values - Activity categories list
	Values *[]string `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ignoredactivitycategories) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
