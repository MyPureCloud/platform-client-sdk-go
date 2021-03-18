package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkupdateshifttradestateresult
type Bulkupdateshifttradestateresult struct { 
	// Entities
	Entities *[]Bulkupdateshifttradestateresultitem `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateshifttradestateresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
