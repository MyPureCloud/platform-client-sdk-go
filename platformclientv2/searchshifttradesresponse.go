package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Searchshifttradesresponse
type Searchshifttradesresponse struct { 
	// Trades - The shift trades that match the search criteria
	Trades *[]Searchshifttraderesponse `json:"trades,omitempty"`

}

// String returns a JSON representation of the model
func (o *Searchshifttradesresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
