package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userprimarysource
type Userprimarysource struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SourceId - The id of the source
	SourceId *string `json:"sourceId,omitempty"`


	// Registered - Whether or not the source is registered
	Registered *bool `json:"registered,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Userprimarysource) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userprimarysource
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SourceId *string `json:"sourceId,omitempty"`
		
		Registered *bool `json:"registered,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SourceId: o.SourceId,
		
		Registered: o.Registered,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userprimarysource) UnmarshalJSON(b []byte) error {
	var UserprimarysourceMap map[string]interface{}
	err := json.Unmarshal(b, &UserprimarysourceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserprimarysourceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UserprimarysourceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SourceId, ok := UserprimarysourceMap["sourceId"].(string); ok {
		o.SourceId = &SourceId
	}
    
	if Registered, ok := UserprimarysourceMap["registered"].(bool); ok {
		o.Registered = &Registered
	}
    
	if SelfUri, ok := UserprimarysourceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userprimarysource) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
