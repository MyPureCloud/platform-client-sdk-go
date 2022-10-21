package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userreferencewithname
type Userreferencewithname struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Userreferencewithname) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userreferencewithname
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userreferencewithname) UnmarshalJSON(b []byte) error {
	var UserreferencewithnameMap map[string]interface{}
	err := json.Unmarshal(b, &UserreferencewithnameMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserreferencewithnameMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UserreferencewithnameMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SelfUri, ok := UserreferencewithnameMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userreferencewithname) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
