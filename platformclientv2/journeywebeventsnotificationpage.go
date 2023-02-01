package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationpage
type Journeywebeventsnotificationpage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Url
	Url *string `json:"url,omitempty"`

	// Title
	Title *string `json:"title,omitempty"`

	// Domain
	Domain *string `json:"domain,omitempty"`

	// Fragment
	Fragment *string `json:"fragment,omitempty"`

	// Hostname
	Hostname *string `json:"hostname,omitempty"`

	// Keywords
	Keywords *string `json:"keywords,omitempty"`

	// Lang
	Lang *string `json:"lang,omitempty"`

	// Pathname
	Pathname *string `json:"pathname,omitempty"`

	// QueryString
	QueryString *string `json:"queryString,omitempty"`

	// Breadcrumb
	Breadcrumb *[]string `json:"breadcrumb,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeywebeventsnotificationpage) SetField(field string, fieldValue interface{}) {
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

func (o Journeywebeventsnotificationpage) MarshalJSON() ([]byte, error) {
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
	type Alias Journeywebeventsnotificationpage
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		Fragment *string `json:"fragment,omitempty"`
		
		Hostname *string `json:"hostname,omitempty"`
		
		Keywords *string `json:"keywords,omitempty"`
		
		Lang *string `json:"lang,omitempty"`
		
		Pathname *string `json:"pathname,omitempty"`
		
		QueryString *string `json:"queryString,omitempty"`
		
		Breadcrumb *[]string `json:"breadcrumb,omitempty"`
		Alias
	}{ 
		Url: o.Url,
		
		Title: o.Title,
		
		Domain: o.Domain,
		
		Fragment: o.Fragment,
		
		Hostname: o.Hostname,
		
		Keywords: o.Keywords,
		
		Lang: o.Lang,
		
		Pathname: o.Pathname,
		
		QueryString: o.QueryString,
		
		Breadcrumb: o.Breadcrumb,
		Alias:    (Alias)(o),
	})
}

func (o *Journeywebeventsnotificationpage) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationpageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationpageMap)
	if err != nil {
		return err
	}
	
	if Url, ok := JourneywebeventsnotificationpageMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Title, ok := JourneywebeventsnotificationpageMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Domain, ok := JourneywebeventsnotificationpageMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if Fragment, ok := JourneywebeventsnotificationpageMap["fragment"].(string); ok {
		o.Fragment = &Fragment
	}
    
	if Hostname, ok := JourneywebeventsnotificationpageMap["hostname"].(string); ok {
		o.Hostname = &Hostname
	}
    
	if Keywords, ok := JourneywebeventsnotificationpageMap["keywords"].(string); ok {
		o.Keywords = &Keywords
	}
    
	if Lang, ok := JourneywebeventsnotificationpageMap["lang"].(string); ok {
		o.Lang = &Lang
	}
    
	if Pathname, ok := JourneywebeventsnotificationpageMap["pathname"].(string); ok {
		o.Pathname = &Pathname
	}
    
	if QueryString, ok := JourneywebeventsnotificationpageMap["queryString"].(string); ok {
		o.QueryString = &QueryString
	}
    
	if Breadcrumb, ok := JourneywebeventsnotificationpageMap["breadcrumb"].([]interface{}); ok {
		BreadcrumbString, _ := json.Marshal(Breadcrumb)
		json.Unmarshal(BreadcrumbString, &o.Breadcrumb)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationpage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
