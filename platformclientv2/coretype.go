package platformclientv2
import (
	"time"
	"encoding/json"
)

// Coretype
type Coretype struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Version - A positive integer denoting the core type's version
	Version *int32 `json:"version,omitempty"`


	// DateCreated - The date the core type was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// Schema - The core type's built-in schema
	Schema *Schema `json:"schema,omitempty"`


	// Current - A boolean indicating if the core type's version is the current one in use by the system
	Current *bool `json:"current,omitempty"`


	// ValidationFields - An array of strings naming the fields of the core type subject to validation.  Validation constraints are specified by a schema author using the core type.
	ValidationFields *[]string `json:"validationFields,omitempty"`


	// ValidationLimits - A structure denoting the system-imposed minimum and maximum string length (for text-based core types) or numeric values (for number-based) core types.  For example, the validationLimits for a text-based core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schemaauthor on a text field.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field.
	ValidationLimits *Validationlimits `json:"validationLimits,omitempty"`


	// ItemValidationFields - Specific to the \"tag\" core type, this is an array of strings naming the tag item fields of the core type subject to validation
	ItemValidationFields *[]string `json:"itemValidationFields,omitempty"`


	// ItemValidationLimits - A structure denoting the system-imposed minimum and maximum string length for string-array based core types such as \"tag\" and \"enum\".  Forexample, the validationLimits for a schema field using a tag core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schema author on individual tags.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field's tags.
	ItemValidationLimits *Itemvalidationlimits `json:"itemValidationLimits,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coretype) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
