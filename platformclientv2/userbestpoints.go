package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userbestpoints
type Userbestpoints struct { 
	// User - The requested user for the best points
	User *Userreference `json:"user,omitempty"`


	// BestPoints - List of best point for the requested user
	BestPoints *[]Userbestpointsitem `json:"bestPoints,omitempty"`

}

func (o *Userbestpoints) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userbestpoints
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		BestPoints *[]Userbestpointsitem `json:"bestPoints,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		BestPoints: o.BestPoints,
		Alias:    (*Alias)(o),
	})
}

func (o *Userbestpoints) UnmarshalJSON(b []byte) error {
	var UserbestpointsMap map[string]interface{}
	err := json.Unmarshal(b, &UserbestpointsMap)
	if err != nil {
		return err
	}
	
	if User, ok := UserbestpointsMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if BestPoints, ok := UserbestpointsMap["bestPoints"].([]interface{}); ok {
		BestPointsString, _ := json.Marshal(BestPoints)
		json.Unmarshal(BestPointsString, &o.BestPoints)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userbestpoints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
