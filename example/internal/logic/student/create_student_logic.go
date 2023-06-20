package student

import (
	"context"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-api/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStudentLogic {
	return &CreateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStudentLogic) CreateStudent(req *types.StudentInfo) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.Student.Create().
		SetNotNilName(req.Name).
		SetNotNilAge(req.Age).
		SetNotNilAgeInt8(req.AgeInt8).
		SetNotNilAgeUint8(req.AgeUint8).
		SetNotNilAgeInt16(req.AgeInt16).
		SetNotNilAgeUint16(req.AgeUint16).
		SetNotNilAgeInt32(req.AgeInt32).
		SetNotNilAgeUint32(req.AgeUint32).
		SetNotNilAgeInt64(req.AgeInt64).
		SetNotNilAgeUint64(req.AgeUint64).
		SetNotNilAgeInt(req.AgeInt).
		SetNotNilAgeUint(req.AgeUint).
		SetNotNilWeightFloat(req.WeightFloat).
		SetNotNilWeightFloat32(req.WeightFloat32).
		SetNotNilClassID(uuidx.ParseUUIDStringToPointer(req.ClassId)).
		SetNotNilTeacherID(req.TeacherId).
		SetNotNilEnrollAt(pointy.GetTimePointer(req.EnrollAt, 0)).
		SetNotNilStatusBool(req.StatusBool).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
