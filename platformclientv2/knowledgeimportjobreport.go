package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeimportjobreport
type Knowledgeimportjobreport struct { 
	// Errors - List of errors occurred during processing import.
	Errors *[]Knowledgeimportjoberror `json:"errors,omitempty"`


	// Statistics - Statistics related to the import job.
	Statistics *Knowledgeimportjobstatistics `json:"statistics,omitempty"`

}

func (o *Knowledgeimportjobreport) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeimportjobreport
	
	return json.Marshal(&struct { 
		Errors *[]Knowledgeimportjoberror `json:"errors,omitempty"`
		
		Statistics *Knowledgeimportjobstatistics `json:"statistics,omitempty"`
		*Alias
	}{ 
		Errors: o.Errors,
		
		Statistics: o.Statistics,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeimportjobreport) UnmarshalJSON(b []byte) error {
	var KnowledgeimportjobreportMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeimportjobreportMap)
	if err != nil {
		return err
	}
	
	if Errors, ok := KnowledgeimportjobreportMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	
	if Statistics, ok := KnowledgeimportjobreportMap["statistics"].(map[string]interface{}); ok {
		StatisticsString, _ := json.Marshal(Statistics)
		json.Unmarshal(StatisticsString, &o.Statistics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobreport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
