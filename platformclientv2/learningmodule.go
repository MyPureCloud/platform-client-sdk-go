package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodule - Learning module response
type Learningmodule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of learning module
	Name *string `json:"name,omitempty"`


	// CreatedBy - The user who created learning module
	CreatedBy *Userreference `json:"createdBy,omitempty"`


	// DateCreated - The date/time learning module was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ModifiedBy - The user who modified learning module
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`


	// DateModified - The date/time learning module was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - The version of published learning module
	Version *int `json:"version,omitempty"`


	// ExternalId - The external ID of the learning module
	ExternalId *string `json:"externalId,omitempty"`


	// Source - The source of the learning module
	Source *string `json:"source,omitempty"`


	// Rule - The rule for learning module; read-only, and only populated when requested via expand param.
	Rule *Learningmodulerule `json:"rule,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// IsArchived - If true, learning module is archived
	IsArchived *bool `json:"isArchived,omitempty"`


	// IsPublished - If true, learning module is published
	IsPublished *bool `json:"isPublished,omitempty"`


	// Description - The description of learning module
	Description *string `json:"description,omitempty"`


	// CompletionTimeInDays - The completion time of learning module in days
	CompletionTimeInDays *int `json:"completionTimeInDays,omitempty"`


	// VarType - The type for the learning module
	VarType *string `json:"type,omitempty"`


	// InformSteps - The list of inform steps in a learning module
	InformSteps *[]Learningmoduleinformstep `json:"informSteps,omitempty"`


	// AssessmentForm - The assessment form for learning module
	AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`


	// SummaryData - The learning module summary data
	SummaryData *Learningmodulesummary `json:"summaryData,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningmodule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
