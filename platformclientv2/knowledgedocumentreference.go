package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentreference
type Knowledgedocumentreference struct { 
	// Id - The globally unique identifier for the document.
	Id *string `json:"id,omitempty"`


	// KnowledgeBase - The knowledge base that the document belongs to.
	KnowledgeBase *Knowledgebasereference `json:"knowledgeBase,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Knowledgedocumentreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		KnowledgeBase *Knowledgebasereference `json:"knowledgeBase,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		KnowledgeBase: o.KnowledgeBase,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentreference) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgedocumentreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if KnowledgeBase, ok := KnowledgedocumentreferenceMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if SelfUri, ok := KnowledgedocumentreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
