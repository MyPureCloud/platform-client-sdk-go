package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responseassetrequest
type Responseassetrequest struct { 
	// Name - Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|
	Name *string `json:"name,omitempty"`


	// DivisionId - Division to associate to this asset. Can only be used with this division.
	DivisionId *string `json:"divisionId,omitempty"`

}

func (o *Responseassetrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responseassetrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		DivisionId: o.DivisionId,
		Alias:    (*Alias)(o),
	})
}

func (o *Responseassetrequest) UnmarshalJSON(b []byte) error {
	var ResponseassetrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ResponseassetrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ResponseassetrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if DivisionId, ok := ResponseassetrequestMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Responseassetrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
