package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Botversionsummary - A version summary for a botConnector bot.
type Botversionsummary struct { 
	// Name - The name of the bot.
	Name *string `json:"name,omitempty"`


	// Id - The id of the bot.
	Id *string `json:"id,omitempty"`


	// Description - An optional description of the bot.
	Description *string `json:"description,omitempty"`


	// BotCompositeTag - A system-generated string that contains metadata about this bot.
	BotCompositeTag *string `json:"botCompositeTag,omitempty"`


	// Version - The name of the version.
	Version *string `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Botversionsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
