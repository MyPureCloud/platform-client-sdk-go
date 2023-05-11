package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Acceleratorspecification - Metadata for a CX infrastructure as code accelerator
type Acceleratorspecification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - name of this accelerator
	Name *string `json:"name,omitempty"`

	// Description - a description of the general purpose of this accelerator
	Description *string `json:"description,omitempty"`

	// Origin - where the accelerator originated
	Origin *string `json:"origin,omitempty"`

	// VarType - type of the artifact
	VarType *string `json:"type,omitempty"`

	// Classification - architectural classification into which the accelerator belongs
	Classification *string `json:"classification,omitempty"`

	// Tags - tags
	Tags *[]string `json:"tags,omitempty"`

	// Permissions - Genesys Cloud permissions required to install the accelerator
	Permissions *[]string `json:"permissions,omitempty"`

	// Products - Genesys Cloud products required to install the accelerator
	Products *[]string `json:"products,omitempty"`

	// Documentation - additional documentation about the artifact
	Documentation *[]Metadatadocumentation `json:"documentation,omitempty"`

	// Presentation - presentation of data fields to be gathered for the accelerator
	Presentation *[]Metadatapresentation `json:"presentation,omitempty"`

	// Results - resources created or modified as a result of running the accelerator
	Results *Metadataresults `json:"results,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Acceleratorspecification) SetField(field string, fieldValue interface{}) {
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

func (o Acceleratorspecification) MarshalJSON() ([]byte, error) {
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
	type Alias Acceleratorspecification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Origin *string `json:"origin,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Classification *string `json:"classification,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Permissions *[]string `json:"permissions,omitempty"`
		
		Products *[]string `json:"products,omitempty"`
		
		Documentation *[]Metadatadocumentation `json:"documentation,omitempty"`
		
		Presentation *[]Metadatapresentation `json:"presentation,omitempty"`
		
		Results *Metadataresults `json:"results,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Origin: o.Origin,
		
		VarType: o.VarType,
		
		Classification: o.Classification,
		
		Tags: o.Tags,
		
		Permissions: o.Permissions,
		
		Products: o.Products,
		
		Documentation: o.Documentation,
		
		Presentation: o.Presentation,
		
		Results: o.Results,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Acceleratorspecification) UnmarshalJSON(b []byte) error {
	var AcceleratorspecificationMap map[string]interface{}
	err := json.Unmarshal(b, &AcceleratorspecificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AcceleratorspecificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AcceleratorspecificationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := AcceleratorspecificationMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Origin, ok := AcceleratorspecificationMap["origin"].(string); ok {
		o.Origin = &Origin
	}
    
	if VarType, ok := AcceleratorspecificationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Classification, ok := AcceleratorspecificationMap["classification"].(string); ok {
		o.Classification = &Classification
	}
    
	if Tags, ok := AcceleratorspecificationMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if Permissions, ok := AcceleratorspecificationMap["permissions"].([]interface{}); ok {
		PermissionsString, _ := json.Marshal(Permissions)
		json.Unmarshal(PermissionsString, &o.Permissions)
	}
	
	if Products, ok := AcceleratorspecificationMap["products"].([]interface{}); ok {
		ProductsString, _ := json.Marshal(Products)
		json.Unmarshal(ProductsString, &o.Products)
	}
	
	if Documentation, ok := AcceleratorspecificationMap["documentation"].([]interface{}); ok {
		DocumentationString, _ := json.Marshal(Documentation)
		json.Unmarshal(DocumentationString, &o.Documentation)
	}
	
	if Presentation, ok := AcceleratorspecificationMap["presentation"].([]interface{}); ok {
		PresentationString, _ := json.Marshal(Presentation)
		json.Unmarshal(PresentationString, &o.Presentation)
	}
	
	if Results, ok := AcceleratorspecificationMap["results"].(map[string]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if SelfUri, ok := AcceleratorspecificationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Acceleratorspecification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
