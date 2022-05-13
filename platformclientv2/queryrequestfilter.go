package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryrequestfilter
type Queryrequestfilter struct { 
	// VarType - The logic used to combine the clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - The list of clauses used to filter the data
	Clauses *[]Queryrequestclause `json:"clauses,omitempty"`

}

func (o *Queryrequestfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryrequestfilter
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Clauses *[]Queryrequestclause `json:"clauses,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Clauses: o.Clauses,
		Alias:    (*Alias)(o),
	})
}

func (o *Queryrequestfilter) UnmarshalJSON(b []byte) error {
	var QueryrequestfilterMap map[string]interface{}
	err := json.Unmarshal(b, &QueryrequestfilterMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := QueryrequestfilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Clauses, ok := QueryrequestfilterMap["clauses"].([]interface{}); ok {
		ClausesString, _ := json.Marshal(Clauses)
		json.Unmarshal(ClausesString, &o.Clauses)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queryrequestfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
