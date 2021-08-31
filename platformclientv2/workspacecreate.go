package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workspacecreate
type Workspacecreate struct { 
	// Name - The workspace name
	Name *string `json:"name,omitempty"`


	// Bucket
	Bucket *string `json:"bucket,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`

}

func (o *Workspacecreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workspacecreate
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Bucket *string `json:"bucket,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Bucket: o.Bucket,
		
		Description: o.Description,
		Alias:    (*Alias)(o),
	})
}

func (o *Workspacecreate) UnmarshalJSON(b []byte) error {
	var WorkspacecreateMap map[string]interface{}
	err := json.Unmarshal(b, &WorkspacecreateMap)
	if err != nil {
		return err
	}
	
	if Name, ok := WorkspacecreateMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Bucket, ok := WorkspacecreateMap["bucket"].(string); ok {
		o.Bucket = &Bucket
	}
	
	if Description, ok := WorkspacecreateMap["description"].(string); ok {
		o.Description = &Description
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workspacecreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
