package handlers

import (
	"context"
	"fmt"
	"formaura/pkg/email"
	"formaura/pkg/output"
	form_repo "formaura/pkg/repositories/form"
	"net/http"
)

type SubmissionHandler struct {
	FormRepo    form_repo.Repository
	emailClient *email.Client
}

func NewSubmissionHandler(
	repo form_repo.Repository,
	emailClient *email.Client) *SubmissionHandler {
	return &SubmissionHandler{
		FormRepo:    repo,
		emailClient: emailClient,
	}
}

func (h *SubmissionHandler) GetForm(w http.ResponseWriter, r *http.Request) (int, error) {

	formUuid, err := GetUUIDFromParams(r)

	if err != nil {
		return http.StatusBadRequest, err
	}

	form, err := h.FormRepo.GetByUUID(r.Context(), *formUuid)

	if err != nil {
		return http.StatusNotFound, fmt.Errorf("Resource not found")
	}

	if form.Status != form_repo.StatusActive {
		return http.StatusForbidden, fmt.Errorf("This form is currently unavailable")
	}

	err = h.FormRepo.IncrementViews(context.Background(), *formUuid)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to view form, please try again later")
	}

	return output.SuccessResponse(w, r, &GetFormResponse{
		Form: form,
	})
}
