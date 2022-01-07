package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Invalidassignment
type Invalidassignment struct { 
	// User - Invalid user for validation
	User *Userreference `json:"user,omitempty"`


	// Message - The reason for the invalid input for validation
	Message *string `json:"message,omitempty"`

}

func (o *Invalidassignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Invalidassignment
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		Message *string `json:"message,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		Message: o.Message,
		Alias:    (*Alias)(o),
	})
}

func (o *Invalidassignment) UnmarshalJSON(b []byte) error {
	var InvalidassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &InvalidassignmentMap)
	if err != nil {
		return err
	}
	
	if User, ok := InvalidassignmentMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Message, ok := InvalidassignmentMap["message"].(string); ok {
		o.Message = &Message
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Invalidassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
