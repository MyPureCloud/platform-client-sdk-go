package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalimportstatuslisting
type Historicalimportstatuslisting struct { 
	// Entities
	Entities *[]Historicalimportstatus `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Historicalimportstatuslisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
