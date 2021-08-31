package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (o *Benefitassessment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Benefitassessment
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Queues *[]Addressableentityref `json:"queues,omitempty"`
		
		KpiAssessments *[]Keyperformanceindicatorassessment `json:"kpiAssessments,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Queues: o.Queues,
		
		KpiAssessments: o.KpiAssessments,
		
		State: o.State,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Benefitassessment) UnmarshalJSON(b []byte) error {
	var BenefitassessmentMap map[string]interface{}
	err := json.Unmarshal(b, &BenefitassessmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BenefitassessmentMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Queues, ok := BenefitassessmentMap["queues"].([]interface{}); ok {
		QueuesString, _ := json.Marshal(Queues)
		json.Unmarshal(QueuesString, &o.Queues)
	}
	
	if KpiAssessments, ok := BenefitassessmentMap["kpiAssessments"].([]interface{}); ok {
		KpiAssessmentsString, _ := json.Marshal(KpiAssessments)
		json.Unmarshal(KpiAssessmentsString, &o.KpiAssessments)
	}
	
	if State, ok := BenefitassessmentMap["state"].(string); ok {
		o.State = &State
	}
	
	if dateCreatedString, ok := BenefitassessmentMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := BenefitassessmentMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := BenefitassessmentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Benefitassessment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
