package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-api/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentByIdLogic {
	return &GetStudentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentByIdLogic) GetStudentById(req *types.IDReq) (*types.StudentInfoResp, error) {
	data, err := l.svcCtx.DB.Student.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.StudentInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.StudentInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.Unix()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.Unix()),
			},
			Name:          &data.Name,
			Age:           &data.Age,
			AgeInt8:       &data.AgeInt8,
			AgeUint8:      &data.AgeUint8,
			AgeInt16:      &data.AgeInt16,
			AgeUint16:     &data.AgeUint16,
			AgeInt32:      &data.AgeInt32,
			AgeUint32:     &data.AgeUint32,
			AgeInt64:      &data.AgeInt64,
			AgeUint64:     &data.AgeUint64,
			AgeInt:        &data.AgeInt,
			AgeUint:       &data.AgeUint,
			WeightFloat:   &data.WeightFloat,
			WeightFloat32: &data.WeightFloat32,
			ClassId:       pointy.GetPointer(data.ClassID.String()),
			TeacherId:     &data.TeacherID,
			EnrollAt:      pointy.GetPointer(data.EnrollAt.UnixMilli()),
			StatusBool:    &data.StatusBool,
		},
	}, nil
}
