package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchsegment
type Patchsegment struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// IsActive - Whether or not the segment is active.
	IsActive *bool `json:"isActive,omitempty"`


	// DisplayName - The display name of the segment.
	DisplayName *string `json:"displayName,omitempty"`


	// Version - The version of the segment.
	Version *int `json:"version,omitempty"`


	// Description - A description of the segment.
	Description *string `json:"description,omitempty"`


	// Color - The hexadecimal color value of the segment.
	Color *string `json:"color,omitempty"`


	// ShouldDisplayToAgent - Whether or not the segment should be displayed to agent/supervisor users.
	ShouldDisplayToAgent *bool `json:"shouldDisplayToAgent,omitempty"`


	// Context - The context of the segment.
	Context *Context `json:"context,omitempty"`


	// Journey - The pattern of rules defining the segment.
	Journey *Journey `json:"journey,omitempty"`


	// ExternalSegment - Details of an entity corresponding to this segment in an external system.
	ExternalSegment *Patchexternalsegment `json:"externalSegment,omitempty"`


	// AssignmentExpirationDays - Time, in days, from when the segment is assigned until it is automatically unassigned.
	AssignmentExpirationDays *int `json:"assignmentExpirationDays,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// CreatedDate - Timestamp indicating when the segment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Timestamp indicating when the the segment was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (o *Patchsegment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchsegment
	
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
		
		Color *string `json:"color,omitempty"`
		
		ShouldDisplayToAgent *bool `json:"shouldDisplayToAgent,omitempty"`
		
		Context *Context `json:"context,omitempty"`
		
		Journey *Journey `json:"journey,omitempty"`
		
		ExternalSegment *Patchexternalsegment `json:"externalSegment,omitempty"`
		
		AssignmentExpirationDays *int `json:"assignmentExpirationDays,omitempty"`
		
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
		
		Color: o.Color,
		
		ShouldDisplayToAgent: o.ShouldDisplayToAgent,
		
		Context: o.Context,
		
		Journey: o.Journey,
		
		ExternalSegment: o.ExternalSegment,
		
		AssignmentExpirationDays: o.AssignmentExpirationDays,
		
		SelfUri: o.SelfUri,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchsegment) UnmarshalJSON(b []byte) error {
	var PatchsegmentMap map[string]interface{}
	err := json.Unmarshal(b, &PatchsegmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PatchsegmentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if IsActive, ok := PatchsegmentMap["isActive"].(bool); ok {
		o.IsActive = &IsActive
	}
    
	if DisplayName, ok := PatchsegmentMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if Version, ok := PatchsegmentMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Description, ok := PatchsegmentMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Color, ok := PatchsegmentMap["color"].(string); ok {
		o.Color = &Color
	}
    
	if ShouldDisplayToAgent, ok := PatchsegmentMap["shouldDisplayToAgent"].(bool); ok {
		o.ShouldDisplayToAgent = &ShouldDisplayToAgent
	}
    
	if Context, ok := PatchsegmentMap["context"].(map[string]interface{}); ok {
		ContextString, _ := json.Marshal(Context)
		json.Unmarshal(ContextString, &o.Context)
	}
	
	if Journey, ok := PatchsegmentMap["journey"].(map[string]interface{}); ok {
		JourneyString, _ := json.Marshal(Journey)
		json.Unmarshal(JourneyString, &o.Journey)
	}
	
	if ExternalSegment, ok := PatchsegmentMap["externalSegment"].(map[string]interface{}); ok {
		ExternalSegmentString, _ := json.Marshal(ExternalSegment)
		json.Unmarshal(ExternalSegmentString, &o.ExternalSegment)
	}
	
	if AssignmentExpirationDays, ok := PatchsegmentMap["assignmentExpirationDays"].(float64); ok {
		AssignmentExpirationDaysInt := int(AssignmentExpirationDays)
		o.AssignmentExpirationDays = &AssignmentExpirationDaysInt
	}
	
	if SelfUri, ok := PatchsegmentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if createdDateString, ok := PatchsegmentMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := PatchsegmentMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
