package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimuserroutinglanguage - The routing language assigned to a user.
type Scimuserroutinglanguage struct { 
	// Name - The case-sensitive name of a routing language configured in Genesys Cloud.
	Name *string `json:"name,omitempty"`


	// Proficiency - A rating from 0.0 to 5.0 that indicates how fluent an agent is in a particular language. ACD interactions are routed to agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`

}

func (o *Scimuserroutinglanguage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimuserroutinglanguage
	
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

func (o *Scimuserroutinglanguage) UnmarshalJSON(b []byte) error {
	var ScimuserroutinglanguageMap map[string]interface{}
	err := json.Unmarshal(b, &ScimuserroutinglanguageMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ScimuserroutinglanguageMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Proficiency, ok := ScimuserroutinglanguageMap["proficiency"].(float64); ok {
		o.Proficiency = &Proficiency
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Scimuserroutinglanguage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
