package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeimportjobsettings
type Knowledgeimportjobsettings struct { 
	// ImportAsNew - If enabled import creates a new document even if update is available.
	ImportAsNew *bool `json:"importAsNew,omitempty"`


	// Visible - If specified, import will override the visibility of the imported documents.
	Visible *bool `json:"visible,omitempty"`


	// CategoryId - If specified, import will override the category of the imported documents.
	CategoryId *string `json:"categoryId,omitempty"`


	// LabelIds - If specified, import will add this labels to the imported documents.
	LabelIds *[]string `json:"labelIds,omitempty"`

}

func (o *Knowledgeimportjobsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeimportjobsettings
	
	return json.Marshal(&struct { 
		ImportAsNew *bool `json:"importAsNew,omitempty"`
		
		Visible *bool `json:"visible,omitempty"`
		
		CategoryId *string `json:"categoryId,omitempty"`
		
		LabelIds *[]string `json:"labelIds,omitempty"`
		*Alias
	}{ 
		ImportAsNew: o.ImportAsNew,
		
		Visible: o.Visible,
		
		CategoryId: o.CategoryId,
		
		LabelIds: o.LabelIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeimportjobsettings) UnmarshalJSON(b []byte) error {
	var KnowledgeimportjobsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeimportjobsettingsMap)
	if err != nil {
		return err
	}
	
	if ImportAsNew, ok := KnowledgeimportjobsettingsMap["importAsNew"].(bool); ok {
		o.ImportAsNew = &ImportAsNew
	}
    
	if Visible, ok := KnowledgeimportjobsettingsMap["visible"].(bool); ok {
		o.Visible = &Visible
	}
    
	if CategoryId, ok := KnowledgeimportjobsettingsMap["categoryId"].(string); ok {
		o.CategoryId = &CategoryId
	}
    
	if LabelIds, ok := KnowledgeimportjobsettingsMap["labelIds"].([]interface{}); ok {
		LabelIdsString, _ := json.Marshal(LabelIds)
		json.Unmarshal(LabelIdsString, &o.LabelIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
