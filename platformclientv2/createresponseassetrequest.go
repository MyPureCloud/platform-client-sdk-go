package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createresponseassetrequest
type Createresponseassetrequest struct { 
	// Name - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
	Name *string `json:"name,omitempty"`


	// DivisionId - Division to associate to this asset. Can only be used with this division.
	DivisionId *string `json:"divisionId,omitempty"`


	// ContentMd5 - Content MD-5 of the file to upload
	ContentMd5 *string `json:"contentMd5,omitempty"`

}

func (o *Createresponseassetrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createresponseassetrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		ContentMd5 *string `json:"contentMd5,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		DivisionId: o.DivisionId,
		
		ContentMd5: o.ContentMd5,
		Alias:    (*Alias)(o),
	})
}

func (o *Createresponseassetrequest) UnmarshalJSON(b []byte) error {
	var CreateresponseassetrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreateresponseassetrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreateresponseassetrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if DivisionId, ok := CreateresponseassetrequestMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
	
	if ContentMd5, ok := CreateresponseassetrequestMap["contentMd5"].(string); ok {
		o.ContentMd5 = &ContentMd5
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createresponseassetrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
