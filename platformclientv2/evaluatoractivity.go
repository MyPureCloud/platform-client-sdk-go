package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluatoractivity
type Evaluatoractivity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Evaluator
	Evaluator *User `json:"evaluator,omitempty"`


	// NumEvaluationsAssigned
	NumEvaluationsAssigned *int `json:"numEvaluationsAssigned,omitempty"`


	// NumEvaluationsStarted
	NumEvaluationsStarted *int `json:"numEvaluationsStarted,omitempty"`


	// NumEvaluationsCompleted
	NumEvaluationsCompleted *int `json:"numEvaluationsCompleted,omitempty"`


	// NumCalibrationsAssigned
	NumCalibrationsAssigned *int `json:"numCalibrationsAssigned,omitempty"`


	// NumCalibrationsStarted
	NumCalibrationsStarted *int `json:"numCalibrationsStarted,omitempty"`


	// NumCalibrationsCompleted
	NumCalibrationsCompleted *int `json:"numCalibrationsCompleted,omitempty"`


	// NumEvaluationsWithoutViewPermission
	NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Evaluatoractivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluatoractivity

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Evaluator *User `json:"evaluator,omitempty"`
		
		NumEvaluationsAssigned *int `json:"numEvaluationsAssigned,omitempty"`
		
		NumEvaluationsStarted *int `json:"numEvaluationsStarted,omitempty"`
		
		NumEvaluationsCompleted *int `json:"numEvaluationsCompleted,omitempty"`
		
		NumCalibrationsAssigned *int `json:"numCalibrationsAssigned,omitempty"`
		
		NumCalibrationsStarted *int `json:"numCalibrationsStarted,omitempty"`
		
		NumCalibrationsCompleted *int `json:"numCalibrationsCompleted,omitempty"`
		
		NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Evaluator: u.Evaluator,
		
		NumEvaluationsAssigned: u.NumEvaluationsAssigned,
		
		NumEvaluationsStarted: u.NumEvaluationsStarted,
		
		NumEvaluationsCompleted: u.NumEvaluationsCompleted,
		
		NumCalibrationsAssigned: u.NumCalibrationsAssigned,
		
		NumCalibrationsStarted: u.NumCalibrationsStarted,
		
		NumCalibrationsCompleted: u.NumCalibrationsCompleted,
		
		NumEvaluationsWithoutViewPermission: u.NumEvaluationsWithoutViewPermission,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluatoractivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
