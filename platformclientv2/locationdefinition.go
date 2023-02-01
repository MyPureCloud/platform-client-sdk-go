package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationdefinition
type Locationdefinition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// ContactUser - Site contact for the location entity
	ContactUser *Addressableentityref `json:"contactUser,omitempty"`

	// EmergencyNumber - Emergency number for the location entity
	EmergencyNumber *Locationemergencynumber `json:"emergencyNumber,omitempty"`

	// Address
	Address *Locationaddress `json:"address,omitempty"`

	// State - Current state of the location entity
	State *string `json:"state,omitempty"`

	// Notes - Notes for the location entity
	Notes *string `json:"notes,omitempty"`

	// Version - Current version of the location entity, value to be supplied should be retrieved by a GET or on create/update response
	Version *int `json:"version,omitempty"`

	// Path - A list of ancestor IDs in order
	Path *[]string `json:"path,omitempty"`

	// ProfileImage - Profile image of the location entity, retrieved with ?expand=images query parameter
	ProfileImage *[]Locationimage `json:"profileImage,omitempty"`

	// FloorplanImage - Floorplan images of the location entity, retrieved with ?expand=images query parameter
	FloorplanImage *[]Locationimage `json:"floorplanImage,omitempty"`

	// AddressVerificationDetails - Address verification information, retrieve dwith the ?expand=addressVerificationDetails query parameter
	AddressVerificationDetails *Locationaddressverificationdetails `json:"addressVerificationDetails,omitempty"`

	// AddressVerified - Boolean field which states if the address has been verified as an actual address
	AddressVerified *bool `json:"addressVerified,omitempty"`

	// AddressStored - Boolean field which states if the address has been stored for E911
	AddressStored *bool `json:"addressStored,omitempty"`

	// Images
	Images *string `json:"images,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Locationdefinition) SetField(field string, fieldValue interface{}) {
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

func (o Locationdefinition) MarshalJSON() ([]byte, error) {
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
	type Alias Locationdefinition
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ContactUser *Addressableentityref `json:"contactUser,omitempty"`
		
		EmergencyNumber *Locationemergencynumber `json:"emergencyNumber,omitempty"`
		
		Address *Locationaddress `json:"address,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Path *[]string `json:"path,omitempty"`
		
		ProfileImage *[]Locationimage `json:"profileImage,omitempty"`
		
		FloorplanImage *[]Locationimage `json:"floorplanImage,omitempty"`
		
		AddressVerificationDetails *Locationaddressverificationdetails `json:"addressVerificationDetails,omitempty"`
		
		AddressVerified *bool `json:"addressVerified,omitempty"`
		
		AddressStored *bool `json:"addressStored,omitempty"`
		
		Images *string `json:"images,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ContactUser: o.ContactUser,
		
		EmergencyNumber: o.EmergencyNumber,
		
		Address: o.Address,
		
		State: o.State,
		
		Notes: o.Notes,
		
		Version: o.Version,
		
		Path: o.Path,
		
		ProfileImage: o.ProfileImage,
		
		FloorplanImage: o.FloorplanImage,
		
		AddressVerificationDetails: o.AddressVerificationDetails,
		
		AddressVerified: o.AddressVerified,
		
		AddressStored: o.AddressStored,
		
		Images: o.Images,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Locationdefinition) UnmarshalJSON(b []byte) error {
	var LocationdefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &LocationdefinitionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LocationdefinitionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := LocationdefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ContactUser, ok := LocationdefinitionMap["contactUser"].(map[string]interface{}); ok {
		ContactUserString, _ := json.Marshal(ContactUser)
		json.Unmarshal(ContactUserString, &o.ContactUser)
	}
	
	if EmergencyNumber, ok := LocationdefinitionMap["emergencyNumber"].(map[string]interface{}); ok {
		EmergencyNumberString, _ := json.Marshal(EmergencyNumber)
		json.Unmarshal(EmergencyNumberString, &o.EmergencyNumber)
	}
	
	if Address, ok := LocationdefinitionMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	
	if State, ok := LocationdefinitionMap["state"].(string); ok {
		o.State = &State
	}
    
	if Notes, ok := LocationdefinitionMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if Version, ok := LocationdefinitionMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Path, ok := LocationdefinitionMap["path"].([]interface{}); ok {
		PathString, _ := json.Marshal(Path)
		json.Unmarshal(PathString, &o.Path)
	}
	
	if ProfileImage, ok := LocationdefinitionMap["profileImage"].([]interface{}); ok {
		ProfileImageString, _ := json.Marshal(ProfileImage)
		json.Unmarshal(ProfileImageString, &o.ProfileImage)
	}
	
	if FloorplanImage, ok := LocationdefinitionMap["floorplanImage"].([]interface{}); ok {
		FloorplanImageString, _ := json.Marshal(FloorplanImage)
		json.Unmarshal(FloorplanImageString, &o.FloorplanImage)
	}
	
	if AddressVerificationDetails, ok := LocationdefinitionMap["addressVerificationDetails"].(map[string]interface{}); ok {
		AddressVerificationDetailsString, _ := json.Marshal(AddressVerificationDetails)
		json.Unmarshal(AddressVerificationDetailsString, &o.AddressVerificationDetails)
	}
	
	if AddressVerified, ok := LocationdefinitionMap["addressVerified"].(bool); ok {
		o.AddressVerified = &AddressVerified
	}
    
	if AddressStored, ok := LocationdefinitionMap["addressStored"].(bool); ok {
		o.AddressStored = &AddressStored
	}
    
	if Images, ok := LocationdefinitionMap["images"].(string); ok {
		o.Images = &Images
	}
    
	if SelfUri, ok := LocationdefinitionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Locationdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
