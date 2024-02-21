package student

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-example-api/internal/logic/student"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
)

// swagger:route post /student/list student GetStudentList
//
// Get student list | 获取Student列表
//
// Get student list | 获取Student列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StudentListReq
//
// Responses:
//  200: StudentListResp

func GetStudentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StudentListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := student.NewGetStudentListLogic(r.Context(), svcCtx)
		resp, err := l.GetStudentList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
