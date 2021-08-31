package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Publishedsurveyformreference
type Publishedsurveyformreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ContextId - The context id of this form.
	ContextId *string `json:"contextId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Publishedsurveyformreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Publishedsurveyformreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ContextId: o.ContextId,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Publishedsurveyformreference) UnmarshalJSON(b []byte) error {
	var PublishedsurveyformreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &PublishedsurveyformreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PublishedsurveyformreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := PublishedsurveyformreferenceMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ContextId, ok := PublishedsurveyformreferenceMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
	
	if SelfUri, ok := PublishedsurveyformreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Publishedsurveyformreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
