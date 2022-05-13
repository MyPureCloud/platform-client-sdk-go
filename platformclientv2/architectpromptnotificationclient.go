package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectpromptnotificationclient - The client who initiated the change.
type Architectpromptnotificationclient struct { 
	// Id - The ID of the client.
	Id *string `json:"id,omitempty"`


	// Name - The name of the client, if available.
	Name *string `json:"name,omitempty"`

}

func (o *Architectpromptnotificationclient) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectpromptnotificationclient
	
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

func (o *Architectpromptnotificationclient) UnmarshalJSON(b []byte) error {
	var ArchitectpromptnotificationclientMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectpromptnotificationclientMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectpromptnotificationclientMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ArchitectpromptnotificationclientMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Architectpromptnotificationclient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
