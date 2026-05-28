package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryopportunityenrollmentmetadata
type Queryopportunityenrollmentmetadata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ModifiedBy - The user who last modified the associated entity. The id may be 'System' if it was an automated process
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// DateModified - The date the associated entity was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// CreatedBy - The user who created the associated entity, if available. The id may be 'System' if it was an automated process
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// DateCreated - The date the associated entity was created, if available. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// Version - The version of the associated entity.  Used to prevent conflicts on concurrent edits
	Version *int `json:"version,omitempty"`

	// ReviewedBy - The user who reviewed the enrollment
	ReviewedBy *Userreference `json:"reviewedBy,omitempty"`

	// DateReviewed - The date the enrollment was reviewed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateReviewed *time.Time `json:"dateReviewed,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queryopportunityenrollmentmetadata) SetField(field string, fieldValue interface{}) {
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

func (o Queryopportunityenrollmentmetadata) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateModified","DateCreated","DateReviewed", }
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
	type Alias Queryopportunityenrollmentmetadata
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateReviewed := new(string)
	if o.DateReviewed != nil {
		
		*DateReviewed = timeutil.Strftime(o.DateReviewed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateReviewed = nil
	}
	
	return json.Marshal(&struct { 
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		ReviewedBy *Userreference `json:"reviewedBy,omitempty"`
		
		DateReviewed *string `json:"dateReviewed,omitempty"`
		Alias
	}{ 
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		Version: o.Version,
		
		ReviewedBy: o.ReviewedBy,
		
		DateReviewed: DateReviewed,
		Alias:    (Alias)(o),
	})
}

func (o *Queryopportunityenrollmentmetadata) UnmarshalJSON(b []byte) error {
	var QueryopportunityenrollmentmetadataMap map[string]interface{}
	err := json.Unmarshal(b, &QueryopportunityenrollmentmetadataMap)
	if err != nil {
		return err
	}
	
	if ModifiedBy, ok := QueryopportunityenrollmentmetadataMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := QueryopportunityenrollmentmetadataMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := QueryopportunityenrollmentmetadataMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := QueryopportunityenrollmentmetadataMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if Version, ok := QueryopportunityenrollmentmetadataMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ReviewedBy, ok := QueryopportunityenrollmentmetadataMap["reviewedBy"].(map[string]interface{}); ok {
		ReviewedByString, _ := json.Marshal(ReviewedBy)
		json.Unmarshal(ReviewedByString, &o.ReviewedBy)
	}
	
	if dateReviewedString, ok := QueryopportunityenrollmentmetadataMap["dateReviewed"].(string); ok {
		DateReviewed, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateReviewedString)
		o.DateReviewed = &DateReviewed
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queryopportunityenrollmentmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
