package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentquery
type Documentquery struct { 
	// Clauses - Documents filter clauses/criteria. Limit of 20 clauses.
	Clauses *[]Documentqueryclause `json:"clauses,omitempty"`


	// Operator - Specifies how the filter clauses will be applied together.
	Operator *string `json:"operator,omitempty"`

}

func (o *Documentquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentquery
	
	return json.Marshal(&struct { 
		Clauses *[]Documentqueryclause `json:"clauses,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		*Alias
	}{ 
		Clauses: o.Clauses,
		
		Operator: o.Operator,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentquery) UnmarshalJSON(b []byte) error {
	var DocumentqueryMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentqueryMap)
	if err != nil {
		return err
	}
	
	if Clauses, ok := DocumentqueryMap["clauses"].([]interface{}); ok {
		ClausesString, _ := json.Marshal(Clauses)
		json.Unmarshal(ClausesString, &o.Clauses)
	}
	
	if Operator, ok := DocumentqueryMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Documentquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
