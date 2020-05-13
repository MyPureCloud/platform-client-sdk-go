package platformclientv2
import (
	"encoding/json"
)

// Searchshifttradesresponse
type Searchshifttradesresponse struct { 
	// Trades - The shift trades that match the search criteria
	Trades *[]Searchshifttraderesponse `json:"trades,omitempty"`

}

// String returns a JSON representation of the model
func (o *Searchshifttradesresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
