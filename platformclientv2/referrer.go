package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Referrer
type Referrer struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Url - Referrer URL.
	Url *string `json:"url,omitempty"`

	// Domain - Referrer URL domain.
	Domain *string `json:"domain,omitempty"`

	// Hostname - Referrer URL hostname.
	Hostname *string `json:"hostname,omitempty"`

	// Keywords - Referrer keywords.
	Keywords *string `json:"keywords,omitempty"`

	// Pathname - Referrer URL pathname.
	Pathname *string `json:"pathname,omitempty"`

	// QueryString - Referrer URL querystring.
	QueryString *string `json:"queryString,omitempty"`

	// Fragment - Referrer URL fragment.
	Fragment *string `json:"fragment,omitempty"`

	// Name - Name of referrer (e.g. Yahoo!, Google, InfoSpace).
	Name *string `json:"name,omitempty"`

	// Medium - Type of referrer (e.g. search, social).
	Medium *string `json:"medium,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Referrer) SetField(field string, fieldValue interface{}) {
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

func (o Referrer) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Referrer
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		Hostname *string `json:"hostname,omitempty"`
		
		Keywords *string `json:"keywords,omitempty"`
		
		Pathname *string `json:"pathname,omitempty"`
		
		QueryString *string `json:"queryString,omitempty"`
		
		Fragment *string `json:"fragment,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Medium *string `json:"medium,omitempty"`
		Alias
	}{ 
		Url: o.Url,
		
		Domain: o.Domain,
		
		Hostname: o.Hostname,
		
		Keywords: o.Keywords,
		
		Pathname: o.Pathname,
		
		QueryString: o.QueryString,
		
		Fragment: o.Fragment,
		
		Name: o.Name,
		
		Medium: o.Medium,
		Alias:    (Alias)(o),
	})
}

func (o *Referrer) UnmarshalJSON(b []byte) error {
	var ReferrerMap map[string]interface{}
	err := json.Unmarshal(b, &ReferrerMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ReferrerMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Domain, ok := ReferrerMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if Hostname, ok := ReferrerMap["hostname"].(string); ok {
		o.Hostname = &Hostname
	}
    
	if Keywords, ok := ReferrerMap["keywords"].(string); ok {
		o.Keywords = &Keywords
	}
    
	if Pathname, ok := ReferrerMap["pathname"].(string); ok {
		o.Pathname = &Pathname
	}
    
	if QueryString, ok := ReferrerMap["queryString"].(string); ok {
		o.QueryString = &QueryString
	}
    
	if Fragment, ok := ReferrerMap["fragment"].(string); ok {
		o.Fragment = &Fragment
	}
    
	if Name, ok := ReferrerMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Medium, ok := ReferrerMap["medium"].(string); ok {
		o.Medium = &Medium
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Referrer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
