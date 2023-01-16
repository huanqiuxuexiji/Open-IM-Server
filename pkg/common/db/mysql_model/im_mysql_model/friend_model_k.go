package im_mysql_model

import (
	"Open_IM/pkg/common/trace_log"
	"Open_IM/pkg/utils"
	"context"
	"gorm.io/gorm"
	"time"
)

type Friend struct {
	OwnerUserID    string    `gorm:"column:owner_user_id;primary_key;size:64"`
	FriendUserID   string    `gorm:"column:friend_user_id;primary_key;size:64"`
	Remark         string    `gorm:"column:remark;size:255"`
	CreateTime     time.Time `gorm:"column:create_time"`
	AddSource      int32     `gorm:"column:add_source"`
	OperatorUserID string    `gorm:"column:operator_user_id;size:64"`
	Ex             string    `gorm:"column:ex;size:1024"`
	db             *gorm.DB  `gorm:"-"`
}

func NewFriend(db *gorm.DB) *Friend {
	return &Friend{db: db}
}

func (f *Friend) Create(ctx context.Context, friends []*Friend) (err error) {
	defer func() {
		trace_log.SetCtxDebug(ctx, utils.GetSelfFuncName(), err, "friends", friends)
	}()
	return utils.Wrap(f.db.Create(&friends).Error, "")
}

func (f *Friend) Delete(ctx context.Context, ownerUserID string, friendUserIDs string) (err error) {
	defer func() {
		trace_log.SetCtxDebug(ctx, utils.GetSelfFuncName(), err, "ownerUserID", ownerUserID, "friendUserIDs", friendUserIDs)
	}()
	err = utils.Wrap(f.db.Where("owner_user_id = ? and friend_user_id = ?", ownerUserID, friendUserIDs).Delete(&Friend{}).Error, "")
	return err
}

func (f *Friend) UpdateByMap(ctx context.Context, ownerUserID string, args map[string]interface{}) (err error) {
	defer func() {
		trace_log.SetCtxDebug(ctx, utils.GetSelfFuncName(), err, "ownerUserID", ownerUserID, "args", args)
	}()
	return utils.Wrap(f.db.Where("owner_user_id = ?", ownerUserID).Updates(args).Error, "")
}

func (f *Friend) Update(ctx context.Context, friends []*Friend) (err error) {
	defer func() {
		trace_log.SetCtxDebug(ctx, utils.GetSelfFuncName(), err, "friends", friends)
	}()
	return utils.Wrap(f.db.Updates(&friends).Error, "")
}

func (f *Friend) UpdateRemark(ctx context.Context, ownerUserID, friendUserID, remark string) (err error) {
	defer func() {
		trace_log.SetCtxDebug(ctx, utils.GetSelfFuncName(), err, "ownerUserID", ownerUserID, "friendUserID", friendUserID, "remark", remark)
	}()
	return utils.Wrap(f.db.Model(f).Where("owner_user_id = ? and friend_user_id = ?", ownerUserID, friendUserID).Update("remark", remark).Error, "")
}

func (f *Friend) Find(ctx context.Context, ownerUserID string) (friends []*Friend, err error) {
	defer func() {
		trace_log.SetCtxDebug(ctx, utils.GetSelfFuncName(), err, "ownerUserID", ownerUserID, "friends", friends)
	}()
	return friends, utils.Wrap(f.db.Where("owner_user_id = ?", ownerUserID).Find(&friends).Error, "")
}

func (f *Friend) Take(ctx context.Context, ownerUserID, friendUserID string) (friend *Friend, err error) {
	friend = &Friend{}
	defer trace_log.SetCtxDebug(ctx, utils.GetSelfFuncName(), err, "ownerUserID", ownerUserID, "friendUserID", friendUserID, "group", *friend)
	return friend, utils.Wrap(f.db.Where("owner_user_id = ? and friend_user_id", ownerUserID, friendUserID).Take(friend).Error, "")
}

func (f *Friend) FindUserState(ctx context.Context, userID1, userID2 string) (friends []*Friend, err error) {
	defer func() {
		trace_log.SetCtxDebug(ctx, utils.GetSelfFuncName(), err, "userID1", userID1, "userID2", userID2)
	}()
	return friends, utils.Wrap(f.db.Where("(owner_user_id = ? and friend_user_id = ?) or (owner_user_id = ? and friend_user_id = ?)", userID1, userID2, userID2, userID1).Find(&friends).Error, "")
}
