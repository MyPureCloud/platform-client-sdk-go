package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Botsearchresponse
type Botsearchresponse struct { 
	// Id - The id of the bot
	Id *string `json:"id,omitempty"`


	// Name - The name of the bot
	Name *string `json:"name,omitempty"`


	// BotType - The provider of the bot
	BotType *string `json:"botType,omitempty"`


	// Description - The description of the bot
	Description *string `json:"description,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Botsearchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Botsearchresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		BotType *string `json:"botType,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		BotType: o.BotType,
		
		Description: o.Description,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Botsearchresponse) UnmarshalJSON(b []byte) error {
	var BotsearchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BotsearchresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BotsearchresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := BotsearchresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if BotType, ok := BotsearchresponseMap["botType"].(string); ok {
		o.BotType = &BotType
	}
    
	if Description, ok := BotsearchresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if SelfUri, ok := BotsearchresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Botsearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
