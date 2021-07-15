package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Botsummary - A summary description for a botConnector bot
type Botsummary struct { 
	// Name - The name of the bot.
	Name *string `json:"name,omitempty"`


	// Id - The id of the bot.
	Id *string `json:"id,omitempty"`


	// Description - An optional description of the bot.
	Description *string `json:"description,omitempty"`


	// BotCompositeTag - A system-generated string that contains metadata about this bot.
	BotCompositeTag *string `json:"botCompositeTag,omitempty"`

}

// String returns a JSON representation of the model
func (o *Botsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
