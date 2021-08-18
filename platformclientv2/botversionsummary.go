package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Botversionsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Botversionsummary

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		BotCompositeTag *string `json:"botCompositeTag,omitempty"`
		
		Version *string `json:"version,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Id: u.Id,
		
		Description: u.Description,
		
		BotCompositeTag: u.BotCompositeTag,
		
		Version: u.Version,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Botversionsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
