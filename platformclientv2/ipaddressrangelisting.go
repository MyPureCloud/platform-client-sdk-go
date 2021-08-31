package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ipaddressrangelisting
type Ipaddressrangelisting struct { 
	// Entities
	Entities *[]Ipaddressrange `json:"entities,omitempty"`

}

func (o *Ipaddressrangelisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ipaddressrangelisting
	
	return json.Marshal(&struct { 
		Entities *[]Ipaddressrange `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Ipaddressrangelisting) UnmarshalJSON(b []byte) error {
	var IpaddressrangelistingMap map[string]interface{}
	err := json.Unmarshal(b, &IpaddressrangelistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := IpaddressrangelistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ipaddressrangelisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
