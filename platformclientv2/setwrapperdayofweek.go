package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Setwrapperdayofweek
type Setwrapperdayofweek struct { 
	// Values
	Values *[]string `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Setwrapperdayofweek) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
