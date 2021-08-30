package controller

import (
	"net/http"
	"strconv"

	"fransimanuel/belajargolangrestfulapi/helper"
	"fransimanuel/belajargolangrestfulapi/model/web"
	"fransimanuel/belajargolangrestfulapi/service"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&categoryCreateRequest)
	// helper.PanicIfError(err)

	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
	// w.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
}

func (controller CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	CategoryUpdateRequest := web.CategoryUpdateRequest{}
	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&CategoryUpdateRequest)
	// helper.PanicIfError(err)
	helper.ReadFromRequestBody(r, &CategoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	CategoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), CategoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
	// w.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
}

func (controller CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	// w.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(w, webResponse)
}

func (controller CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryRespponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRespponse,
	}

	// w.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(w, webResponse)
}

func (controller CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryRespponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRespponses,
	}

	// w.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(w)
	// err := encoder.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(w, webResponse)
}
