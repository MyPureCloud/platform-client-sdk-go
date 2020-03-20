package platformclientv2
import (
	"time"
	"encoding/json"
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


	// CreatedDate - When the request was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
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

// String returns a JSON representation of the model
func (o *Gdprrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
