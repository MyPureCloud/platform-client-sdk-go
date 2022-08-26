package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentsuggestionresultdocument
type Knowledgedocumentsuggestionresultdocument struct { 
	// Id - The globally unique identifier for the document.
	Id *string `json:"id,omitempty"`


	// KnowledgeBase - The knowledge base that the document belongs to.
	KnowledgeBase *Knowledgebasereference `json:"knowledgeBase,omitempty"`


	// Title - The title of the document.
	Title *string `json:"title,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Knowledgedocumentsuggestionresultdocument) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentsuggestionresultdocument
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		KnowledgeBase *Knowledgebasereference `json:"knowledgeBase,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		KnowledgeBase: o.KnowledgeBase,
		
		Title: o.Title,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentsuggestionresultdocument) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentsuggestionresultdocumentMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentsuggestionresultdocumentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgedocumentsuggestionresultdocumentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if KnowledgeBase, ok := KnowledgedocumentsuggestionresultdocumentMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if Title, ok := KnowledgedocumentsuggestionresultdocumentMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if SelfUri, ok := KnowledgedocumentsuggestionresultdocumentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsuggestionresultdocument) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
