package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregatequeryrequestfilter
type Learningassignmentaggregatequeryrequestfilter struct { 
	// VarType - The logic used to combine the clauses
	VarType *string `json:"type,omitempty"`


	// Clauses - The list of clauses used to filter the data. Note that clauses must filter by attendeeId and a maximum of 100 user IDs are allowed
	Clauses *[]Learningassignmentaggregatequeryrequestclause `json:"clauses,omitempty"`

}

func (o *Learningassignmentaggregatequeryrequestfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentaggregatequeryrequestfilter
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Clauses *[]Learningassignmentaggregatequeryrequestclause `json:"clauses,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Clauses: o.Clauses,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentaggregatequeryrequestfilter) UnmarshalJSON(b []byte) error {
	var LearningassignmentaggregatequeryrequestfilterMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentaggregatequeryrequestfilterMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := LearningassignmentaggregatequeryrequestfilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Clauses, ok := LearningassignmentaggregatequeryrequestfilterMap["clauses"].([]interface{}); ok {
		ClausesString, _ := json.Marshal(Clauses)
		json.Unmarshal(ClausesString, &o.Clauses)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryrequestfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
