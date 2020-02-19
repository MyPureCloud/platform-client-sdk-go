package platformclientv2
import (
	"encoding/json"
)

// Ticker
type Ticker struct { 
	// Symbol - The ticker symbol for this organization. Example: ININ, AAPL, MSFT, etc.
	Symbol *string `json:"symbol,omitempty"`


	// Exchange - The exchange for this ticker symbol. Examples: NYSE, FTSE, NASDAQ, etc.
	Exchange *string `json:"exchange,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ticker) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
