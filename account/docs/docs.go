// Package docs finansiyer Service API.
//
// Documentation for finansiyer Service API
//
//	Schemes: https, http
//	BasePath: ./
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package docs

// swagger:parameters healthRequest
type healthRequest struct{}

// Success
// swagger:response healthResponse
type healthResponse struct {
	// in:body
	Body struct {
		Data   *healthData `json:"data"`
		Result string      `json:"result,omitempty"`
	}
}

type healthData struct {
	Ping string `json:"ping"`
}
