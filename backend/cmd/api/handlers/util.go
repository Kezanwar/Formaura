package handlers

import (
	"fmt"
	"formaura/pkg/constants"
	user_repo "formaura/pkg/repositories/user"
	"net/http"
)

func GetUserFromCtx(r *http.Request) (*user_repo.Model, error) {
	usr, ok := r.Context().Value(constants.USER_CTX).(*user_repo.Model)

	if !ok {
		return nil, fmt.Errorf("handlers.GetUserFromContext: cant find user in r.Context")
	}

	return usr, nil

}
