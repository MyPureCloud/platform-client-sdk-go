package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentreq
type Knowledgedocumentreq struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Title - Document title.
	Title *string `json:"title,omitempty"`


	// Visible - Indicates if the knowledge document should be included in search results.
	Visible *bool `json:"visible,omitempty"`


	// Alternatives - List of alternate phrases related to the title which improves search results.
	Alternatives *[]Knowledgedocumentalternative `json:"alternatives,omitempty"`


	// CategoryId - The category associated with the document.
	CategoryId *string `json:"categoryId,omitempty"`


	// LabelIds - The ids of labels associated with the document.
	LabelIds *[]string `json:"labelIds,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Knowledgedocumentreq) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentreq
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Visible *bool `json:"visible,omitempty"`
		
		Alternatives *[]Knowledgedocumentalternative `json:"alternatives,omitempty"`
		
		CategoryId *string `json:"categoryId,omitempty"`
		
		LabelIds *[]string `json:"labelIds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Title: o.Title,
		
		Visible: o.Visible,
		
		Alternatives: o.Alternatives,
		
		CategoryId: o.CategoryId,
		
		LabelIds: o.LabelIds,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentreq) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentreqMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentreqMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgedocumentreqMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Title, ok := KnowledgedocumentreqMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Visible, ok := KnowledgedocumentreqMap["visible"].(bool); ok {
		o.Visible = &Visible
	}
    
	if Alternatives, ok := KnowledgedocumentreqMap["alternatives"].([]interface{}); ok {
		AlternativesString, _ := json.Marshal(Alternatives)
		json.Unmarshal(AlternativesString, &o.Alternatives)
	}
	
	if CategoryId, ok := KnowledgedocumentreqMap["categoryId"].(string); ok {
		o.CategoryId = &CategoryId
	}
    
	if LabelIds, ok := KnowledgedocumentreqMap["labelIds"].([]interface{}); ok {
		LabelIdsString, _ := json.Marshal(LabelIds)
		json.Unmarshal(LabelIdsString, &o.LabelIds)
	}
	
	if SelfUri, ok := KnowledgedocumentreqMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentreq) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
