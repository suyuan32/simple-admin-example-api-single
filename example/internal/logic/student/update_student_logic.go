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
	err := l.svcCtx.DB.Student.UpdateOneID(req.Id).
		SetNotEmptyName(req.Name).
		SetNotEmptyAge(req.Age).
		SetNotEmptyAgeInt8(req.AgeInt8).
		SetNotEmptyAgeUint8(req.AgeUint8).
		SetNotEmptyAgeInt16(req.AgeInt16).
		SetNotEmptyAgeUint16(req.AgeUint16).
		SetNotEmptyAgeInt32(req.AgeInt32).
		SetNotEmptyAgeUint32(req.AgeUint32).
		SetNotEmptyAgeInt64(req.AgeInt64).
		SetNotEmptyAgeUint64(req.AgeUint64).
		SetNotEmptyAgeInt(req.AgeInt).
		SetNotEmptyAgeUint(req.AgeUint).
		SetNotEmptyWeightFloat(req.WeightFloat).
		SetNotEmptyWeightFloat32(req.WeightFloat32).
		SetNotEmptyClassID(uuidx.ParseUUIDString(req.ClassId)).
		SetNotEmptyEnrollAt(time.Unix(req.EnrollAt, 0)).
		SetNotEmptyStatusBool(req.StatusBool).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: i18n.UpdateSuccess}, nil
}
