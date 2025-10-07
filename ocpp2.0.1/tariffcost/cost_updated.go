package tariffcost

import (
	"reflect"

	"github.com/lorenzodonini/ocpp-go/ocpp2.0.1/types"
)

// -------------------- Cost Updated (CSMS -> CS) --------------------

const CostUpdatedFeatureName = "CostUpdated"

// The field definition of the CostUpdated request payload sent by the CSMS to the Charging Station.
type CostUpdatedRequest struct {
	TotalCost     float64           `json:"totalCost" validate:"required"`
	TransactionID string            `json:"transactionId" validate:"required,max=36"`
	CustomData    *types.CustomData `json:"customData,omitempty" validate:"omitempty"`
}

// This field definition of the CostUpdated response payload, sent by the Charging Station to the CSMS in response to a CostUpdatedRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type CostUpdatedResponse struct {
	CustomData *types.CustomData `json:"customData,omitempty" validate:"omitempty"`
}

// The driver wants to know how much the running total cost is, updated at a relevant interval, while a transaction is ongoing.
// To fulfill this requirement, the CSMS sends a CostUpdatedRequest to the Charging Station to update the current total cost, every Y seconds.
// Upon receipt of the CostUpdatedRequest, the Charging Station responds with a CostUpdatedResponse, then shows the updated cost to the driver.
type CostUpdatedFeature struct{}

func (f CostUpdatedFeature) GetFeatureName() string {
	return CostUpdatedFeatureName
}

func (f CostUpdatedFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(CostUpdatedRequest{})
}

func (f CostUpdatedFeature) GetResponseType() reflect.Type {
	return reflect.TypeOf(CostUpdatedResponse{})
}

func (r CostUpdatedRequest) GetFeatureName() string {
	return CostUpdatedFeatureName
}

func (c CostUpdatedResponse) GetFeatureName() string {
	return CostUpdatedFeatureName
}

// Creates a new CostUpdatedRequest, containing all required fields. There are no optional fields for this message.
func NewCostUpdatedRequest(totalCost float64, transactionID string) *CostUpdatedRequest {
	return &CostUpdatedRequest{TotalCost: totalCost, TransactionID: transactionID}
}

// Creates a new CostUpdatedResponse, which doesn't contain any required or optional fields.
func NewCostUpdatedResponse() *CostUpdatedResponse {
	return &CostUpdatedResponse{}
}
