package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Unansweredgroup
type Unansweredgroup struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Label - Knowledge base unanswered group label
	Label *string `json:"label,omitempty"`


	// PhraseGroups - Represents a list of phrase groups inside an unanswered group
	PhraseGroups *[]Unansweredphrasegroup `json:"phraseGroups,omitempty"`


	// SuggestedDocuments - Represents a list of documents that may be linked to an unanswered group
	SuggestedDocuments *[]Unansweredgroupsuggesteddocument `json:"suggestedDocuments,omitempty"`


	// Statistics - Statistics object containing the various hit counts for an unanswered group
	Statistics *Knowledgegroupstatistics `json:"statistics,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Unansweredgroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Unansweredgroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Label *string `json:"label,omitempty"`
		
		PhraseGroups *[]Unansweredphrasegroup `json:"phraseGroups,omitempty"`
		
		SuggestedDocuments *[]Unansweredgroupsuggesteddocument `json:"suggestedDocuments,omitempty"`
		
		Statistics *Knowledgegroupstatistics `json:"statistics,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Label: o.Label,
		
		PhraseGroups: o.PhraseGroups,
		
		SuggestedDocuments: o.SuggestedDocuments,
		
		Statistics: o.Statistics,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Unansweredgroup) UnmarshalJSON(b []byte) error {
	var UnansweredgroupMap map[string]interface{}
	err := json.Unmarshal(b, &UnansweredgroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UnansweredgroupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Label, ok := UnansweredgroupMap["label"].(string); ok {
		o.Label = &Label
	}
    
	if PhraseGroups, ok := UnansweredgroupMap["phraseGroups"].([]interface{}); ok {
		PhraseGroupsString, _ := json.Marshal(PhraseGroups)
		json.Unmarshal(PhraseGroupsString, &o.PhraseGroups)
	}
	
	if SuggestedDocuments, ok := UnansweredgroupMap["suggestedDocuments"].([]interface{}); ok {
		SuggestedDocumentsString, _ := json.Marshal(SuggestedDocuments)
		json.Unmarshal(SuggestedDocumentsString, &o.SuggestedDocuments)
	}
	
	if Statistics, ok := UnansweredgroupMap["statistics"].(map[string]interface{}); ok {
		StatisticsString, _ := json.Marshal(Statistics)
		json.Unmarshal(StatisticsString, &o.Statistics)
	}
	
	if SelfUri, ok := UnansweredgroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Unansweredgroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
