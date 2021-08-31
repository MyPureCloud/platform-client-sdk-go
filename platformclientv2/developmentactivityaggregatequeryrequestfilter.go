package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregatequeryrequestfilter
type Developmentactivityaggregatequeryrequestfilter struct { 
	// VarType - The logic used to combine the clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - The list of clauses used to filter the data. Note that clauses must filter by attendeeId and a maximum of 100 user IDs are allowed
	Clauses *[]Developmentactivityaggregatequeryrequestclause `json:"clauses,omitempty"`

}

func (o *Developmentactivityaggregatequeryrequestfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivityaggregatequeryrequestfilter
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Clauses *[]Developmentactivityaggregatequeryrequestclause `json:"clauses,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Clauses: o.Clauses,
		Alias:    (*Alias)(o),
	})
}

func (o *Developmentactivityaggregatequeryrequestfilter) UnmarshalJSON(b []byte) error {
	var DevelopmentactivityaggregatequeryrequestfilterMap map[string]interface{}
	err := json.Unmarshal(b, &DevelopmentactivityaggregatequeryrequestfilterMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DevelopmentactivityaggregatequeryrequestfilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Clauses, ok := DevelopmentactivityaggregatequeryrequestfilterMap["clauses"].([]interface{}); ok {
		ClausesString, _ := json.Marshal(Clauses)
		json.Unmarshal(ClausesString, &o.Clauses)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryrequestfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
