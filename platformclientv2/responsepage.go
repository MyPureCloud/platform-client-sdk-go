package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Responsepage
type Responsepage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Url - The page URL.
	Url *string `json:"url,omitempty"`

	// Title - Title of the page.
	Title *string `json:"title,omitempty"`

	// Domain - Domain of the page's URL.
	Domain *string `json:"domain,omitempty"`

	// Fragment - Fragment or hash of the page's URL.
	Fragment *string `json:"fragment,omitempty"`

	// Hostname - Hostname of the page's URL.
	Hostname *string `json:"hostname,omitempty"`

	// Keywords - Keywords from the HTML <meta> tag of the page.
	Keywords *string `json:"keywords,omitempty"`

	// Lang - ISO 639-1 language code for the page as defined in the <html> tag.
	Lang *string `json:"lang,omitempty"`

	// Pathname - Path name of the page for the event.
	Pathname *string `json:"pathname,omitempty"`

	// QueryString - Query string that is passed to the page in the current event.
	QueryString *string `json:"queryString,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Responsepage) SetField(field string, fieldValue interface{}) {
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

func (o Responsepage) MarshalJSON() ([]byte, error) {
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
	type Alias Responsepage
	
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
		Alias:    (Alias)(o),
	})
}

func (o *Responsepage) UnmarshalJSON(b []byte) error {
	var ResponsepageMap map[string]interface{}
	err := json.Unmarshal(b, &ResponsepageMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ResponsepageMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Title, ok := ResponsepageMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Domain, ok := ResponsepageMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if Fragment, ok := ResponsepageMap["fragment"].(string); ok {
		o.Fragment = &Fragment
	}
    
	if Hostname, ok := ResponsepageMap["hostname"].(string); ok {
		o.Hostname = &Hostname
	}
    
	if Keywords, ok := ResponsepageMap["keywords"].(string); ok {
		o.Keywords = &Keywords
	}
    
	if Lang, ok := ResponsepageMap["lang"].(string); ok {
		o.Lang = &Lang
	}
    
	if Pathname, ok := ResponsepageMap["pathname"].(string); ok {
		o.Pathname = &Pathname
	}
    
	if QueryString, ok := ResponsepageMap["queryString"].(string); ok {
		o.QueryString = &QueryString
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Responsepage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
