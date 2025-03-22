package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Socialtopicresponse
type Socialtopicresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - ID of the social topic.
	Id *string `json:"id,omitempty"`

	// Name - The name of the social topic.
	Name *string `json:"name,omitempty"`

	// Description - A description of the social topic.
	Description *string `json:"description,omitempty"`

	// DateCreated - Timestamp indicating when the social topic was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Timestamp indicating when the social topic was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DivisionId - The ID of the division to which the social topic belongs to.
	DivisionId *string `json:"divisionId,omitempty"`

	// Status - The status of the social topic.
	Status *string `json:"status,omitempty"`

	// DataIngestionRulesMetadata - The data ingestion rule metadata.
	DataIngestionRulesMetadata *[]Dataingestionrulesmetadata `json:"dataIngestionRulesMetadata,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Socialtopicresponse) SetField(field string, fieldValue interface{}) {
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

func (o Socialtopicresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Socialtopicresponse
	
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
		
		Description *string `json:"description,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DataIngestionRulesMetadata *[]Dataingestionrulesmetadata `json:"dataIngestionRulesMetadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DivisionId: o.DivisionId,
		
		Status: o.Status,
		
		DataIngestionRulesMetadata: o.DataIngestionRulesMetadata,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Socialtopicresponse) UnmarshalJSON(b []byte) error {
	var SocialtopicresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SocialtopicresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SocialtopicresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SocialtopicresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := SocialtopicresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateCreatedString, ok := SocialtopicresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := SocialtopicresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if DivisionId, ok := SocialtopicresponseMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if Status, ok := SocialtopicresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if DataIngestionRulesMetadata, ok := SocialtopicresponseMap["dataIngestionRulesMetadata"].([]interface{}); ok {
		DataIngestionRulesMetadataString, _ := json.Marshal(DataIngestionRulesMetadata)
		json.Unmarshal(DataIngestionRulesMetadataString, &o.DataIngestionRulesMetadata)
	}
	
	if SelfUri, ok := SocialtopicresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Socialtopicresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
