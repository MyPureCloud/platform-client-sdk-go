package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainedgesoftwareversiondto
type Domainedgesoftwareversiondto struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// EdgeVersion
	EdgeVersion *string `json:"edgeVersion,omitempty"`

	// PublishDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PublishDate *time.Time `json:"publishDate,omitempty"`

	// EdgeUri
	EdgeUri *string `json:"edgeUri,omitempty"`

	// Current
	Current *bool `json:"current,omitempty"`

	// LatestRelease
	LatestRelease *bool `json:"latestRelease,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Domainedgesoftwareversiondto) SetField(field string, fieldValue interface{}) {
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

func (o Domainedgesoftwareversiondto) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "PublishDate", }
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
	type Alias Domainedgesoftwareversiondto
	
	PublishDate := new(string)
	if o.PublishDate != nil {
		
		*PublishDate = timeutil.Strftime(o.PublishDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PublishDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		EdgeVersion *string `json:"edgeVersion,omitempty"`
		
		PublishDate *string `json:"publishDate,omitempty"`
		
		EdgeUri *string `json:"edgeUri,omitempty"`
		
		Current *bool `json:"current,omitempty"`
		
		LatestRelease *bool `json:"latestRelease,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		EdgeVersion: o.EdgeVersion,
		
		PublishDate: PublishDate,
		
		EdgeUri: o.EdgeUri,
		
		Current: o.Current,
		
		LatestRelease: o.LatestRelease,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Domainedgesoftwareversiondto) UnmarshalJSON(b []byte) error {
	var DomainedgesoftwareversiondtoMap map[string]interface{}
	err := json.Unmarshal(b, &DomainedgesoftwareversiondtoMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DomainedgesoftwareversiondtoMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DomainedgesoftwareversiondtoMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if EdgeVersion, ok := DomainedgesoftwareversiondtoMap["edgeVersion"].(string); ok {
		o.EdgeVersion = &EdgeVersion
	}
    
	if publishDateString, ok := DomainedgesoftwareversiondtoMap["publishDate"].(string); ok {
		PublishDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", publishDateString)
		o.PublishDate = &PublishDate
	}
	
	if EdgeUri, ok := DomainedgesoftwareversiondtoMap["edgeUri"].(string); ok {
		o.EdgeUri = &EdgeUri
	}
    
	if Current, ok := DomainedgesoftwareversiondtoMap["current"].(bool); ok {
		o.Current = &Current
	}
    
	if LatestRelease, ok := DomainedgesoftwareversiondtoMap["latestRelease"].(bool); ok {
		o.LatestRelease = &LatestRelease
	}
    
	if SelfUri, ok := DomainedgesoftwareversiondtoMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Domainedgesoftwareversiondto) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
