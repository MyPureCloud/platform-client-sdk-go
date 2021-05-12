package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Listwrappershiftstartvariance
type Listwrappershiftstartvariance struct { 
	// Values
	Values *[]Shiftstartvariance `json:"values,omitempty"`

}

// String returns a JSON representation of the model
func (o *Listwrappershiftstartvariance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
