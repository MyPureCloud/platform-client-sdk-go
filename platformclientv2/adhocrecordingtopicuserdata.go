package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adhocrecordingtopicuserdata
type Adhocrecordingtopicuserdata struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

func (o *Adhocrecordingtopicuserdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adhocrecordingtopicuserdata
	
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

func (o *Adhocrecordingtopicuserdata) UnmarshalJSON(b []byte) error {
	var AdhocrecordingtopicuserdataMap map[string]interface{}
	err := json.Unmarshal(b, &AdhocrecordingtopicuserdataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AdhocrecordingtopicuserdataMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AdhocrecordingtopicuserdataMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Adhocrecordingtopicuserdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
