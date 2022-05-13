package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Audittopicdomainentityref
type Audittopicdomainentityref struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Audittopicdomainentityref) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Audittopicdomainentityref
	
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

func (o *Audittopicdomainentityref) UnmarshalJSON(b []byte) error {
	var AudittopicdomainentityrefMap map[string]interface{}
	err := json.Unmarshal(b, &AudittopicdomainentityrefMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AudittopicdomainentityrefMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AudittopicdomainentityrefMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SelfUri, ok := AudittopicdomainentityrefMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Audittopicdomainentityref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
