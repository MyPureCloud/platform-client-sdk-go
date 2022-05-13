package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ticker
type Ticker struct { 
	// Symbol - The ticker symbol for this organization. Example: ININ, AAPL, MSFT, etc.
	Symbol *string `json:"symbol,omitempty"`


	// Exchange - The exchange for this ticker symbol. Examples: NYSE, FTSE, NASDAQ, etc.
	Exchange *string `json:"exchange,omitempty"`

}

func (o *Ticker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ticker
	
	return json.Marshal(&struct { 
		Symbol *string `json:"symbol,omitempty"`
		
		Exchange *string `json:"exchange,omitempty"`
		*Alias
	}{ 
		Symbol: o.Symbol,
		
		Exchange: o.Exchange,
		Alias:    (*Alias)(o),
	})
}

func (o *Ticker) UnmarshalJSON(b []byte) error {
	var TickerMap map[string]interface{}
	err := json.Unmarshal(b, &TickerMap)
	if err != nil {
		return err
	}
	
	if Symbol, ok := TickerMap["symbol"].(string); ok {
		o.Symbol = &Symbol
	}
    
	if Exchange, ok := TickerMap["exchange"].(string); ok {
		o.Exchange = &Exchange
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Ticker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
