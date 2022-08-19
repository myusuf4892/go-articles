package helper

import (
	"net/http"
)

func ResponseStatusOkNoData(msg string) (int, map[string]interface{}) {
	return http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": msg,
	}
}

func ResponseStatusOkWithData(msg string, data interface{}) (int, map[string]interface{}) {
	return http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": msg,
		"data":    data,
	}
}

func ResponseStatusOkWithDataPage(msg string, page int64, data interface{}) (int, map[string]interface{}) {
	return http.StatusOK, map[string]interface{}{
		"code":      "200",
		"message":   msg,
		"totalpage": page,
		"data":      data,
	}
}

func ResponseBadRequest(msg string) (int, map[string]interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"code":    "400",
		"message": msg,
	}
}

func ResponseNotFound(msg string) (int, map[string]interface{}) {
	return http.StatusNotFound, map[string]interface{}{
		"code":    "404",
		"message": msg,
	}
}

func ResponseInternalServerError(msg string) (int, map[string]interface{}) {
	return http.StatusInternalServerError, map[string]interface{}{
		"code":    "500",
		"message": msg,
	}
}

func ResponseNoContent(msg string) (int, map[string]interface{}) {
	return http.StatusNoContent, map[string]interface{}{
		"code":    "204",
		"message": msg,
	}
}

func ResponseCreateSuccess(msg string) (int, map[string]interface{}) {
	return http.StatusCreated, map[string]interface{}{
		"code":    "201",
		"message": msg,
	}
}

func ResponseForbidden(msg string) (int, map[string]interface{}) {
	return http.StatusForbidden, map[string]interface{}{
		"code":    "403",
		"message": msg,
	}
}
