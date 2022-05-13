package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Replacerequest
type Replacerequest struct { 
	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// AuthToken
	AuthToken *string `json:"authToken,omitempty"`

}

func (o *Replacerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Replacerequest
	
	return json.Marshal(&struct { 
		ChangeNumber *int `json:"changeNumber,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		AuthToken *string `json:"authToken,omitempty"`
		*Alias
	}{ 
		ChangeNumber: o.ChangeNumber,
		
		Name: o.Name,
		
		AuthToken: o.AuthToken,
		Alias:    (*Alias)(o),
	})
}

func (o *Replacerequest) UnmarshalJSON(b []byte) error {
	var ReplacerequestMap map[string]interface{}
	err := json.Unmarshal(b, &ReplacerequestMap)
	if err != nil {
		return err
	}
	
	if ChangeNumber, ok := ReplacerequestMap["changeNumber"].(float64); ok {
		ChangeNumberInt := int(ChangeNumber)
		o.ChangeNumber = &ChangeNumberInt
	}
	
	if Name, ok := ReplacerequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if AuthToken, ok := ReplacerequestMap["authToken"].(string); ok {
		o.AuthToken = &AuthToken
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Replacerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
