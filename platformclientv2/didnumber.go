package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Didnumber - Represents an unassigned or assigned DID in a DID Pool.
type Didnumber struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Number - The number of the DID formatted as E164.
	Number *string `json:"number,omitempty"`

	// Assigned - True if this DID is assigned to an entity.  False otherwise.
	Assigned *bool `json:"assigned,omitempty"`

	// DidPool - A Uri reference to the DID Pool this DID is a part of.
	DidPool *Addressableentityref `json:"didPool,omitempty"`

	// Owner - A Uri reference to the owner of this DID.  The owner's type can be found in ownerType.  If the DID is unassigned, this will be NULL.
	Owner *Domainentityref `json:"owner,omitempty"`

	// OwnerType - The type of the entity that owns this DID.  If the DID is unassigned, this will be NULL.
	OwnerType *string `json:"ownerType,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Didnumber) SetField(field string, fieldValue interface{}) {
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

func (o Didnumber) MarshalJSON() ([]byte, error) {
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
	type Alias Didnumber
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Number *string `json:"number,omitempty"`
		
		Assigned *bool `json:"assigned,omitempty"`
		
		DidPool *Addressableentityref `json:"didPool,omitempty"`
		
		Owner *Domainentityref `json:"owner,omitempty"`
		
		OwnerType *string `json:"ownerType,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Number: o.Number,
		
		Assigned: o.Assigned,
		
		DidPool: o.DidPool,
		
		Owner: o.Owner,
		
		OwnerType: o.OwnerType,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Didnumber) UnmarshalJSON(b []byte) error {
	var DidnumberMap map[string]interface{}
	err := json.Unmarshal(b, &DidnumberMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DidnumberMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DidnumberMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Number, ok := DidnumberMap["number"].(string); ok {
		o.Number = &Number
	}
    
	if Assigned, ok := DidnumberMap["assigned"].(bool); ok {
		o.Assigned = &Assigned
	}
    
	if DidPool, ok := DidnumberMap["didPool"].(map[string]interface{}); ok {
		DidPoolString, _ := json.Marshal(DidPool)
		json.Unmarshal(DidPoolString, &o.DidPool)
	}
	
	if Owner, ok := DidnumberMap["owner"].(map[string]interface{}); ok {
		OwnerString, _ := json.Marshal(Owner)
		json.Unmarshal(OwnerString, &o.Owner)
	}
	
	if OwnerType, ok := DidnumberMap["ownerType"].(string); ok {
		o.OwnerType = &OwnerType
	}
    
	if SelfUri, ok := DidnumberMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Didnumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
