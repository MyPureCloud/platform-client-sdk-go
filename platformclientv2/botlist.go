package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Botlist - A list of BotConnectorBots
type Botlist struct { 
	// ChatBots - A list of botConnector Bots. Max 50
	ChatBots *[]Botconnectorbot `json:"chatBots,omitempty"`

}

func (u *Botlist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Botlist

	

	return json.Marshal(&struct { 
		ChatBots *[]Botconnectorbot `json:"chatBots,omitempty"`
		*Alias
	}{ 
		ChatBots: u.ChatBots,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Botlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
