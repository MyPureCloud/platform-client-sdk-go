package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatroutingtarget
type Webchatroutingtarget struct { 
	// TargetType - The target type of the routing target, such as 'QUEUE'.
	TargetType *string `json:"targetType,omitempty"`


	// TargetAddress - The target of the route, in the format appropriate given the 'targetType'.
	TargetAddress *string `json:"targetAddress,omitempty"`


	// Skills - The list of skill names to use for routing.
	Skills *[]string `json:"skills,omitempty"`


	// Language - The language name to use for routing.
	Language *string `json:"language,omitempty"`


	// Priority - The priority to assign to the conversation for routing.
	Priority *int `json:"priority,omitempty"`

}

func (u *Webchatroutingtarget) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatroutingtarget

	

	return json.Marshal(&struct { 
		TargetType *string `json:"targetType,omitempty"`
		
		TargetAddress *string `json:"targetAddress,omitempty"`
		
		Skills *[]string `json:"skills,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		*Alias
	}{ 
		TargetType: u.TargetType,
		
		TargetAddress: u.TargetAddress,
		
		Skills: u.Skills,
		
		Language: u.Language,
		
		Priority: u.Priority,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webchatroutingtarget) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
