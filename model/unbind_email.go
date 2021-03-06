package model

import (
	"fmt"

	"github.com/qiuchengw/go-user/db"
)

// 解绑邮箱认证.
//  调用该函数前, 请确认:
//  1. 该用户存在并且 verified
//  2. 该用户除了邮箱认证还有别的认证
func UnbindEmail(userId int64) (err error) {
	if err = removeFromCache(userId); err != nil {
		return
	}
	if err = unbindEmail(userId); err != nil {
		return
	}
	return syncToCache(userId)
}

func unbindEmail(userId int64) (err error) {
	para := struct {
		UserId      int64    `sqlx:"user_id"`
		NotBindType BindType `sqlx:"not_bind_type"`
	}{
		UserId:      userId,
		NotBindType: BindTypeMask &^ BindTypeEmail,
	}

	tx, err := db.GetDB().Beginx()
	if err != nil {
		return
	}

	// user_email 表删除一个 item
	stmt1, err := tx.Prepare("delete from user_email where user_id=? and verified=1")
	if err != nil {
		tx.Rollback()
		return
	}
	rslt1, err := stmt1.Exec(para.UserId)
	if err != nil {
		tx.Rollback()
		return
	}
	rowsAffected1, err := rslt1.RowsAffected()
	if err != nil {
		tx.Rollback()
		return
	}

	// user 更新 item
	stmt2, err := tx.PrepareNamed("update user set bind_types = bind_types&:not_bind_type where id=:user_id and verified=1 and bind_types&:not_bind_type<>0")
	if err != nil {
		tx.Rollback()
		return
	}
	rslt2, err := stmt2.Exec(para)
	if err != nil {
		tx.Rollback()
		return
	}
	rowsAffected2, err := rslt2.RowsAffected()
	if err != nil {
		tx.Rollback()
		return
	}

	if rowsAffected1 != rowsAffected2 {
		err = fmt.Errorf("用户 %d 解绑邮箱失败", para.UserId)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}
