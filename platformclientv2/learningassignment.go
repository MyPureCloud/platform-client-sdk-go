package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignment - Learning module assignment with user information
type Learningassignment struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Assessment - The assessment associated with this assignment
	Assessment *Learningassessment `json:"assessment,omitempty"`


	// CreatedBy - The user who created the assignment
	CreatedBy *Userreference `json:"createdBy,omitempty"`


	// DateCreated - The date when the assignment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ModifiedBy - The user who modified the assignment
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`


	// DateModified - The date when the assignment was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// IsOverdue - True if the assignment is overdue
	IsOverdue *bool `json:"isOverdue,omitempty"`


	// PercentageScore - The user's percentage score for this assignment
	PercentageScore *float32 `json:"percentageScore,omitempty"`


	// IsRule - True if this assignment was created by a Rule
	IsRule *bool `json:"isRule,omitempty"`


	// IsManual - True if this assignment was created manually
	IsManual *bool `json:"isManual,omitempty"`


	// IsPassed - True if the assessment was passed
	IsPassed *bool `json:"isPassed,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// State - The Learning Assignment state
	State *string `json:"state,omitempty"`


	// DateRecommendedForCompletion - The recommended completion date of the assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateRecommendedForCompletion *time.Time `json:"dateRecommendedForCompletion,omitempty"`


	// Version - The version of Learning module assigned
	Version *int `json:"version,omitempty"`


	// Module - The Learning module object associated with this assignment
	Module *Learningmodule `json:"module,omitempty"`


	// User - The user to whom the assignment is assigned
	User *Userreference `json:"user,omitempty"`


	// AssessmentForm - The assessment form associated with this assignment
	AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`

}

func (u *Learningassignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignment

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateRecommendedForCompletion := new(string)
	if u.DateRecommendedForCompletion != nil {
		
		*DateRecommendedForCompletion = timeutil.Strftime(u.DateRecommendedForCompletion, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateRecommendedForCompletion = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Assessment *Learningassessment `json:"assessment,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		IsOverdue *bool `json:"isOverdue,omitempty"`
		
		PercentageScore *float32 `json:"percentageScore,omitempty"`
		
		IsRule *bool `json:"isRule,omitempty"`
		
		IsManual *bool `json:"isManual,omitempty"`
		
		IsPassed *bool `json:"isPassed,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateRecommendedForCompletion *string `json:"dateRecommendedForCompletion,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Module *Learningmodule `json:"module,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Assessment: u.Assessment,
		
		CreatedBy: u.CreatedBy,
		
		DateCreated: DateCreated,
		
		ModifiedBy: u.ModifiedBy,
		
		DateModified: DateModified,
		
		IsOverdue: u.IsOverdue,
		
		PercentageScore: u.PercentageScore,
		
		IsRule: u.IsRule,
		
		IsManual: u.IsManual,
		
		IsPassed: u.IsPassed,
		
		SelfUri: u.SelfUri,
		
		State: u.State,
		
		DateRecommendedForCompletion: DateRecommendedForCompletion,
		
		Version: u.Version,
		
		Module: u.Module,
		
		User: u.User,
		
		AssessmentForm: u.AssessmentForm,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
