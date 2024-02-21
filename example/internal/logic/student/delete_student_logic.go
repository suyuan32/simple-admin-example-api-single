package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-api/ent/student"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-api/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStudentLogic {
	return &DeleteStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStudentLogic) DeleteStudent(req *types.UUIDsReq) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.Student.Delete().Where(student.IDIn(uuidx.ParseUUIDSlice(req.Ids)...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
