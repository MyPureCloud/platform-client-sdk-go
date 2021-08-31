package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Botsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Botsummary
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		BotCompositeTag *string `json:"botCompositeTag,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Id: o.Id,
		
		Description: o.Description,
		
		BotCompositeTag: o.BotCompositeTag,
		Alias:    (*Alias)(o),
	})
}

func (o *Botsummary) UnmarshalJSON(b []byte) error {
	var BotsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &BotsummaryMap)
	if err != nil {
		return err
	}
	
	if Name, ok := BotsummaryMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Id, ok := BotsummaryMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Description, ok := BotsummaryMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if BotCompositeTag, ok := BotsummaryMap["botCompositeTag"].(string); ok {
		o.BotCompositeTag = &BotCompositeTag
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Botsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
