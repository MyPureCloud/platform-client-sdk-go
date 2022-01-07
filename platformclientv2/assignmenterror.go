package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assignmenterror
type Assignmenterror struct { 
	// User - A user that is failed to be removed from the performance profile
	User *Userreference `json:"user,omitempty"`


	// Message - Error message from membership assignment
	Message *string `json:"message,omitempty"`

}

func (o *Assignmenterror) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assignmenterror
	
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

func (o *Assignmenterror) UnmarshalJSON(b []byte) error {
	var AssignmenterrorMap map[string]interface{}
	err := json.Unmarshal(b, &AssignmenterrorMap)
	if err != nil {
		return err
	}
	
	if User, ok := AssignmenterrorMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Message, ok := AssignmenterrorMap["message"].(string); ok {
		o.Message = &Message
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assignmenterror) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
