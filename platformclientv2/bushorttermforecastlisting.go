package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bushorttermforecastlisting
type Bushorttermforecastlisting struct { 
	// Entities
	Entities *[]Bushorttermforecastlistitem `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bushorttermforecastlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
