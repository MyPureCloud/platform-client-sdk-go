package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userlicenses
type Userlicenses struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Licenses
	Licenses *[]string `json:"licenses,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Userlicenses) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userlicenses
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Licenses *[]string `json:"licenses,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Licenses: o.Licenses,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userlicenses) UnmarshalJSON(b []byte) error {
	var UserlicensesMap map[string]interface{}
	err := json.Unmarshal(b, &UserlicensesMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserlicensesMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Licenses, ok := UserlicensesMap["licenses"].([]interface{}); ok {
		LicensesString, _ := json.Marshal(Licenses)
		json.Unmarshal(LicensesString, &o.Licenses)
	}
	
	if SelfUri, ok := UserlicensesMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userlicenses) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
