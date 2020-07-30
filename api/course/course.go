package course

import (
	"medrepo-server/model"
	"medrepo-server/api"

	"github.com/kataras/iris/v12"
	"gopkg.in/go-playground/validator.v9"
)

type NewParam struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func Create(ctx iris.Context) {
	data := createParam(ctx)
	if data == nil {
		api.Error(ctx, 80001, "create failed!", nil)
		return
	}

	res := model.DB().Create(data)
	if err := res.Error; err != nil {
		api.Error(ctx, 80001, "create failed!", err)
		return
	}

	api.Success(ctx, "create success!", nil)
}

func createParam(ctx iris.Context) *model.Course {
	params := NewParam{}
	if err := ctx.ReadJSON(&params); err != nil {
		api.Error(ctx, 80400, "params error!", err)
		return nil
	}

	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		api.Error(ctx, 80400, "params error!", err)
		return nil
	}

	return &model.Course {
		Name: params.Name,
		Description: params.Description,
	}
}

func Get(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if id == 0 || err != nil {
		api.Error(ctx, 70400, "params error", err)
		return
	}

	course := model.Course{}
	if err := model.DB().Find(&course, id).Error; err != nil {
		api.Error(ctx, 80003, "get course failed", err)
		return
	}

	var chapters []model.Chapter
	model.DB().Model(&course).Related(&chapters)
	course.Chapters = chapters

	api.Success(ctx, "success", course)
}

func AddChapter(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if id == 0 || err != nil {
		api.Error(ctx, 70400, "params error", err)
		return
	}

	course := model.Course{}
	if err := model.DB().Find(&course, id).Error; err != nil {
		api.Error(ctx, 80003, "course not found", err)
		return
	}

	var chapters []model.Chapter
	model.DB().Model(&course).Related(&chapters)

	chapter := addChapterParam(ctx)
	chapters = append(chapters, *chapter)
	
	course.Chapters = chapters
	model.DB().Save(&course)

	api.Success(ctx, "success", course)
}

type NewChapterParam struct {
	Name     string `json:"name"`
	CourseID uint `json:"course_id"`
}

func addChapterParam(ctx iris.Context) *model.Chapter {
	params := NewChapterParam{}
	if err := ctx.ReadJSON(&params); err != nil {
		api.Error(ctx, 80400, "params error!", err)
		return nil
	}

	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		api.Error(ctx, 80400, "params error!", err)
		return nil
	}

	return &model.Chapter {
		Name: params.Name,
		CourseID: params.CourseID,
	}
}