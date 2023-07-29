package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-api/ent/predicate"
	"github.com/suyuan32/simple-admin-example-api/ent/student"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-api/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentListLogic {
	return &GetStudentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentListLogic) GetStudentList(req *types.StudentListReq) (*types.StudentListResp, error) {
	var predicates []predicate.Student
	if req.Name != nil {
		predicates = append(predicates, student.NameContains(*req.Name))
	}
	data, err := l.svcCtx.DB.Student.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.StudentListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.StudentInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.Unix()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.Unix()),
				},
				Name:          &v.Name,
				Age:           &v.Age,
				AgeInt32:      &v.AgeInt32,
				AgeInt64:      &v.AgeInt64,
				AgeUint:       &v.AgeUint,
				AgeUint32:     &v.AgeUint32,
				AgeUint64:     &v.AgeUint64,
				WeightFloat:   &v.WeightFloat,
				WeightFloat32: &v.WeightFloat32,
				ClassId:       pointy.GetPointer(v.ClassID.String()),
				EnrollAt:      pointy.GetPointer(v.EnrollAt.UnixMilli()),
				StatusBool:    &v.StatusBool,
			})
	}

	return resp, nil
}
