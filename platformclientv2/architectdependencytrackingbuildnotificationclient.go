package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectdependencytrackingbuildnotificationclient
type Architectdependencytrackingbuildnotificationclient struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

func (o *Architectdependencytrackingbuildnotificationclient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectdependencytrackingbuildnotificationclient
	
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

func (o *Architectdependencytrackingbuildnotificationclient) UnmarshalJSON(b []byte) error {
	var ArchitectdependencytrackingbuildnotificationclientMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectdependencytrackingbuildnotificationclientMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectdependencytrackingbuildnotificationclientMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ArchitectdependencytrackingbuildnotificationclientMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectdependencytrackingbuildnotificationclient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
