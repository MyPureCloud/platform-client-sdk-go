package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Botlist - A list of BotConnectorBots
type Botlist struct { 
	// ChatBots - A list of botConnector Bots. Max 50
	ChatBots *[]Botconnectorbot `json:"chatBots,omitempty"`

}

// String returns a JSON representation of the model
func (o *Botlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
