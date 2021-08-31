package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adjacents
type Adjacents struct { 
	// Superiors
	Superiors *[]User `json:"superiors,omitempty"`


	// Siblings
	Siblings *[]User `json:"siblings,omitempty"`


	// DirectReports
	DirectReports *[]User `json:"directReports,omitempty"`

}

func (o *Adjacents) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adjacents
	
	return json.Marshal(&struct { 
		Superiors *[]User `json:"superiors,omitempty"`
		
		Siblings *[]User `json:"siblings,omitempty"`
		
		DirectReports *[]User `json:"directReports,omitempty"`
		*Alias
	}{ 
		Superiors: o.Superiors,
		
		Siblings: o.Siblings,
		
		DirectReports: o.DirectReports,
		Alias:    (*Alias)(o),
	})
}

func (o *Adjacents) UnmarshalJSON(b []byte) error {
	var AdjacentsMap map[string]interface{}
	err := json.Unmarshal(b, &AdjacentsMap)
	if err != nil {
		return err
	}
	
	if Superiors, ok := AdjacentsMap["superiors"].([]interface{}); ok {
		SuperiorsString, _ := json.Marshal(Superiors)
		json.Unmarshal(SuperiorsString, &o.Superiors)
	}
	
	if Siblings, ok := AdjacentsMap["siblings"].([]interface{}); ok {
		SiblingsString, _ := json.Marshal(Siblings)
		json.Unmarshal(SiblingsString, &o.Siblings)
	}
	
	if DirectReports, ok := AdjacentsMap["directReports"].([]interface{}); ok {
		DirectReportsString, _ := json.Marshal(DirectReports)
		json.Unmarshal(DirectReportsString, &o.DirectReports)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Adjacents) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
