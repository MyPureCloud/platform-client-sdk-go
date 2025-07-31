package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Clientpublicapiusage
type Clientpublicapiusage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Date - The date of the usage. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	Date *time.Time `json:"date,omitempty"`

	// Platform - The platform the request(s) is/were made on.
	Platform *string `json:"platform,omitempty"`

	// HttpMethod - The http method of the request(s)
	HttpMethod *string `json:"httpMethod,omitempty"`

	// TemplateUri - The templateUri of the request(s).
	TemplateUri *string `json:"templateUri,omitempty"`

	// RequestCount - The total number of requests.
	RequestCount *int `json:"requestCount,omitempty"`

	// Status200 - The number of requests resulting in a 2xx HTTP status code.
	Status200 *int `json:"status200,omitempty"`

	// Status300 - The number of requests resulting in a 3xx HTTP status code.
	Status300 *int `json:"status300,omitempty"`

	// Status400 - The number of requests resulting in a 4xx HTTP status code.
	Status400 *int `json:"status400,omitempty"`

	// Status429 - The number of requests resulting in a 429 HTTP status code.
	Status429 *int `json:"status429,omitempty"`

	// Status500 - The number of requests resulting in a 5xx HTTP status code.
	Status500 *int `json:"status500,omitempty"`

	// OrganizationId - The organizationId that made the request.
	OrganizationId *string `json:"organizationId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Clientpublicapiusage) SetField(field string, fieldValue interface{}) {
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

func (o Clientpublicapiusage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "Date", }

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
	type Alias Clientpublicapiusage
	
	Date := new(string)
	if o.Date != nil {
		*Date = timeutil.Strftime(o.Date, "%Y-%m-%d")
	} else {
		Date = nil
	}
	
	return json.Marshal(&struct { 
		Date *string `json:"date,omitempty"`
		
		Platform *string `json:"platform,omitempty"`
		
		HttpMethod *string `json:"httpMethod,omitempty"`
		
		TemplateUri *string `json:"templateUri,omitempty"`
		
		RequestCount *int `json:"requestCount,omitempty"`
		
		Status200 *int `json:"status200,omitempty"`
		
		Status300 *int `json:"status300,omitempty"`
		
		Status400 *int `json:"status400,omitempty"`
		
		Status429 *int `json:"status429,omitempty"`
		
		Status500 *int `json:"status500,omitempty"`
		
		OrganizationId *string `json:"organizationId,omitempty"`
		Alias
	}{ 
		Date: Date,
		
		Platform: o.Platform,
		
		HttpMethod: o.HttpMethod,
		
		TemplateUri: o.TemplateUri,
		
		RequestCount: o.RequestCount,
		
		Status200: o.Status200,
		
		Status300: o.Status300,
		
		Status400: o.Status400,
		
		Status429: o.Status429,
		
		Status500: o.Status500,
		
		OrganizationId: o.OrganizationId,
		Alias:    (Alias)(o),
	})
}

func (o *Clientpublicapiusage) UnmarshalJSON(b []byte) error {
	var ClientpublicapiusageMap map[string]interface{}
	err := json.Unmarshal(b, &ClientpublicapiusageMap)
	if err != nil {
		return err
	}
	
	if dateString, ok := ClientpublicapiusageMap["date"].(string); ok {
		Date, _ := time.Parse("2006-01-02", dateString)
		o.Date = &Date
	}
	
	if Platform, ok := ClientpublicapiusageMap["platform"].(string); ok {
		o.Platform = &Platform
	}
    
	if HttpMethod, ok := ClientpublicapiusageMap["httpMethod"].(string); ok {
		o.HttpMethod = &HttpMethod
	}
    
	if TemplateUri, ok := ClientpublicapiusageMap["templateUri"].(string); ok {
		o.TemplateUri = &TemplateUri
	}
    
	if RequestCount, ok := ClientpublicapiusageMap["requestCount"].(float64); ok {
		RequestCountInt := int(RequestCount)
		o.RequestCount = &RequestCountInt
	}
	
	if Status200, ok := ClientpublicapiusageMap["status200"].(float64); ok {
		Status200Int := int(Status200)
		o.Status200 = &Status200Int
	}
	
	if Status300, ok := ClientpublicapiusageMap["status300"].(float64); ok {
		Status300Int := int(Status300)
		o.Status300 = &Status300Int
	}
	
	if Status400, ok := ClientpublicapiusageMap["status400"].(float64); ok {
		Status400Int := int(Status400)
		o.Status400 = &Status400Int
	}
	
	if Status429, ok := ClientpublicapiusageMap["status429"].(float64); ok {
		Status429Int := int(Status429)
		o.Status429 = &Status429Int
	}
	
	if Status500, ok := ClientpublicapiusageMap["status500"].(float64); ok {
		Status500Int := int(Status500)
		o.Status500 = &Status500Int
	}
	
	if OrganizationId, ok := ClientpublicapiusageMap["organizationId"].(string); ok {
		o.OrganizationId = &OrganizationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Clientpublicapiusage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
