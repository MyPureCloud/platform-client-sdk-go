package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgegroupstatistics
type Knowledgegroupstatistics struct { 
	// UnlinkedPhraseCount - Knowledge Group unique phrase count
	UnlinkedPhraseCount *int `json:"unlinkedPhraseCount,omitempty"`


	// UnlinkedPhraseHitCount - Knowledge Group unlinked phrases hit count
	UnlinkedPhraseHitCount *int `json:"unlinkedPhraseHitCount,omitempty"`


	// TotalPhraseHitCount - Total number of phrase hit counts of an unanswered group
	TotalPhraseHitCount *int `json:"totalPhraseHitCount,omitempty"`

}

func (o *Knowledgegroupstatistics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgegroupstatistics
	
	return json.Marshal(&struct { 
		UnlinkedPhraseCount *int `json:"unlinkedPhraseCount,omitempty"`
		
		UnlinkedPhraseHitCount *int `json:"unlinkedPhraseHitCount,omitempty"`
		
		TotalPhraseHitCount *int `json:"totalPhraseHitCount,omitempty"`
		*Alias
	}{ 
		UnlinkedPhraseCount: o.UnlinkedPhraseCount,
		
		UnlinkedPhraseHitCount: o.UnlinkedPhraseHitCount,
		
		TotalPhraseHitCount: o.TotalPhraseHitCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgegroupstatistics) UnmarshalJSON(b []byte) error {
	var KnowledgegroupstatisticsMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgegroupstatisticsMap)
	if err != nil {
		return err
	}
	
	if UnlinkedPhraseCount, ok := KnowledgegroupstatisticsMap["unlinkedPhraseCount"].(float64); ok {
		UnlinkedPhraseCountInt := int(UnlinkedPhraseCount)
		o.UnlinkedPhraseCount = &UnlinkedPhraseCountInt
	}
	
	if UnlinkedPhraseHitCount, ok := KnowledgegroupstatisticsMap["unlinkedPhraseHitCount"].(float64); ok {
		UnlinkedPhraseHitCountInt := int(UnlinkedPhraseHitCount)
		o.UnlinkedPhraseHitCount = &UnlinkedPhraseHitCountInt
	}
	
	if TotalPhraseHitCount, ok := KnowledgegroupstatisticsMap["totalPhraseHitCount"].(float64); ok {
		TotalPhraseHitCountInt := int(TotalPhraseHitCount)
		o.TotalPhraseHitCount = &TotalPhraseHitCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgegroupstatistics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
