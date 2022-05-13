package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Surveyassignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyassignment
	
	return json.Marshal(&struct { 
		SurveyForm *Publishedsurveyformreference `json:"surveyForm,omitempty"`
		
		Flow *Domainentityref `json:"flow,omitempty"`
		
		InviteTimeInterval *string `json:"inviteTimeInterval,omitempty"`
		
		SendingUser *string `json:"sendingUser,omitempty"`
		
		SendingDomain *string `json:"sendingDomain,omitempty"`
		*Alias
	}{ 
		SurveyForm: o.SurveyForm,
		
		Flow: o.Flow,
		
		InviteTimeInterval: o.InviteTimeInterval,
		
		SendingUser: o.SendingUser,
		
		SendingDomain: o.SendingDomain,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveyassignment) UnmarshalJSON(b []byte) error {
	var SurveyassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyassignmentMap)
	if err != nil {
		return err
	}
	
	if SurveyForm, ok := SurveyassignmentMap["surveyForm"].(map[string]interface{}); ok {
		SurveyFormString, _ := json.Marshal(SurveyForm)
		json.Unmarshal(SurveyFormString, &o.SurveyForm)
	}
	
	if Flow, ok := SurveyassignmentMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if InviteTimeInterval, ok := SurveyassignmentMap["inviteTimeInterval"].(string); ok {
		o.InviteTimeInterval = &InviteTimeInterval
	}
    
	if SendingUser, ok := SurveyassignmentMap["sendingUser"].(string); ok {
		o.SendingUser = &SendingUser
	}
    
	if SendingDomain, ok := SurveyassignmentMap["sendingDomain"].(string); ok {
		o.SendingDomain = &SendingDomain
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
