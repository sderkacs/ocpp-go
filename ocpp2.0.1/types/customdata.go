package types

import (
	"encoding/json"
)

// Since all classes in JSON schema have the additionalProperties attribute set to false,
// the standard extends each class with a CustomData data type, which allows for the
// creation of custom extensions and adds some flexibility.
//
// Note: Since this type inherently involves transferring custom data, this library cannot
// guarantee the validity and/or specific encoding of the data. Ensuring data type and
// formatting conformity is the user's responsibility.
type CustomData struct {
	VendorID string                 `json:"vendorId" validate:"required,max=255"`
	Values   map[string]interface{} `json:"-"` // Ignore during default JSON unmarshaling
}

// Creates a new CustomData structure with the specified vendorId and an empty Values map.
func NewCustomData(vendorId string) *CustomData {
	return &CustomData{
		VendorID: vendorId,
		Values:   make(map[string]interface{}, 0),
	}
}

func (c *CustomData) UnmarshalJSON(data []byte) error {
	temp := make(map[string]interface{})

	// Unmarshal all fields into a temporary map.
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// We're trying to retrieve the vendorId value from the map.
	// We're not returning an error here; it will fail later during validation.
	rawVendorID := temp["vendorId"]
	vendorID := ""
	if rawVendorID != nil {
		s, ok := rawVendorID.(string)
		if ok {
			vendorID = s
		}
	}

	*c = CustomData{
		VendorID: vendorID,
		Values:   temp,
	}

	// Remove vendorId from Values to avoid duplication.
	delete(c.Values, "vendorId")

	return nil
}

func (c CustomData) MarshalJSON() ([]byte, error) {
	output := make(map[string]interface{})

	for k, v := range c.Values {
		// Skip the key named vendorId to avoid inconsistencies
		// during marshalling/unmarshalling.
		if k == "vendorId" {
			continue
		}
		output[k] = v
	}
	output["vendorId"] = c.VendorID

	return json.Marshal(output)
}
