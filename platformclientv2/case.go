package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Case
type Case struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the Case.
	Name *string `json:"name,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Starrabledivision `json:"division,omitempty"`

	// Version - The version of the Case.
	Version *int `json:"version,omitempty"`

	// Reference - The reference identifier of the Case.
	Reference *string `json:"reference,omitempty"`

	// Caseplan - The Caseplan the case was created from.
	Caseplan *Caseplanreference `json:"caseplan,omitempty"`

	// Summary - Overview information for the Case.
	Summary *string `json:"summary,omitempty"`

	// Owner - The owner of the Case.
	Owner *Userreference `json:"owner,omitempty"`

	// Status - The status of the Case.
	Status *string `json:"status,omitempty"`

	// Priority - The priority of the Case.
	Priority *string `json:"priority,omitempty"`

	// DateDue - The due date of the Case. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDue *time.Time `json:"dateDue,omitempty"`

	// DateStarted - The start time of the Case. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`

	// DateClosed - The completion time of the Case. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateClosed *time.Time `json:"dateClosed,omitempty"`

	// DateCreated - The date the Case was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The date the Case was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// ModifiedBy - The id of the User who modified the Case.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// ExternalContact - The External Contact associated with the Case.
	ExternalContact *Caseexternalcontactreference `json:"externalContact,omitempty"`

	// CustomerIntent - The customer intent for the Case.
	CustomerIntent *Customerintentreference `json:"customerIntent,omitempty"`

	// CreationStatus - The creation status of the Case
	CreationStatus *string `json:"creationStatus,omitempty"`

	// TtlSeconds - The time-to-live in seconds for the lifetime of the Case.
	TtlSeconds *int `json:"ttlSeconds,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Case) SetField(field string, fieldValue interface{}) {
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

func (o Case) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateDue","DateStarted","DateClosed","DateCreated","DateModified", }
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
	type Alias Case
	
	DateDue := new(string)
	if o.DateDue != nil {
		
		*DateDue = timeutil.Strftime(o.DateDue, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDue = nil
	}
	
	DateStarted := new(string)
	if o.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(o.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	
	DateClosed := new(string)
	if o.DateClosed != nil {
		
		*DateClosed = timeutil.Strftime(o.DateClosed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateClosed = nil
	}
	
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
		
		Division *Starrabledivision `json:"division,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Reference *string `json:"reference,omitempty"`
		
		Caseplan *Caseplanreference `json:"caseplan,omitempty"`
		
		Summary *string `json:"summary,omitempty"`
		
		Owner *Userreference `json:"owner,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Priority *string `json:"priority,omitempty"`
		
		DateDue *string `json:"dateDue,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		
		DateClosed *string `json:"dateClosed,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		ExternalContact *Caseexternalcontactreference `json:"externalContact,omitempty"`
		
		CustomerIntent *Customerintentreference `json:"customerIntent,omitempty"`
		
		CreationStatus *string `json:"creationStatus,omitempty"`
		
		TtlSeconds *int `json:"ttlSeconds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Version: o.Version,
		
		Reference: o.Reference,
		
		Caseplan: o.Caseplan,
		
		Summary: o.Summary,
		
		Owner: o.Owner,
		
		Status: o.Status,
		
		Priority: o.Priority,
		
		DateDue: DateDue,
		
		DateStarted: DateStarted,
		
		DateClosed: DateClosed,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		ExternalContact: o.ExternalContact,
		
		CustomerIntent: o.CustomerIntent,
		
		CreationStatus: o.CreationStatus,
		
		TtlSeconds: o.TtlSeconds,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Case) UnmarshalJSON(b []byte) error {
	var CaseMap map[string]interface{}
	err := json.Unmarshal(b, &CaseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CaseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CaseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := CaseMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Version, ok := CaseMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Reference, ok := CaseMap["reference"].(string); ok {
		o.Reference = &Reference
	}
    
	if Caseplan, ok := CaseMap["caseplan"].(map[string]interface{}); ok {
		CaseplanString, _ := json.Marshal(Caseplan)
		json.Unmarshal(CaseplanString, &o.Caseplan)
	}
	
	if Summary, ok := CaseMap["summary"].(string); ok {
		o.Summary = &Summary
	}
    
	if Owner, ok := CaseMap["owner"].(map[string]interface{}); ok {
		OwnerString, _ := json.Marshal(Owner)
		json.Unmarshal(OwnerString, &o.Owner)
	}
	
	if Status, ok := CaseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Priority, ok := CaseMap["priority"].(string); ok {
		o.Priority = &Priority
	}
    
	if dateDueString, ok := CaseMap["dateDue"].(string); ok {
		DateDue, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDueString)
		o.DateDue = &DateDue
	}
	
	if dateStartedString, ok := CaseMap["dateStarted"].(string); ok {
		DateStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartedString)
		o.DateStarted = &DateStarted
	}
	
	if dateClosedString, ok := CaseMap["dateClosed"].(string); ok {
		DateClosed, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateClosedString)
		o.DateClosed = &DateClosed
	}
	
	if dateCreatedString, ok := CaseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := CaseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := CaseMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if ExternalContact, ok := CaseMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if CustomerIntent, ok := CaseMap["customerIntent"].(map[string]interface{}); ok {
		CustomerIntentString, _ := json.Marshal(CustomerIntent)
		json.Unmarshal(CustomerIntentString, &o.CustomerIntent)
	}
	
	if CreationStatus, ok := CaseMap["creationStatus"].(string); ok {
		o.CreationStatus = &CreationStatus
	}
    
	if TtlSeconds, ok := CaseMap["ttlSeconds"].(float64); ok {
		TtlSecondsInt := int(TtlSeconds)
		o.TtlSeconds = &TtlSecondsInt
	}
	
	if SelfUri, ok := CaseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Case) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
