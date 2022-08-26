package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeexportjobfilter
type Knowledgeexportjobfilter struct { 
	// DocumentsFilter - Filters for narrowing down which documents to export.
	DocumentsFilter *Knowledgeexportjobdocumentsfilter `json:"documentsFilter,omitempty"`


	// VersionFilter - Specifies what version should be exported.
	VersionFilter *string `json:"versionFilter,omitempty"`

}

func (o *Knowledgeexportjobfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeexportjobfilter
	
	return json.Marshal(&struct { 
		DocumentsFilter *Knowledgeexportjobdocumentsfilter `json:"documentsFilter,omitempty"`
		
		VersionFilter *string `json:"versionFilter,omitempty"`
		*Alias
	}{ 
		DocumentsFilter: o.DocumentsFilter,
		
		VersionFilter: o.VersionFilter,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeexportjobfilter) UnmarshalJSON(b []byte) error {
	var KnowledgeexportjobfilterMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeexportjobfilterMap)
	if err != nil {
		return err
	}
	
	if DocumentsFilter, ok := KnowledgeexportjobfilterMap["documentsFilter"].(map[string]interface{}); ok {
		DocumentsFilterString, _ := json.Marshal(DocumentsFilter)
		json.Unmarshal(DocumentsFilterString, &o.DocumentsFilter)
	}
	
	if VersionFilter, ok := KnowledgeexportjobfilterMap["versionFilter"].(string); ok {
		o.VersionFilter = &VersionFilter
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeexportjobfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
