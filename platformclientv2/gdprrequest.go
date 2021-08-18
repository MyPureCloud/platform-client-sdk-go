package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Gdprrequest
type Gdprrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// CreatedBy - The user that created this request
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ReplacementTerms - The replacement terms for the provided search terms, in the case of a GDPR_UPDATE request
	ReplacementTerms *[]Replacementterm `json:"replacementTerms,omitempty"`


	// RequestType - The type of GDPR request
	RequestType *string `json:"requestType,omitempty"`


	// CreatedDate - When the request was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// Status - The status of the request
	Status *string `json:"status,omitempty"`


	// Subject - The subject of the GDPR request
	Subject *Gdprsubject `json:"subject,omitempty"`


	// ResultsUrl - The location where the results of the request can be retrieved
	ResultsUrl *string `json:"resultsUrl,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Gdprrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Gdprrequest

	
	CreatedDate := new(string)
	if u.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(u.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		ReplacementTerms *[]Replacementterm `json:"replacementTerms,omitempty"`
		
		RequestType *string `json:"requestType,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Subject *Gdprsubject `json:"subject,omitempty"`
		
		ResultsUrl *string `json:"resultsUrl,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		CreatedBy: u.CreatedBy,
		
		ReplacementTerms: u.ReplacementTerms,
		
		RequestType: u.RequestType,
		
		CreatedDate: CreatedDate,
		
		Status: u.Status,
		
		Subject: u.Subject,
		
		ResultsUrl: u.ResultsUrl,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Gdprrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
