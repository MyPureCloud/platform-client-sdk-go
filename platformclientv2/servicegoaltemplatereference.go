package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Servicegoaltemplatereference
type Servicegoaltemplatereference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Servicegoaltemplatereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Servicegoaltemplatereference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Servicegoaltemplatereference) UnmarshalJSON(b []byte) error {
	var ServicegoaltemplatereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &ServicegoaltemplatereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ServicegoaltemplatereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if SelfUri, ok := ServicegoaltemplatereferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Servicegoaltemplatereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
