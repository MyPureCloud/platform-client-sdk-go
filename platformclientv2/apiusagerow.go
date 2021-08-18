package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	Status200 *int `json:"status200,omitempty"`


	// Status300 - Number of requests resulting in a 3xx HTTP status code
	Status300 *int `json:"status300,omitempty"`


	// Status400 - Number of requests resulting in a 4xx HTTP status code
	Status400 *int `json:"status400,omitempty"`


	// Status500 - Number of requests resulting in a 5xx HTTP status code
	Status500 *int `json:"status500,omitempty"`


	// Status429 - Number of requests resulting in a 429 HTTP status code, this is a subset of the count returned with status400
	Status429 *int `json:"status429,omitempty"`


	// Requests - Total number of requests
	Requests *int `json:"requests,omitempty"`


	// Date - Date of requests, based on granularity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Date *time.Time `json:"date,omitempty"`

}

func (u *Apiusagerow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Apiusagerow

	
	Date := new(string)
	if u.Date != nil {
		
		*Date = timeutil.Strftime(u.Date, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Date = nil
	}
	

	return json.Marshal(&struct { 
		ClientId *string `json:"clientId,omitempty"`
		
		ClientName *string `json:"clientName,omitempty"`
		
		OrganizationId *string `json:"organizationId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		TemplateUri *string `json:"templateUri,omitempty"`
		
		HttpMethod *string `json:"httpMethod,omitempty"`
		
		Status200 *int `json:"status200,omitempty"`
		
		Status300 *int `json:"status300,omitempty"`
		
		Status400 *int `json:"status400,omitempty"`
		
		Status500 *int `json:"status500,omitempty"`
		
		Status429 *int `json:"status429,omitempty"`
		
		Requests *int `json:"requests,omitempty"`
		
		Date *string `json:"date,omitempty"`
		*Alias
	}{ 
		ClientId: u.ClientId,
		
		ClientName: u.ClientName,
		
		OrganizationId: u.OrganizationId,
		
		UserId: u.UserId,
		
		TemplateUri: u.TemplateUri,
		
		HttpMethod: u.HttpMethod,
		
		Status200: u.Status200,
		
		Status300: u.Status300,
		
		Status400: u.Status400,
		
		Status500: u.Status500,
		
		Status429: u.Status429,
		
		Requests: u.Requests,
		
		Date: Date,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Apiusagerow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
