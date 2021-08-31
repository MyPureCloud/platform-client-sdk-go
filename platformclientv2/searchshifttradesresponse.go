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

func (o *Searchshifttradesresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Searchshifttradesresponse
	
	return json.Marshal(&struct { 
		Trades *[]Searchshifttraderesponse `json:"trades,omitempty"`
		*Alias
	}{ 
		Trades: o.Trades,
		Alias:    (*Alias)(o),
	})
}

func (o *Searchshifttradesresponse) UnmarshalJSON(b []byte) error {
	var SearchshifttradesresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SearchshifttradesresponseMap)
	if err != nil {
		return err
	}
	
	if Trades, ok := SearchshifttradesresponseMap["trades"].([]interface{}); ok {
		TradesString, _ := json.Marshal(Trades)
		json.Unmarshal(TradesString, &o.Trades)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Searchshifttradesresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
