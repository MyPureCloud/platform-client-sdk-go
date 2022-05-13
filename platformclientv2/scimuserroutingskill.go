package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimuserroutingskill - The routing skill assigned to a user.
type Scimuserroutingskill struct { 
	// Name - The case-sensitive name of a routing skill configured in Genesys Cloud.
	Name *string `json:"name,omitempty"`


	// Proficiency - A rating from 0.0 to 5.0 that indicates how adept an agent is at a particular skill. When \"Best available skills\" is enabled for a queue in Genesys Cloud, ACD interactions in that queue are routed to agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`

}

func (o *Scimuserroutingskill) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimuserroutingskill
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Proficiency *float64 `json:"proficiency,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Proficiency: o.Proficiency,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimuserroutingskill) UnmarshalJSON(b []byte) error {
	var ScimuserroutingskillMap map[string]interface{}
	err := json.Unmarshal(b, &ScimuserroutingskillMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ScimuserroutingskillMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Proficiency, ok := ScimuserroutingskillMap["proficiency"].(float64); ok {
		o.Proficiency = &Proficiency
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Scimuserroutingskill) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
