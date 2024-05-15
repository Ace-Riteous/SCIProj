package api

import (
	"github.com/gin-gonic/gin"
)

// BaseApi 基本结构
type BaseApi struct {
	Ctx    *gin.Context
	Errors error
}

func NewBaseApi() BaseApi {
	return BaseApi{}
}

func (m *BaseApi) SetError(err error) {
	m.Errors = err
}

// GetError 获取改对象存在的error
func (m *BaseApi) GetError() error {
	return m.Errors
}

type BuildRequestOption struct {
	Ctx     *gin.Context
	DTO     any
	BindUri bool //用于判断是否从uri取数据
	BindAll bool //所有获取渠道

}

// BuildRequest 构建请求
func (m *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {

	//整合错误
	var errResult error

	//绑定ctx
	m.Ctx = option.Ctx

	//绑定请求数据
	if option.DTO != nil {
		//只有dto非空才进行绑定
		//判断是否从uri获取数据
		if option.BindUri || option.BindAll {
			errResult = m.Ctx.ShouldBindUri(option.DTO)

		}

		if !option.BindUri || option.BindAll {
			errResult = m.Ctx.ShouldBind(option.DTO)
		}

		//判断是否有错误
		if errResult != nil {
			//开始解释这个错误到我们的提示语言
			//errResult = m.ParseValidateError(errResult.(validator.ValidationErrors), option.DTO) //弃用，不应该断言

			m.SetError(errResult)
			m.Fail(ResponseJson{
				Code: 400,
				Msg:  m.GetError().Error(),
			})
		}
	}

	return m
}

// Fail 内部处理请求返回(进一步封装)
func (m *BaseApi) Fail(resp ResponseJson) {
	Fail(m.Ctx, resp)
}
func (m *BaseApi) OK(resp ResponseJson) {
	OK(m.Ctx, resp)
}
func (m *BaseApi) ServerFail(resp ResponseJson) {
	ServerFail(m.Ctx, resp)
}
