package model

import (
	"fmt"

	"github.com/qiuchengw/go-user/db"
)

// 绑定邮箱(一般在认证后进行操作).
//  调用该函数前, 请确认:
//  1. 该用户存在并且 verified
//  2. 该用户未绑定邮箱
//  3. 该邮箱未绑定用户
func BindEmail(userId int64, email string) (err error) {
	if err = removeFromCache(userId); err != nil {
		return
	}
	if err = bindEmail(userId, email); err != nil {
		return
	}
	return syncToCache(userId)
}

func bindEmail(userId int64, email string) (err error) {
	para := struct {
		UserId   int64    `sqlx:"user_id"`
		Email    string   `sqlx:"email"`
		BindType BindType `sqlx:"bind_type"`
	}{
		UserId:   userId,
		Email:    email,
		BindType: BindTypeEmail,
	}

	tx, err := db.GetDB().Beginx()
	if err != nil {
		return
	}

	// user_email 表增加一个 item
	stmt1, err := tx.Prepare("insert into user_email(user_id, email, verified) values(?, ?, 1)")
	if err != nil {
		tx.Rollback()
		return
	}
	if _, err = stmt1.Exec(para.UserId, para.Email); err != nil {
		tx.Rollback()
		return
	}

	// user 更新 item
	stmt2, err := tx.PrepareNamed("update user set bind_types = bind_types|:bind_type where id=:user_id and verified=1 and bind_types&:bind_type=0")
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
	if rowsAffected2 != 1 {
		err = fmt.Errorf("绑定邮箱 %s 到用户 %d 失败", para.Email, para.UserId)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}
