package platformclientv2
import (
	"encoding/json"
)

// Genesysbotconnector
type Genesysbotconnector struct { 
	// QueryParameters - User defined name/value parameters passed to the BotConnector bot.
	QueryParameters *map[string]string `json:"queryParameters,omitempty"`

}

// String returns a JSON representation of the model
func (o *Genesysbotconnector) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
