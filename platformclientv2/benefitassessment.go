package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Benefitassessment
type Benefitassessment struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Queues - The list of queues that are assessed for Predictive Routing benefit.
	Queues *[]Addressableentityref `json:"queues,omitempty"`


	// KpiAssessments - A set of key performance indicators applied on the queue to determine suitability of Predictive Routing.
	KpiAssessments *[]Keyperformanceindicatorassessment `json:"kpiAssessments,omitempty"`


	// State - State of the benefit assessment.
	State *string `json:"state,omitempty"`


	// DateCreated - Creation Date of the benefit assessment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Modified Date of the benefit assessment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Benefitassessment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
