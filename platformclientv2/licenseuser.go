package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Licenseuser
type Licenseuser struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Licenses
	Licenses *[]Licensedefinition `json:"licenses,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Licenseuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Licenseuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Licenses *[]Licensedefinition `json:"licenses,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Licenses: o.Licenses,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Licenseuser) UnmarshalJSON(b []byte) error {
	var LicenseuserMap map[string]interface{}
	err := json.Unmarshal(b, &LicenseuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LicenseuserMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Licenses, ok := LicenseuserMap["licenses"].([]interface{}); ok {
		LicensesString, _ := json.Marshal(Licenses)
		json.Unmarshal(LicensesString, &o.Licenses)
	}
	
	if SelfUri, ok := LicenseuserMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Licenseuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
