package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainresourceconditionvalue
type Domainresourceconditionvalue struct { 
	// User
	User *User `json:"user,omitempty"`


	// Queue
	Queue *Queue `json:"queue,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (u *Domainresourceconditionvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainresourceconditionvalue

	

	return json.Marshal(&struct { 
		User *User `json:"user,omitempty"`
		
		Queue *Queue `json:"queue,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		Queue: u.Queue,
		
		Value: u.Value,
		
		VarType: u.VarType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domainresourceconditionvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
