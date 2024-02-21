package student

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-example-api/internal/logic/student"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
)

// swagger:route post /student student GetStudentById
//
// Get student by ID | 通过ID获取Student
//
// Get student by ID | 通过ID获取Student
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UUIDReq
//
// Responses:
//  200: StudentInfoResp

func GetStudentByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UUIDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := student.NewGetStudentByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetStudentById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
