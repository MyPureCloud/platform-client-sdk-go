package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Generaltopicsentitylisting
type Generaltopicsentitylisting struct { 
	// Entities
	Entities *[]Generaltopic `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Generaltopicsentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
