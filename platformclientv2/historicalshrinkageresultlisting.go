package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalshrinkageresultlisting
type Historicalshrinkageresultlisting struct { 
	// Entities
	Entities *[]Historicalshrinkageresult `json:"entities,omitempty"`

}

func (o *Historicalshrinkageresultlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicalshrinkageresultlisting
	
	return json.Marshal(&struct { 
		Entities *[]Historicalshrinkageresult `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Historicalshrinkageresultlisting) UnmarshalJSON(b []byte) error {
	var HistoricalshrinkageresultlistingMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricalshrinkageresultlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := HistoricalshrinkageresultlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicalshrinkageresultlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
