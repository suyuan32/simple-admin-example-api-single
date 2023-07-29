package student

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-api/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentLogic {
	return &UpdateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStudentLogic) UpdateStudent(req *types.StudentInfo) (*types.BaseMsgResp, error) {
	err := l.svcCtx.DB.Student.UpdateOneID(*req.Id).
		SetNotNilName(req.Name).
		SetNotNilAge(req.Age).
		SetNotNilAgeInt32(req.AgeInt32).
		SetNotNilAgeInt64(req.AgeInt64).
		SetNotNilAgeUint(req.AgeUint).
		SetNotNilAgeUint32(req.AgeUint32).
		SetNotNilAgeUint64(req.AgeUint64).
		SetNotNilWeightFloat(req.WeightFloat).
		SetNotNilWeightFloat32(req.WeightFloat32).
		SetNotNilClassID(uuidx.ParseUUIDStringToPointer(req.ClassId)).
		SetNotNilEnrollAt(pointy.GetTimePointer(req.EnrollAt, 0)).
		SetNotNilStatusBool(req.StatusBool).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
