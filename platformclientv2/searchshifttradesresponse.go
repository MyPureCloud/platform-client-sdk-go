package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Searchshifttradesresponse
type Searchshifttradesresponse struct { 
	// Trades - The shift trades that match the search criteria
	Trades *[]Searchshifttraderesponse `json:"trades,omitempty"`

}

func (u *Searchshifttradesresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Searchshifttradesresponse

	

	return json.Marshal(&struct { 
		Trades *[]Searchshifttraderesponse `json:"trades,omitempty"`
		*Alias
	}{ 
		Trades: u.Trades,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Searchshifttradesresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
