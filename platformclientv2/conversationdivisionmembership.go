package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationdivisionmembership
type Conversationdivisionmembership struct { 
	// Division - A division the conversation belongs to.
	Division *Domainentityref `json:"division,omitempty"`


	// Entities - The entities on the conversation within the division. These are the users, queues, work flows, etc. that can be on conversations and and be assigned to different divisions.
	Entities *[]Domainentityref `json:"entities,omitempty"`

}

func (o *Conversationdivisionmembership) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationdivisionmembership
	
	return json.Marshal(&struct { 
		Division *Domainentityref `json:"division,omitempty"`
		
		Entities *[]Domainentityref `json:"entities,omitempty"`
		*Alias
	}{ 
		Division: o.Division,
		
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationdivisionmembership) UnmarshalJSON(b []byte) error {
	var ConversationdivisionmembershipMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationdivisionmembershipMap)
	if err != nil {
		return err
	}
	
	if Division, ok := ConversationdivisionmembershipMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Entities, ok := ConversationdivisionmembershipMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationdivisionmembership) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
