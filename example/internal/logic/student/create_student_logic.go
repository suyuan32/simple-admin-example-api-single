package student

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-api/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

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
		SetName(req.Name).
		SetAge(req.Age).
		SetAgeInt8(req.AgeInt8).
		SetAgeUint8(req.AgeUint8).
		SetAgeInt16(req.AgeInt16).
		SetAgeUint16(req.AgeUint16).
		SetAgeInt32(req.AgeInt32).
		SetAgeUint32(req.AgeUint32).
		SetAgeInt64(req.AgeInt64).
		SetAgeUint64(req.AgeUint64).
		SetAgeInt(req.AgeInt).
		SetAgeUint(req.AgeUint).
		SetWeightFloat(req.WeightFloat).
		SetWeightFloat32(req.WeightFloat32).
		SetClassID(uuidx.ParseUUIDString(req.ClassId)).
		SetEnrollAt(time.Unix(req.EnrollAt, 0)).
		SetStatusBool(req.StatusBool).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
