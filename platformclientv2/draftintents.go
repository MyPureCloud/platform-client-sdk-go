package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Draftintents
type Draftintents struct { 
	// Id - Id for an intent.
	Id *string `json:"id,omitempty"`


	// Name - Name/Label for an intent.
	Name *string `json:"name,omitempty"`


	// Utterances - The utterances that are extracted for an Intent.
	Utterances *[]string `json:"utterances,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Draftintents) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Draftintents
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Utterances *[]string `json:"utterances,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Utterances: o.Utterances,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Draftintents) UnmarshalJSON(b []byte) error {
	var DraftintentsMap map[string]interface{}
	err := json.Unmarshal(b, &DraftintentsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DraftintentsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DraftintentsMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Utterances, ok := DraftintentsMap["utterances"].([]interface{}); ok {
		UtterancesString, _ := json.Marshal(Utterances)
		json.Unmarshal(UtterancesString, &o.Utterances)
	}
	
	if SelfUri, ok := DraftintentsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Draftintents) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
