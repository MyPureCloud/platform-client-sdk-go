package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeycaseassociation - The representation of a case association on a journey session.
type Journeycaseassociation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the association.
	Id *string `json:"id,omitempty"`

	// AssociatedCase - The case that was associated with the journey session.
	AssociatedCase *Addressableentityref `json:"associatedCase,omitempty"`

	// CaseReference - The reference for the case that was associated with the journey session.
	CaseReference *string `json:"caseReference,omitempty"`

	// DateAssociated - The date of the association. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateAssociated *time.Time `json:"dateAssociated,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeycaseassociation) SetField(field string, fieldValue interface{}) {
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

func (o Journeycaseassociation) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateAssociated", }
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
	type Alias Journeycaseassociation
	
	DateAssociated := new(string)
	if o.DateAssociated != nil {
		
		*DateAssociated = timeutil.Strftime(o.DateAssociated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateAssociated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		AssociatedCase *Addressableentityref `json:"associatedCase,omitempty"`
		
		CaseReference *string `json:"caseReference,omitempty"`
		
		DateAssociated *string `json:"dateAssociated,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		AssociatedCase: o.AssociatedCase,
		
		CaseReference: o.CaseReference,
		
		DateAssociated: DateAssociated,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Journeycaseassociation) UnmarshalJSON(b []byte) error {
	var JourneycaseassociationMap map[string]interface{}
	err := json.Unmarshal(b, &JourneycaseassociationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneycaseassociationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if AssociatedCase, ok := JourneycaseassociationMap["associatedCase"].(map[string]interface{}); ok {
		AssociatedCaseString, _ := json.Marshal(AssociatedCase)
		json.Unmarshal(AssociatedCaseString, &o.AssociatedCase)
	}
	
	if CaseReference, ok := JourneycaseassociationMap["caseReference"].(string); ok {
		o.CaseReference = &CaseReference
	}
    
	if dateAssociatedString, ok := JourneycaseassociationMap["dateAssociated"].(string); ok {
		DateAssociated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAssociatedString)
		o.DateAssociated = &DateAssociated
	}
	
	if SelfUri, ok := JourneycaseassociationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeycaseassociation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
