package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappavailablephonenumberdetailslisting
type Whatsappavailablephonenumberdetailslisting struct { 
	// Entities
	Entities *[]Whatsappavailablephonenumberdetails `json:"entities,omitempty"`

}

func (o *Whatsappavailablephonenumberdetailslisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Whatsappavailablephonenumberdetailslisting
	
	return json.Marshal(&struct { 
		Entities *[]Whatsappavailablephonenumberdetails `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Whatsappavailablephonenumberdetailslisting) UnmarshalJSON(b []byte) error {
	var WhatsappavailablephonenumberdetailslistingMap map[string]interface{}
	err := json.Unmarshal(b, &WhatsappavailablephonenumberdetailslistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := WhatsappavailablephonenumberdetailslistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Whatsappavailablephonenumberdetailslisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
