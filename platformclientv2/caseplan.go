package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Caseplan
type Caseplan struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the Caseplan.
	Name *string `json:"name,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Starrabledivision `json:"division,omitempty"`

	// Description - The description of the Caseplan.
	Description *string `json:"description,omitempty"`

	// ReferencePrefix - The prefix used when creating the reference for Cases from the Caseplan.
	ReferencePrefix *string `json:"referencePrefix,omitempty"`

	// DefaultDueDurationInSeconds - The default due duration in seconds for Cases created from the Caseplan.
	DefaultDueDurationInSeconds *int `json:"defaultDueDurationInSeconds,omitempty"`

	// DefaultTtlSeconds - The default TTL in seconds for Cases created from the Caseplan.
	DefaultTtlSeconds *int `json:"defaultTtlSeconds,omitempty"`

	// DefaultCaseOwner - The default case owner for Cases created from the Caseplan.
	DefaultCaseOwner *Userreference `json:"defaultCaseOwner,omitempty"`

	// Latest - The latest version of the Caseplan.
	Latest *int `json:"latest,omitempty"`

	// Published - The published version of the Caseplan.
	Published *int `json:"published,omitempty"`

	// DateCreated - The Caseplan creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The Caseplan modification date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DatePublished - The Caseplan publication date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePublished *time.Time `json:"datePublished,omitempty"`

	// ModifiedBy - The id of the User who modified the Caseplan.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// CustomerIntent - The customer intent for the Cases created from the caseplan.
	CustomerIntent *Customerintentreference `json:"customerIntent,omitempty"`

	// VersionState - The version state of the Caseplan.
	VersionState *string `json:"versionState,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Caseplan) SetField(field string, fieldValue interface{}) {
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

func (o Caseplan) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DatePublished", }
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
	type Alias Caseplan
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DatePublished := new(string)
	if o.DatePublished != nil {
		
		*DatePublished = timeutil.Strftime(o.DatePublished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DatePublished = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Starrabledivision `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ReferencePrefix *string `json:"referencePrefix,omitempty"`
		
		DefaultDueDurationInSeconds *int `json:"defaultDueDurationInSeconds,omitempty"`
		
		DefaultTtlSeconds *int `json:"defaultTtlSeconds,omitempty"`
		
		DefaultCaseOwner *Userreference `json:"defaultCaseOwner,omitempty"`
		
		Latest *int `json:"latest,omitempty"`
		
		Published *int `json:"published,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DatePublished *string `json:"datePublished,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		CustomerIntent *Customerintentreference `json:"customerIntent,omitempty"`
		
		VersionState *string `json:"versionState,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		ReferencePrefix: o.ReferencePrefix,
		
		DefaultDueDurationInSeconds: o.DefaultDueDurationInSeconds,
		
		DefaultTtlSeconds: o.DefaultTtlSeconds,
		
		DefaultCaseOwner: o.DefaultCaseOwner,
		
		Latest: o.Latest,
		
		Published: o.Published,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DatePublished: DatePublished,
		
		ModifiedBy: o.ModifiedBy,
		
		CustomerIntent: o.CustomerIntent,
		
		VersionState: o.VersionState,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Caseplan) UnmarshalJSON(b []byte) error {
	var CaseplanMap map[string]interface{}
	err := json.Unmarshal(b, &CaseplanMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CaseplanMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CaseplanMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := CaseplanMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := CaseplanMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ReferencePrefix, ok := CaseplanMap["referencePrefix"].(string); ok {
		o.ReferencePrefix = &ReferencePrefix
	}
    
	if DefaultDueDurationInSeconds, ok := CaseplanMap["defaultDueDurationInSeconds"].(float64); ok {
		DefaultDueDurationInSecondsInt := int(DefaultDueDurationInSeconds)
		o.DefaultDueDurationInSeconds = &DefaultDueDurationInSecondsInt
	}
	
	if DefaultTtlSeconds, ok := CaseplanMap["defaultTtlSeconds"].(float64); ok {
		DefaultTtlSecondsInt := int(DefaultTtlSeconds)
		o.DefaultTtlSeconds = &DefaultTtlSecondsInt
	}
	
	if DefaultCaseOwner, ok := CaseplanMap["defaultCaseOwner"].(map[string]interface{}); ok {
		DefaultCaseOwnerString, _ := json.Marshal(DefaultCaseOwner)
		json.Unmarshal(DefaultCaseOwnerString, &o.DefaultCaseOwner)
	}
	
	if Latest, ok := CaseplanMap["latest"].(float64); ok {
		LatestInt := int(Latest)
		o.Latest = &LatestInt
	}
	
	if Published, ok := CaseplanMap["published"].(float64); ok {
		PublishedInt := int(Published)
		o.Published = &PublishedInt
	}
	
	if dateCreatedString, ok := CaseplanMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := CaseplanMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if datePublishedString, ok := CaseplanMap["datePublished"].(string); ok {
		DatePublished, _ := time.Parse("2006-01-02T15:04:05.999999Z", datePublishedString)
		o.DatePublished = &DatePublished
	}
	
	if ModifiedBy, ok := CaseplanMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if CustomerIntent, ok := CaseplanMap["customerIntent"].(map[string]interface{}); ok {
		CustomerIntentString, _ := json.Marshal(CustomerIntent)
		json.Unmarshal(CustomerIntentString, &o.CustomerIntent)
	}
	
	if VersionState, ok := CaseplanMap["versionState"].(string); ok {
		o.VersionState = &VersionState
	}
    
	if SelfUri, ok := CaseplanMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Caseplan) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
