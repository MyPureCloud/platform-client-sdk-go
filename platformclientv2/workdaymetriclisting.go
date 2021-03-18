package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Workdaymetriclisting
type Workdaymetriclisting struct { 
	// Entities
	Entities *[]Workdaymetric `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workdaymetriclisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
