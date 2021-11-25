package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Usertimeoffrequestreference
type Usertimeoffrequestreference struct { 
	// Id - The id of the time off request
	Id *string `json:"id,omitempty"`


	// User - The ID of the user to whom the time off request applies
	User *Userreference `json:"user,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Usertimeoffrequestreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Usertimeoffrequestreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		User: o.User,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Usertimeoffrequestreference) UnmarshalJSON(b []byte) error {
	var UsertimeoffrequestreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &UsertimeoffrequestreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UsertimeoffrequestreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if User, ok := UsertimeoffrequestreferenceMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if SelfUri, ok := UsertimeoffrequestreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Usertimeoffrequestreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
