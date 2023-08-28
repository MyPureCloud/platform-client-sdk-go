package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Apiusagesimplesearch
type Apiusagesimplesearch struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Interval - Behaves like one clause in a SQL WHERE. Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`

	// Metrics - Behaves like a SQL SELECT clause. Enables retrieving only named metrics. If omitted, all metrics that are available will be returned (like SELECT *).
	Metrics *[]string `json:"metrics,omitempty"`

	// OauthClientNames - Behaves like a SQL WHERE with multiple IN operators. Specifies a list of OAuth client names to be queried.
	OauthClientNames *[]string `json:"oauthClientNames,omitempty"`

	// HttpMethods - Behaves like a SQL WHERE with multiple IN operators. Specifies a list of HTTP methods to be queried.
	HttpMethods *[]string `json:"httpMethods,omitempty"`

	// TemplateUris - Behaves like a SQL WHERE with multiple IN operators. Specifies a list of Template Uris to be queried.
	TemplateUris *[]string `json:"templateUris,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Apiusagesimplesearch) SetField(field string, fieldValue interface{}) {
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

func (o Apiusagesimplesearch) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Apiusagesimplesearch
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]string `json:"metrics,omitempty"`
		
		OauthClientNames *[]string `json:"oauthClientNames,omitempty"`
		
		HttpMethods *[]string `json:"httpMethods,omitempty"`
		
		TemplateUris *[]string `json:"templateUris,omitempty"`
		Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		
		OauthClientNames: o.OauthClientNames,
		
		HttpMethods: o.HttpMethods,
		
		TemplateUris: o.TemplateUris,
		Alias:    (Alias)(o),
	})
}

func (o *Apiusagesimplesearch) UnmarshalJSON(b []byte) error {
	var ApiusagesimplesearchMap map[string]interface{}
	err := json.Unmarshal(b, &ApiusagesimplesearchMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := ApiusagesimplesearchMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Metrics, ok := ApiusagesimplesearchMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if OauthClientNames, ok := ApiusagesimplesearchMap["oauthClientNames"].([]interface{}); ok {
		OauthClientNamesString, _ := json.Marshal(OauthClientNames)
		json.Unmarshal(OauthClientNamesString, &o.OauthClientNames)
	}
	
	if HttpMethods, ok := ApiusagesimplesearchMap["httpMethods"].([]interface{}); ok {
		HttpMethodsString, _ := json.Marshal(HttpMethods)
		json.Unmarshal(HttpMethodsString, &o.HttpMethods)
	}
	
	if TemplateUris, ok := ApiusagesimplesearchMap["templateUris"].([]interface{}); ok {
		TemplateUrisString, _ := json.Marshal(TemplateUris)
		json.Unmarshal(TemplateUrisString, &o.TemplateUris)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Apiusagesimplesearch) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
