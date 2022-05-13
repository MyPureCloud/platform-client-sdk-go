package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentupload
type Documentupload struct { 
	// Name - The name of the document
	Name *string `json:"name,omitempty"`


	// Workspace - The workspace the document will be uploaded to
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// TagIds
	TagIds *[]string `json:"tagIds,omitempty"`

}

func (o *Documentupload) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentupload
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		TagIds *[]string `json:"tagIds,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Workspace: o.Workspace,
		
		Tags: o.Tags,
		
		TagIds: o.TagIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentupload) UnmarshalJSON(b []byte) error {
	var DocumentuploadMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentuploadMap)
	if err != nil {
		return err
	}
	
	if Name, ok := DocumentuploadMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Workspace, ok := DocumentuploadMap["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if Tags, ok := DocumentuploadMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if TagIds, ok := DocumentuploadMap["tagIds"].([]interface{}); ok {
		TagIdsString, _ := json.Marshal(TagIds)
		json.Unmarshal(TagIdsString, &o.TagIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentupload) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
