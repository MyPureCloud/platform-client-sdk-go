package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeimportjobrequest
type Knowledgeimportjobrequest struct { 
	// UploadKey - Upload key
	UploadKey *string `json:"uploadKey,omitempty"`


	// FileType - File type of the document
	FileType *string `json:"fileType,omitempty"`


	// Settings - Additional optional settings
	Settings *Knowledgeimportjobsettings `json:"settings,omitempty"`

}

func (o *Knowledgeimportjobrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeimportjobrequest
	
	return json.Marshal(&struct { 
		UploadKey *string `json:"uploadKey,omitempty"`
		
		FileType *string `json:"fileType,omitempty"`
		
		Settings *Knowledgeimportjobsettings `json:"settings,omitempty"`
		*Alias
	}{ 
		UploadKey: o.UploadKey,
		
		FileType: o.FileType,
		
		Settings: o.Settings,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeimportjobrequest) UnmarshalJSON(b []byte) error {
	var KnowledgeimportjobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeimportjobrequestMap)
	if err != nil {
		return err
	}
	
	if UploadKey, ok := KnowledgeimportjobrequestMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if FileType, ok := KnowledgeimportjobrequestMap["fileType"].(string); ok {
		o.FileType = &FileType
	}
    
	if Settings, ok := KnowledgeimportjobrequestMap["settings"].(map[string]interface{}); ok {
		SettingsString, _ := json.Marshal(Settings)
		json.Unmarshal(SettingsString, &o.Settings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
