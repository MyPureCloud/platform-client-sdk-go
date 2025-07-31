package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Organizationpublicapiusageresultsresponse
type Organizationpublicapiusageresultsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name
	Name *string `json:"name,omitempty"`

	// QueryStatus - The status of the query.
	QueryStatus *string `json:"queryStatus,omitempty"`

	// ErrorBody - The reason for the failure. This will only be present if the query is in a FAILURE state but may not be included even if the state is FAILURE
	ErrorBody *Errorbody `json:"errorBody,omitempty"`

	// NextUri - The uri to get the next set of results. Will only be included if there is another page to retrieve.
	NextUri *string `json:"nextUri,omitempty"`

	// Entities - The results of the query.
	Entities *[]Organizationpublicapiusage `json:"entities,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Organizationpublicapiusageresultsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Organizationpublicapiusageresultsresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Organizationpublicapiusageresultsresponse
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		QueryStatus *string `json:"queryStatus,omitempty"`
		
		ErrorBody *Errorbody `json:"errorBody,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		Entities *[]Organizationpublicapiusage `json:"entities,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		QueryStatus: o.QueryStatus,
		
		ErrorBody: o.ErrorBody,
		
		NextUri: o.NextUri,
		
		Entities: o.Entities,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Organizationpublicapiusageresultsresponse) UnmarshalJSON(b []byte) error {
	var OrganizationpublicapiusageresultsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &OrganizationpublicapiusageresultsresponseMap)
	if err != nil {
		return err
	}
	
	if Name, ok := OrganizationpublicapiusageresultsresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if QueryStatus, ok := OrganizationpublicapiusageresultsresponseMap["queryStatus"].(string); ok {
		o.QueryStatus = &QueryStatus
	}
    
	if ErrorBody, ok := OrganizationpublicapiusageresultsresponseMap["errorBody"].(map[string]interface{}); ok {
		ErrorBodyString, _ := json.Marshal(ErrorBody)
		json.Unmarshal(ErrorBodyString, &o.ErrorBody)
	}
	
	if NextUri, ok := OrganizationpublicapiusageresultsresponseMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if Entities, ok := OrganizationpublicapiusageresultsresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if SelfUri, ok := OrganizationpublicapiusageresultsresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Organizationpublicapiusageresultsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
