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

func (o *Botconnectorbot) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Versions: o.Versions,
		
		BotCompositeTag: o.BotCompositeTag,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Botconnectorbot) UnmarshalJSON(b []byte) error {
	var BotconnectorbotMap map[string]interface{}
	err := json.Unmarshal(b, &BotconnectorbotMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BotconnectorbotMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := BotconnectorbotMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := BotconnectorbotMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Versions, ok := BotconnectorbotMap["versions"].([]interface{}); ok {
		VersionsString, _ := json.Marshal(Versions)
		json.Unmarshal(VersionsString, &o.Versions)
	}
	
	if BotCompositeTag, ok := BotconnectorbotMap["botCompositeTag"].(string); ok {
		o.BotCompositeTag = &BotCompositeTag
	}
    
	if SelfUri, ok := BotconnectorbotMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Botconnectorbot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
