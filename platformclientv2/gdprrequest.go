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

func (o *Gdprrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Gdprrequest
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		CreatedBy: o.CreatedBy,
		
		ReplacementTerms: o.ReplacementTerms,
		
		RequestType: o.RequestType,
		
		CreatedDate: CreatedDate,
		
		Status: o.Status,
		
		Subject: o.Subject,
		
		ResultsUrl: o.ResultsUrl,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Gdprrequest) UnmarshalJSON(b []byte) error {
	var GdprrequestMap map[string]interface{}
	err := json.Unmarshal(b, &GdprrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GdprrequestMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := GdprrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if CreatedBy, ok := GdprrequestMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ReplacementTerms, ok := GdprrequestMap["replacementTerms"].([]interface{}); ok {
		ReplacementTermsString, _ := json.Marshal(ReplacementTerms)
		json.Unmarshal(ReplacementTermsString, &o.ReplacementTerms)
	}
	
	if RequestType, ok := GdprrequestMap["requestType"].(string); ok {
		o.RequestType = &RequestType
	}
	
	if createdDateString, ok := GdprrequestMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if Status, ok := GdprrequestMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if Subject, ok := GdprrequestMap["subject"].(map[string]interface{}); ok {
		SubjectString, _ := json.Marshal(Subject)
		json.Unmarshal(SubjectString, &o.Subject)
	}
	
	if ResultsUrl, ok := GdprrequestMap["resultsUrl"].(string); ok {
		o.ResultsUrl = &ResultsUrl
	}
	
	if SelfUri, ok := GdprrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Gdprrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
