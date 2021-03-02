package platformclientv2
import (
	"encoding/json"
)

// Surveyassignment
type Surveyassignment struct { 
	// SurveyForm - The survey form used for this survey.
	SurveyForm *Publishedsurveyformreference `json:"surveyForm,omitempty"`


	// Flow - The URI reference to the flow associated with this survey.
	Flow *Domainentityref `json:"flow,omitempty"`


	// InviteTimeInterval - An ISO 8601 repeated interval consisting of the number of repetitions, the start datetime, and the interval (e.g. R2/2018-03-01T13:00:00Z/P1M10DT2H30M). Total duration must not exceed 90 days.
	InviteTimeInterval *string `json:"inviteTimeInterval,omitempty"`


	// SendingUser - User together with sendingDomain used to send email, null to use no-reply
	SendingUser *string `json:"sendingUser,omitempty"`


	// SendingDomain - Validated email domain, required
	SendingDomain *string `json:"sendingDomain,omitempty"`

}

// String returns a JSON representation of the model
func (o *Surveyassignment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
