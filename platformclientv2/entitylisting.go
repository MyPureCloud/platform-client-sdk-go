package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Entitylisting
type Entitylisting struct { 
	// Entities
	Entities *[]interface{} `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Entitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
