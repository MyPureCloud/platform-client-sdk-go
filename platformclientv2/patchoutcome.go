package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchoutcome
type Patchoutcome struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// IsActive - Whether or not the outcome is active.
	IsActive *bool `json:"isActive,omitempty"`


	// DisplayName - The display name of the outcome.
	DisplayName *string `json:"displayName,omitempty"`


	// Version - The version of the outcome.
	Version *int `json:"version,omitempty"`


	// Description - A description of the outcome.
	Description *string `json:"description,omitempty"`


	// IsPositive - Whether or not the outcome is positive.
	IsPositive *bool `json:"isPositive,omitempty"`


	// Context - The context of the outcome.
	Context *Context `json:"context,omitempty"`


	// Journey - The pattern of rules defining the filter of the outcome.
	Journey *Journey `json:"journey,omitempty"`


	// AssociatedValueField - The field from the event indicating the associated value.
	AssociatedValueField *Associatedvaluefield `json:"associatedValueField,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// CreatedDate - Timestamp indicating when the outcome was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Timestamp indicating when the outcome was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (o *Patchoutcome) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchoutcome
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		IsActive *bool `json:"isActive,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		IsPositive *bool `json:"isPositive,omitempty"`
		
		Context *Context `json:"context,omitempty"`
		
		Journey *Journey `json:"journey,omitempty"`
		
		AssociatedValueField *Associatedvaluefield `json:"associatedValueField,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		IsActive: o.IsActive,
		
		DisplayName: o.DisplayName,
		
		Version: o.Version,
		
		Description: o.Description,
		
		IsPositive: o.IsPositive,
		
		Context: o.Context,
		
		Journey: o.Journey,
		
		AssociatedValueField: o.AssociatedValueField,
		
		SelfUri: o.SelfUri,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchoutcome) UnmarshalJSON(b []byte) error {
	var PatchoutcomeMap map[string]interface{}
	err := json.Unmarshal(b, &PatchoutcomeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PatchoutcomeMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if IsActive, ok := PatchoutcomeMap["isActive"].(bool); ok {
		o.IsActive = &IsActive
	}
	
	if DisplayName, ok := PatchoutcomeMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
	
	if Version, ok := PatchoutcomeMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Description, ok := PatchoutcomeMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if IsPositive, ok := PatchoutcomeMap["isPositive"].(bool); ok {
		o.IsPositive = &IsPositive
	}
	
	if Context, ok := PatchoutcomeMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	
	if Journey, ok := PatchoutcomeMap["journey"].(map[string]interface{}); ok {
		JourneyString, _ := json.Marshal(Journey)
		json.Unmarshal(JourneyString, &o.Journey)
	}
	
	if AssociatedValueField, ok := PatchoutcomeMap["associatedValueField"].(map[string]interface{}); ok {
		AssociatedValueFieldString, _ := json.Marshal(AssociatedValueField)
		json.Unmarshal(AssociatedValueFieldString, &o.AssociatedValueField)
	}
	
	if SelfUri, ok := PatchoutcomeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if createdDateString, ok := PatchoutcomeMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := PatchoutcomeMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchoutcome) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
