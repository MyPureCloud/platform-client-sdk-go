package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignconfigchangeurireference
type Dialercampaignconfigchangeurireference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

func (o *Dialercampaignconfigchangeurireference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignconfigchangeurireference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignconfigchangeurireference) UnmarshalJSON(b []byte) error {
	var DialercampaignconfigchangeurireferenceMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignconfigchangeurireferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialercampaignconfigchangeurireferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DialercampaignconfigchangeurireferenceMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangeurireference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
