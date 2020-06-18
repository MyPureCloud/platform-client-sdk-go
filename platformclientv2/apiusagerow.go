package platformclientv2
import (
	"time"
	"encoding/json"
)

// Apiusagerow
type Apiusagerow struct { 
	// ClientId - Client Id associated with this query result
	ClientId *string `json:"clientId,omitempty"`


	// ClientName - Client Name associated with this query result
	ClientName *string `json:"clientName,omitempty"`


	// OrganizationId - Organization Id associated with this query result
	OrganizationId *string `json:"organizationId,omitempty"`


	// UserId - User Id associated with this query result
	UserId *string `json:"userId,omitempty"`


	// TemplateUri - Template Uri associated with this query result
	TemplateUri *string `json:"templateUri,omitempty"`


	// HttpMethod - HTTP Method associated with this query result
	HttpMethod *string `json:"httpMethod,omitempty"`


	// Status200 - Number of requests resulting in a 2xx HTTP status code
	Status200 *int64 `json:"status200,omitempty"`


	// Status300 - Number of requests resulting in a 3xx HTTP status code
	Status300 *int64 `json:"status300,omitempty"`


	// Status400 - Number of requests resulting in a 4xx HTTP status code
	Status400 *int64 `json:"status400,omitempty"`


	// Status500 - Number of requests resulting in a 5xx HTTP status code
	Status500 *int64 `json:"status500,omitempty"`


	// Status429 - Number of requests resulting in a 429 HTTP status code, this is a subset of the count returned with status400
	Status429 *int64 `json:"status429,omitempty"`


	// Requests - Total number of requests
	Requests *int64 `json:"requests,omitempty"`


	// Date - Date of requests, based on granularity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	Date *time.Time `json:"date,omitempty"`

}

// String returns a JSON representation of the model
func (o *Apiusagerow) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
