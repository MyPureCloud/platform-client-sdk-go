package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Apiusagerow
type Apiusagerow struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Apiusagerow) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Apiusagerow) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Date", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		Alias
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
		Alias:    (Alias)(o),
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
