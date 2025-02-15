// internal/api/handlers/status.go
package handlers

import (
	"assignment_1/internal/core/status"
	"assignment_1/internal/responses"
	"net/http"
)

var statusSvc status.Service

// InitStatusService initializes the status service
func InitStatusService(svc status.Service) {
	statusSvc = svc
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONResponse(w, http.StatusMethodNotAllowed, responses.ErrorResponse{
			Error: "Method not allowed",
		})
		return
	}

	status, err := statusSvc.GetStatus(r.Context())
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, responses.ErrorResponse{
			Error: "Failed to get status",
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, status)
}
