// Code generated by microgen 1.0.4. DO NOT EDIT.

package transport

import types "esgrs/pkg/types"

type (
	GenerateVisualizationRequest struct {
		Result types.ESGRatingResult `json:"result"`
	}
	GenerateVisualizationResponse struct {
		Vis types.Visualization `json:"vis"`
	}
)