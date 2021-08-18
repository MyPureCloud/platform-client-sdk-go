package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Botconnectorbot - A botConnector Bot Instance
type Botconnectorbot struct { 
	// Id - The Botconnector Bot Id - this is configurable by the user when put
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description - An optional description of the bot.  This can be up to 256 characters long and must be comprised of displayable characters without leading or trailing whitespace
	Description *string `json:"description,omitempty"`


	// Versions - This bots versions, limit of 50 per bot
	Versions *[]Botconnectorbotversion `json:"versions,omitempty"`


	// BotCompositeTag - A system-generated string that contains metadata about this bot.
	BotCompositeTag *string `json:"botCompositeTag,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Botconnectorbot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Botconnectorbot

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Versions *[]Botconnectorbotversion `json:"versions,omitempty"`
		
		BotCompositeTag *string `json:"botCompositeTag,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Versions: u.Versions,
		
		BotCompositeTag: u.BotCompositeTag,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Botconnectorbot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
