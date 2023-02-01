package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Numberplan
type Numberplan struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`

	// Description - The resource's description.
	Description *string `json:"description,omitempty"`

	// Version - The current version of the resource.
	Version *int `json:"version,omitempty"`

	// DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// ModifiedBy - The ID of the user that last modified the resource.
	ModifiedBy *string `json:"modifiedBy,omitempty"`

	// CreatedBy - The ID of the user that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// State - Indicates if the resource is active, inactive, or deleted.
	State *string `json:"state,omitempty"`

	// ModifiedByApp - The application that last modified the resource.
	ModifiedByApp *string `json:"modifiedByApp,omitempty"`

	// CreatedByApp - The application that created the resource.
	CreatedByApp *string `json:"createdByApp,omitempty"`

	// Match
	Match *string `json:"match,omitempty"`

	// NormalizedFormat
	NormalizedFormat *string `json:"normalizedFormat,omitempty"`

	// Priority
	Priority *int `json:"priority,omitempty"`

	// Numbers
	Numbers *[]Number `json:"numbers,omitempty"`

	// DigitLength
	DigitLength *Digitlength `json:"digitLength,omitempty"`

	// Classification
	Classification *string `json:"classification,omitempty"`

	// MatchType
	MatchType *string `json:"matchType,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Numberplan) SetField(field string, fieldValue interface{}) {
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

func (o Numberplan) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Numberplan
	
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
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ModifiedByApp *string `json:"modifiedByApp,omitempty"`
		
		CreatedByApp *string `json:"createdByApp,omitempty"`
		
		Match *string `json:"match,omitempty"`
		
		NormalizedFormat *string `json:"normalizedFormat,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Numbers *[]Number `json:"numbers,omitempty"`
		
		DigitLength *Digitlength `json:"digitLength,omitempty"`
		
		Classification *string `json:"classification,omitempty"`
		
		MatchType *string `json:"matchType,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		CreatedBy: o.CreatedBy,
		
		State: o.State,
		
		ModifiedByApp: o.ModifiedByApp,
		
		CreatedByApp: o.CreatedByApp,
		
		Match: o.Match,
		
		NormalizedFormat: o.NormalizedFormat,
		
		Priority: o.Priority,
		
		Numbers: o.Numbers,
		
		DigitLength: o.DigitLength,
		
		Classification: o.Classification,
		
		MatchType: o.MatchType,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Numberplan) UnmarshalJSON(b []byte) error {
	var NumberplanMap map[string]interface{}
	err := json.Unmarshal(b, &NumberplanMap)
	if err != nil {
		return err
	}
	
	if Id, ok := NumberplanMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := NumberplanMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := NumberplanMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := NumberplanMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Version, ok := NumberplanMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := NumberplanMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := NumberplanMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := NumberplanMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    
	if CreatedBy, ok := NumberplanMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
    
	if State, ok := NumberplanMap["state"].(string); ok {
		o.State = &State
	}
    
	if ModifiedByApp, ok := NumberplanMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
    
	if CreatedByApp, ok := NumberplanMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
    
	if Match, ok := NumberplanMap["match"].(string); ok {
		o.Match = &Match
	}
    
	if NormalizedFormat, ok := NumberplanMap["normalizedFormat"].(string); ok {
		o.NormalizedFormat = &NormalizedFormat
	}
    
	if Priority, ok := NumberplanMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Numbers, ok := NumberplanMap["numbers"].([]interface{}); ok {
		NumbersString, _ := json.Marshal(Numbers)
		json.Unmarshal(NumbersString, &o.Numbers)
	}
	
	if DigitLength, ok := NumberplanMap["digitLength"].(map[string]interface{}); ok {
		DigitLengthString, _ := json.Marshal(DigitLength)
		json.Unmarshal(DigitLengthString, &o.DigitLength)
	}
	
	if Classification, ok := NumberplanMap["classification"].(string); ok {
		o.Classification = &Classification
	}
    
	if MatchType, ok := NumberplanMap["matchType"].(string); ok {
		o.MatchType = &MatchType
	}
    
	if SelfUri, ok := NumberplanMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Numberplan) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
