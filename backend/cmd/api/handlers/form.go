package handlers

import (
	"encoding/json"
	"fmt"
	user_memory_cache "formaura/pkg/cache/user_memory"
	"formaura/pkg/email"
	"formaura/pkg/output"
	form_repo "formaura/pkg/repositories/form"
	"formaura/pkg/validate"
	"net/http"
	"time"
)

type FormHandler struct {
	FormRepo    form_repo.Repository
	authCache   *user_memory_cache.Cache
	emailClient *email.Client
}

func NewFormHandler(
	repo form_repo.Repository,
	authCache *user_memory_cache.Cache,
	emailClient *email.Client) *FormHandler {
	return &FormHandler{
		FormRepo:    repo,
		authCache:   authCache,
		emailClient: emailClient,
	}
}

type GetListingResponse struct {
	Forms *[]*form_repo.ListingModel `json:"forms"`
}

type GetFormResponse struct {
	Form *form_repo.FormModel `json:"form"`
}

func (h *FormHandler) GetDetailedListing(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetUserFromCtx(r)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	time.Sleep(2 * time.Second)

	listing, err := h.FormRepo.GetDetailedListingByUserID(r.Context(), usr.ID)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Internal server error")
	}

	return output.SuccessResponse(w, r, &GetListingResponse{
		Forms: &listing,
	})
}

func (h *FormHandler) NewForm(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetUserFromCtx(r)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	listing, err := h.FormRepo.GetBasicListingByUserID(r.Context(), usr.ID)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Unable to create a new form")
	}

	newTitle := form_repo.GenerateFormUntitledName(listing)

	blankForm := form_repo.FormData{
		Steps: []form_repo.Step{},
	}

	newForm, err := h.FormRepo.Create(r.Context(), usr.ID, newTitle, nil, blankForm)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("FormRepo.Create: Unable to create a new form")
	}

	return output.SuccessResponse(w, r, &GetFormResponse{
		Form: newForm,
	})
}

func (h *FormHandler) GetForm(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetUserFromCtx(r)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	formUuid, err := GetUUIDFromParams(r)

	if err != nil {
		return http.StatusBadRequest, err
	}

	form, err := h.FormRepo.GetByUUID(r.Context(), *formUuid)

	if err != nil {
		return http.StatusNotFound, fmt.Errorf("Resource not found")
	}

	if form.UserID != usr.ID {
		return http.StatusForbidden, fmt.Errorf("Resource not found")
	}

	return output.SuccessResponse(w, r, &GetFormResponse{
		Form: form,
	})
}

func (h *FormHandler) UpdateFormData(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetUserFromCtx(r)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	return output.SuccessResponse(w, r, &AutoAuthResp{
		User: usr,
	})
}

type UpdateFormMetaReqBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *UpdateFormMetaReqBody) validate() error {
	if !validate.StrNotEmpty(r.Name) {
		return fmt.Errorf("Request body invalid")
	}
	return nil
}

func (h *FormHandler) UpdateFormMeta(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetUserFromCtx(r)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	var body UpdateFormMetaReqBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return http.StatusBadRequest, err
	}
	if err := body.validate(); err != nil {
		return http.StatusBadRequest, err
	}

	formUuid, err := GetUUIDFromParams(r)

	if err != nil {
		return http.StatusBadRequest, err
	}

	form, err := h.FormRepo.GetByUUID(r.Context(), *formUuid)

	if err != nil {
		return http.StatusNotFound, fmt.Errorf("Resource not found")
	}

	if form.UserID != usr.ID {
		return http.StatusForbidden, fmt.Errorf("Resource not found")
	}

	updated, err := h.FormRepo.UpdateFormMeta(r.Context(), form.ID, body.Name, body.Description)

	if err != nil {
		return http.StatusForbidden, fmt.Errorf("Resource not found")
	}

	return output.SuccessResponse(w, r, &GetFormResponse{
		Form: updated,
	})
}

func (h *FormHandler) UpdateFormAffiliates(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetUserFromCtx(r)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	return output.SuccessResponse(w, r, &AutoAuthResp{
		User: usr,
	})
}

func (h *FormHandler) DeleteForm(w http.ResponseWriter, r *http.Request) (int, error) {
	usr, err := GetUserFromCtx(r)

	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Unauthorized")
	}

	return output.SuccessResponse(w, r, &AutoAuthResp{
		User: usr,
	})
}
