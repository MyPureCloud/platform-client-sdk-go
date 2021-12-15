package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationclient - The client who initiated the change.
type Architectflownotificationclient struct { 
	// Id - The ID of the client.
	Id *string `json:"id,omitempty"`


	// Name - The name of the client, if available.
	Name *string `json:"name,omitempty"`

}

func (o *Architectflownotificationclient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflownotificationclient
	
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

func (o *Architectflownotificationclient) UnmarshalJSON(b []byte) error {
	var ArchitectflownotificationclientMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflownotificationclientMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectflownotificationclientMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ArchitectflownotificationclientMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflownotificationclient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
