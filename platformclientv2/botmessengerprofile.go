package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Botmessengerprofile
type Botmessengerprofile struct { 
	// Name - Name of the Bot
	Name *string `json:"name,omitempty"`


	// AvatarUrl - Avatar for Bot
	AvatarUrl *string `json:"avatarUrl,omitempty"`

}

func (o *Botmessengerprofile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Botmessengerprofile
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		AvatarUrl *string `json:"avatarUrl,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		AvatarUrl: o.AvatarUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Botmessengerprofile) UnmarshalJSON(b []byte) error {
	var BotmessengerprofileMap map[string]interface{}
	err := json.Unmarshal(b, &BotmessengerprofileMap)
	if err != nil {
		return err
	}
	
	if Name, ok := BotmessengerprofileMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if AvatarUrl, ok := BotmessengerprofileMap["avatarUrl"].(string); ok {
		o.AvatarUrl = &AvatarUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Botmessengerprofile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
