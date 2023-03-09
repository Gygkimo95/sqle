// Package passport code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkpassport

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *PassportService {
	p := &PassportService{config: config}
	p.Session = &session{service: p}
	return p
}

type PassportService struct {
	config  *larkcore.Config
	Session *session // 登录态
}

type session struct {
	service *PassportService
}

// 批量获取用户登录信息（脱敏）
//
// - 该接口用于查询用户的登录信息
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/passport-v1/session/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/passportv1/query_session.go
func (s *session) Query(ctx context.Context, req *QuerySessionReq, options ...larkcore.RequestOptionFunc) (*QuerySessionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/passport/v1/sessions/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QuerySessionResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}