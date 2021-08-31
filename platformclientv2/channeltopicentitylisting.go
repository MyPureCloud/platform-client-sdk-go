package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Channeltopicentitylisting
type Channeltopicentitylisting struct { 
	// Entities
	Entities *[]Channeltopic `json:"entities,omitempty"`

}

func (o *Channeltopicentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Channeltopicentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Channeltopic `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Channeltopicentitylisting) UnmarshalJSON(b []byte) error {
	var ChanneltopicentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &ChanneltopicentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ChanneltopicentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Channeltopicentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
