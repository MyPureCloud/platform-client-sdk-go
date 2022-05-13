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

func (o *Apiusagerow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Apiusagerow
	
	Date := new(string)
	if o.Date != nil {
		
		*Date = timeutil.Strftime(o.Date, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		ClientId: o.ClientId,
		
		ClientName: o.ClientName,
		
		OrganizationId: o.OrganizationId,
		
		UserId: o.UserId,
		
		TemplateUri: o.TemplateUri,
		
		HttpMethod: o.HttpMethod,
		
		Status200: o.Status200,
		
		Status300: o.Status300,
		
		Status400: o.Status400,
		
		Status500: o.Status500,
		
		Status429: o.Status429,
		
		Requests: o.Requests,
		
		Date: Date,
		Alias:    (*Alias)(o),
	})
}

func (o *Apiusagerow) UnmarshalJSON(b []byte) error {
	var ApiusagerowMap map[string]interface{}
	err := json.Unmarshal(b, &ApiusagerowMap)
	if err != nil {
		return err
	}
	
	if ClientId, ok := ApiusagerowMap["clientId"].(string); ok {
		o.ClientId = &ClientId
	}
    
	if ClientName, ok := ApiusagerowMap["clientName"].(string); ok {
		o.ClientName = &ClientName
	}
    
	if OrganizationId, ok := ApiusagerowMap["organizationId"].(string); ok {
		o.OrganizationId = &OrganizationId
	}
    
	if UserId, ok := ApiusagerowMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if TemplateUri, ok := ApiusagerowMap["templateUri"].(string); ok {
		o.TemplateUri = &TemplateUri
	}
    
	if HttpMethod, ok := ApiusagerowMap["httpMethod"].(string); ok {
		o.HttpMethod = &HttpMethod
	}
    
	if Status200, ok := ApiusagerowMap["status200"].(float64); ok {
		Status200Int := int(Status200)
		o.Status200 = &Status200Int
	}
	
	if Status300, ok := ApiusagerowMap["status300"].(float64); ok {
		Status300Int := int(Status300)
		o.Status300 = &Status300Int
	}
	
	if Status400, ok := ApiusagerowMap["status400"].(float64); ok {
		Status400Int := int(Status400)
		o.Status400 = &Status400Int
	}
	
	if Status500, ok := ApiusagerowMap["status500"].(float64); ok {
		Status500Int := int(Status500)
		o.Status500 = &Status500Int
	}
	
	if Status429, ok := ApiusagerowMap["status429"].(float64); ok {
		Status429Int := int(Status429)
		o.Status429 = &Status429Int
	}
	
	if Requests, ok := ApiusagerowMap["requests"].(float64); ok {
		RequestsInt := int(Requests)
		o.Requests = &RequestsInt
	}
	
	if dateString, ok := ApiusagerowMap["date"].(string); ok {
		Date, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateString)
		o.Date = &Date
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Apiusagerow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
