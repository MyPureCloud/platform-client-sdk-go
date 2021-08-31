package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Channelentitylisting
type Channelentitylisting struct { 
	// Entities
	Entities *[]Channel `json:"entities,omitempty"`

}

func (o *Channelentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Channelentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Channel `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Channelentitylisting) UnmarshalJSON(b []byte) error {
	var ChannelentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &ChannelentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ChannelentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Channelentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
