package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Calibrationcreate
type Calibrationcreate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Calibrator
	Calibrator *User `json:"calibrator,omitempty"`


	// Agent
	Agent *User `json:"agent,omitempty"`


	// Conversation - The conversation to use for the calibration.
	Conversation *Conversation `json:"conversation,omitempty"`


	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// ContextId
	ContextId *string `json:"contextId,omitempty"`


	// AverageScore
	AverageScore *int `json:"averageScore,omitempty"`


	// HighScore
	HighScore *int `json:"highScore,omitempty"`


	// LowScore
	LowScore *int `json:"lowScore,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// Evaluations
	Evaluations *[]Evaluation `json:"evaluations,omitempty"`


	// Evaluators
	Evaluators *[]User `json:"evaluators,omitempty"`


	// ScoringIndex
	ScoringIndex *Evaluation `json:"scoringIndex,omitempty"`


	// ExpertEvaluator
	ExpertEvaluator *User `json:"expertEvaluator,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Calibrationcreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Calibrationcreate

	
	CreatedDate := new(string)
	if u.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(u.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Calibrator *User `json:"calibrator,omitempty"`
		
		Agent *User `json:"agent,omitempty"`
		
		Conversation *Conversation `json:"conversation,omitempty"`
		
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		AverageScore *int `json:"averageScore,omitempty"`
		
		HighScore *int `json:"highScore,omitempty"`
		
		LowScore *int `json:"lowScore,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		Evaluations *[]Evaluation `json:"evaluations,omitempty"`
		
		Evaluators *[]User `json:"evaluators,omitempty"`
		
		ScoringIndex *Evaluation `json:"scoringIndex,omitempty"`
		
		ExpertEvaluator *User `json:"expertEvaluator,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Calibrator: u.Calibrator,
		
		Agent: u.Agent,
		
		Conversation: u.Conversation,
		
		EvaluationForm: u.EvaluationForm,
		
		ContextId: u.ContextId,
		
		AverageScore: u.AverageScore,
		
		HighScore: u.HighScore,
		
		LowScore: u.LowScore,
		
		CreatedDate: CreatedDate,
		
		Evaluations: u.Evaluations,
		
		Evaluators: u.Evaluators,
		
		ScoringIndex: u.ScoringIndex,
		
		ExpertEvaluator: u.ExpertEvaluator,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Calibrationcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
