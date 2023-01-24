package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateuser
type Updateuser struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Chat
	Chat *Chat `json:"chat,omitempty"`

	// Department
	Department *string `json:"department,omitempty"`

	// Email
	Email *string `json:"email,omitempty"`

	// PrimaryContactInfo - The address(s) used for primary contact. Updates to the corresponding address in the addresses list will be reflected here.
	PrimaryContactInfo *[]Contact `json:"primaryContactInfo,omitempty"`

	// Addresses - Email address, phone number, and/or extension for this user. One entry is allowed per media type
	Addresses *[]Contact `json:"addresses,omitempty"`

	// Title
	Title *string `json:"title,omitempty"`

	// Username
	Username *string `json:"username,omitempty"`

	// Manager
	Manager *string `json:"manager,omitempty"`

	// Images
	Images *[]Userimage `json:"images,omitempty"`

	// Version - This value should be the current version of the user. The current version can be obtained with a GET on the user before doing a PATCH.
	Version *int `json:"version,omitempty"`

	// ProfileSkills - Profile skills possessed by the user
	ProfileSkills *[]string `json:"profileSkills,omitempty"`

	// Locations - The user placement at each site location.
	Locations *[]Location `json:"locations,omitempty"`

	// Groups - The groups the user is a member of
	Groups *[]Group `json:"groups,omitempty"`

	// State - The state of the user. This property can be used to restore a deleted user or transition between active and inactive. If specified, it is the only modifiable field.
	State *string `json:"state,omitempty"`

	// AcdAutoAnswer - The value that denotes if acdAutoAnswer is set on the user
	AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`

	// Certifications
	Certifications *[]string `json:"certifications,omitempty"`

	// Biography
	Biography *Biography `json:"biography,omitempty"`

	// EmployerInfo
	EmployerInfo *Employerinfo `json:"employerInfo,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updateuser) SetField(field string, fieldValue interface{}) {
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

func (o Updateuser) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Updateuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Chat *Chat `json:"chat,omitempty"`
		
		Department *string `json:"department,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		PrimaryContactInfo *[]Contact `json:"primaryContactInfo,omitempty"`
		
		Addresses *[]Contact `json:"addresses,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Username *string `json:"username,omitempty"`
		
		Manager *string `json:"manager,omitempty"`
		
		Images *[]Userimage `json:"images,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		ProfileSkills *[]string `json:"profileSkills,omitempty"`
		
		Locations *[]Location `json:"locations,omitempty"`
		
		Groups *[]Group `json:"groups,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`
		
		Certifications *[]string `json:"certifications,omitempty"`
		
		Biography *Biography `json:"biography,omitempty"`
		
		EmployerInfo *Employerinfo `json:"employerInfo,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Chat: o.Chat,
		
		Department: o.Department,
		
		Email: o.Email,
		
		PrimaryContactInfo: o.PrimaryContactInfo,
		
		Addresses: o.Addresses,
		
		Title: o.Title,
		
		Username: o.Username,
		
		Manager: o.Manager,
		
		Images: o.Images,
		
		Version: o.Version,
		
		ProfileSkills: o.ProfileSkills,
		
		Locations: o.Locations,
		
		Groups: o.Groups,
		
		State: o.State,
		
		AcdAutoAnswer: o.AcdAutoAnswer,
		
		Certifications: o.Certifications,
		
		Biography: o.Biography,
		
		EmployerInfo: o.EmployerInfo,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Updateuser) UnmarshalJSON(b []byte) error {
	var UpdateuserMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UpdateuserMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UpdateuserMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Chat, ok := UpdateuserMap["chat"].(map[string]interface{}); ok {
		ChatString, _ := json.Marshal(Chat)
		json.Unmarshal(ChatString, &o.Chat)
	}
	
	if Department, ok := UpdateuserMap["department"].(string); ok {
		o.Department = &Department
	}
    
	if Email, ok := UpdateuserMap["email"].(string); ok {
		o.Email = &Email
	}
    
	if PrimaryContactInfo, ok := UpdateuserMap["primaryContactInfo"].([]interface{}); ok {
		PrimaryContactInfoString, _ := json.Marshal(PrimaryContactInfo)
		json.Unmarshal(PrimaryContactInfoString, &o.PrimaryContactInfo)
	}
	
	if Addresses, ok := UpdateuserMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if Title, ok := UpdateuserMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Username, ok := UpdateuserMap["username"].(string); ok {
		o.Username = &Username
	}
    
	if Manager, ok := UpdateuserMap["manager"].(string); ok {
		o.Manager = &Manager
	}
    
	if Images, ok := UpdateuserMap["images"].([]interface{}); ok {
		ImagesString, _ := json.Marshal(Images)
		json.Unmarshal(ImagesString, &o.Images)
	}
	
	if Version, ok := UpdateuserMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ProfileSkills, ok := UpdateuserMap["profileSkills"].([]interface{}); ok {
		ProfileSkillsString, _ := json.Marshal(ProfileSkills)
		json.Unmarshal(ProfileSkillsString, &o.ProfileSkills)
	}
	
	if Locations, ok := UpdateuserMap["locations"].([]interface{}); ok {
		LocationsString, _ := json.Marshal(Locations)
		json.Unmarshal(LocationsString, &o.Locations)
	}
	
	if Groups, ok := UpdateuserMap["groups"].([]interface{}); ok {
		GroupsString, _ := json.Marshal(Groups)
		json.Unmarshal(GroupsString, &o.Groups)
	}
	
	if State, ok := UpdateuserMap["state"].(string); ok {
		o.State = &State
	}
    
	if AcdAutoAnswer, ok := UpdateuserMap["acdAutoAnswer"].(bool); ok {
		o.AcdAutoAnswer = &AcdAutoAnswer
	}
    
	if Certifications, ok := UpdateuserMap["certifications"].([]interface{}); ok {
		CertificationsString, _ := json.Marshal(Certifications)
		json.Unmarshal(CertificationsString, &o.Certifications)
	}
	
	if Biography, ok := UpdateuserMap["biography"].(map[string]interface{}); ok {
		BiographyString, _ := json.Marshal(Biography)
		json.Unmarshal(BiographyString, &o.Biography)
	}
	
	if EmployerInfo, ok := UpdateuserMap["employerInfo"].(map[string]interface{}); ok {
		EmployerInfoString, _ := json.Marshal(EmployerInfo)
		json.Unmarshal(EmployerInfoString, &o.EmployerInfo)
	}
	
	if SelfUri, ok := UpdateuserMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Updateuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
