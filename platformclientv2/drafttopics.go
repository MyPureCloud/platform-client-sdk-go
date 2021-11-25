package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Drafttopics
type Drafttopics struct { 
	// Id - Id for a topic.
	Id *string `json:"id,omitempty"`


	// Name - Name/Label for a topic.
	Name *string `json:"name,omitempty"`


	// Phrases - The phrases that are extracted for a topic.
	Phrases *[]string `json:"phrases,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Drafttopics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Drafttopics
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Phrases *[]string `json:"phrases,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Phrases: o.Phrases,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Drafttopics) UnmarshalJSON(b []byte) error {
	var DrafttopicsMap map[string]interface{}
	err := json.Unmarshal(b, &DrafttopicsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DrafttopicsMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DrafttopicsMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Phrases, ok := DrafttopicsMap["phrases"].([]interface{}); ok {
		PhrasesString, _ := json.Marshal(Phrases)
		json.Unmarshal(PhrasesString, &o.Phrases)
	}
	
	if SelfUri, ok := DrafttopicsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Drafttopics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
