package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturnaction
type Reportingturnaction struct { 
	// ActionId - The ID of the action in the bot flow.
	ActionId *string `json:"actionId,omitempty"`


	// ActionName - The name of the action in the bot flow.
	ActionName *string `json:"actionName,omitempty"`


	// ActionNumber - The number of the action in the bot flow.
	ActionNumber *int `json:"actionNumber,omitempty"`


	// ActionType
	ActionType *string `json:"actionType,omitempty"`

}

func (u *Reportingturnaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingturnaction

	

	return json.Marshal(&struct { 
		ActionId *string `json:"actionId,omitempty"`
		
		ActionName *string `json:"actionName,omitempty"`
		
		ActionNumber *int `json:"actionNumber,omitempty"`
		
		ActionType *string `json:"actionType,omitempty"`
		*Alias
	}{ 
		ActionId: u.ActionId,
		
		ActionName: u.ActionName,
		
		ActionNumber: u.ActionNumber,
		
		ActionType: u.ActionType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportingturnaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
