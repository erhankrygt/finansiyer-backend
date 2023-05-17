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

// swagger:parameters registerRequest
type registerRequest struct {
	// in: body
	// required: true
	Body struct {
		// example: John
		FirstName string `json:"firstName"`
		// example: Doe
		LastName string `json:"lastName"`
		// example: john@finansiyer.com
		Email string `json:"email"`
		// example: 539xxxxxxx
		PhoneNumber string `json:"phoneNumber"`
		// example: 12345678
		Password string `json:"password"`
		// example: 12345678
		ConfirmPassword string `json:"confirmPassword"`
	}
}

// Success
// swagger:response registerResponse
type registerResponse struct {
	// in:body
	Body struct {
		Data   *registerData `json:"data"`
		Result *apiError     `json:"result,omitempty"`
	}
}

type registerData struct {
	IsSuccessful bool `json:"isSuccessful"`
}

type apiError struct {
	Code    int
	Message string
}
