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

func (u *Documentupload) MarshalJSON() ([]byte, error) {
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
		Name: u.Name,
		
		Workspace: u.Workspace,
		
		Tags: u.Tags,
		
		TagIds: u.TagIds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Documentupload) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
