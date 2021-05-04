package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludomainversion
type Nludomainversion struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Domain - The NLU domain of the version.
	Domain **Nludomain `json:"domain,omitempty"`


	// Description - The description of the NLU domain version.
	Description *string `json:"description,omitempty"`


	// Language - The language that the NLU domain version supports.
	Language *string `json:"language,omitempty"`


	// Published - Whether this NLU domain version has been published.
	Published *bool `json:"published,omitempty"`


	// DateCreated - The date when the NLU domain version was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date when the NLU domain version was updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// DateTrained - The date when the NLU domain version was trained. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateTrained *time.Time `json:"dateTrained,omitempty"`


	// DatePublished - The date when the NLU domain version was published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePublished *time.Time `json:"datePublished,omitempty"`


	// TrainingStatus - The training status of the NLU domain version.
	TrainingStatus *string `json:"trainingStatus,omitempty"`


	// EvaluationStatus - The evaluation status of the NLU domain version.
	EvaluationStatus *string `json:"evaluationStatus,omitempty"`


	// Intents - The intents defined for this NLU domain version.
	Intents *[]Intentdefinition `json:"intents,omitempty"`


	// EntityTypes - The entity types defined for this NLU domain version.
	EntityTypes *[]Namedentitytypedefinition `json:"entityTypes,omitempty"`


	// Entities - The entities defined for this NLU domain version.This field is mutually exclusive with entityTypeBindings
	Entities *[]Namedentitydefinition `json:"entities,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nludomainversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
