package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Organizationpublicapiusagequeryrequest
type Organizationpublicapiusagequeryrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Interval - Specify the interval to query on. Start and end are inclusive. Start date cannot be more than a year ago. End date cannot be more than 90 days after the start. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`

	// Granularity - Specify the granularity to aggregate the data to.
	Granularity *string `json:"granularity,omitempty"`

	// SortBy - Specify how to sort the returned data.
	SortBy *[]Usagequerysortby `json:"sortBy,omitempty"`

	// Metrics - Specify which metrics you want returned (all will be returned by default).
	Metrics *[]string `json:"metrics,omitempty"`

	// TemplateUris - Specify if you only want a subset of templateUris represented in the query.
	TemplateUris *[]string `json:"templateUris,omitempty"`

	// HttpMethods - Specify if you only want a subset of httpMethods represented in the query.
	HttpMethods *[]string `json:"httpMethods,omitempty"`

	// Platforms - Specify if you only want a subset of platforms represented in the query.
	Platforms *[]string `json:"platforms,omitempty"`

	// GroupBy - Specify how to aggregate the data (by default the data is not aggregated).
	GroupBy *[]string `json:"groupBy,omitempty"`

	// UserIds - Specify if you only want a subset of userIds represented in the query.
	UserIds *[]string `json:"userIds,omitempty"`

	// OauthClientIds - Specify if you only want a subset of oauthClientIds represented in the query.
	OauthClientIds *[]string `json:"oauthClientIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Organizationpublicapiusagequeryrequest) SetField(field string, fieldValue interface{}) {
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

func (o Organizationpublicapiusagequeryrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Organizationpublicapiusagequeryrequest
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		SortBy *[]Usagequerysortby `json:"sortBy,omitempty"`
		
		Metrics *[]string `json:"metrics,omitempty"`
		
		TemplateUris *[]string `json:"templateUris,omitempty"`
		
		HttpMethods *[]string `json:"httpMethods,omitempty"`
		
		Platforms *[]string `json:"platforms,omitempty"`
		
		GroupBy *[]string `json:"groupBy,omitempty"`
		
		UserIds *[]string `json:"userIds,omitempty"`
		
		OauthClientIds *[]string `json:"oauthClientIds,omitempty"`
		Alias
	}{ 
		Interval: o.Interval,
		
		Granularity: o.Granularity,
		
		SortBy: o.SortBy,
		
		Metrics: o.Metrics,
		
		TemplateUris: o.TemplateUris,
		
		HttpMethods: o.HttpMethods,
		
		Platforms: o.Platforms,
		
		GroupBy: o.GroupBy,
		
		UserIds: o.UserIds,
		
		OauthClientIds: o.OauthClientIds,
		Alias:    (Alias)(o),
	})
}

func (o *Organizationpublicapiusagequeryrequest) UnmarshalJSON(b []byte) error {
	var OrganizationpublicapiusagequeryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &OrganizationpublicapiusagequeryrequestMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := OrganizationpublicapiusagequeryrequestMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if Granularity, ok := OrganizationpublicapiusagequeryrequestMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if SortBy, ok := OrganizationpublicapiusagequeryrequestMap["sortBy"].([]interface{}); ok {
		SortByString, _ := json.Marshal(SortBy)
		json.Unmarshal(SortByString, &o.SortBy)
	}
	
	if Metrics, ok := OrganizationpublicapiusagequeryrequestMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if TemplateUris, ok := OrganizationpublicapiusagequeryrequestMap["templateUris"].([]interface{}); ok {
		TemplateUrisString, _ := json.Marshal(TemplateUris)
		json.Unmarshal(TemplateUrisString, &o.TemplateUris)
	}
	
	if HttpMethods, ok := OrganizationpublicapiusagequeryrequestMap["httpMethods"].([]interface{}); ok {
		HttpMethodsString, _ := json.Marshal(HttpMethods)
		json.Unmarshal(HttpMethodsString, &o.HttpMethods)
	}
	
	if Platforms, ok := OrganizationpublicapiusagequeryrequestMap["platforms"].([]interface{}); ok {
		PlatformsString, _ := json.Marshal(Platforms)
		json.Unmarshal(PlatformsString, &o.Platforms)
	}
	
	if GroupBy, ok := OrganizationpublicapiusagequeryrequestMap["groupBy"].([]interface{}); ok {
		GroupByString, _ := json.Marshal(GroupBy)
		json.Unmarshal(GroupByString, &o.GroupBy)
	}
	
	if UserIds, ok := OrganizationpublicapiusagequeryrequestMap["userIds"].([]interface{}); ok {
		UserIdsString, _ := json.Marshal(UserIds)
		json.Unmarshal(UserIdsString, &o.UserIds)
	}
	
	if OauthClientIds, ok := OrganizationpublicapiusagequeryrequestMap["oauthClientIds"].([]interface{}); ok {
		OauthClientIdsString, _ := json.Marshal(OauthClientIds)
		json.Unmarshal(OauthClientIdsString, &o.OauthClientIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Organizationpublicapiusagequeryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
