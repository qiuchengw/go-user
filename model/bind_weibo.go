package model

import (
	"fmt"

	"github.com/aiyi/go-user/db"
)

// 绑定微博(一般在认证后进行操作).
//  调用该函数前, 请确认:
//  1. 该用户存在并且 has_fixed
//  2. 该用户未绑定微博
//  3. 该微博未绑定用户
func BindWeibo(userId int64, openid, nickname string) (err error) {
	para := struct {
		UserId   int64  `sqlx:"user_id"`
		OpenId   string `sqlx:"openid"`
		Nickname string `sqlx:"nickname"`
		AuthType int64  `sqlx:"auth_type"`
	}{
		UserId:   userId,
		OpenId:   openid,
		Nickname: nickname,
		AuthType: AuthTypeWeibo,
	}

	tx, err := db.GetDB().Beginx()
	if err != nil {
		return
	}

	// user_weibo 表增加一个 item
	stmt1, err := tx.Prepare("insert into user_weibo(user_id, nickname, openid, has_fixed) values(?, ?, ?, 1)")
	if err != nil {
		tx.Rollback()
		return
	}
	if _, err = stmt1.Exec(para.UserId, para.Nickname, para.OpenId); err != nil {
		tx.Rollback()
		return
	}

	// user 更新 item
	stmt2, err := tx.PrepareNamed("update user set auth_types = auth_types|:auth_type where id=:user_id and has_fixed=1 and auth_types&:auth_type=0")
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
		err = fmt.Errorf("绑定微博 %s 到用户 %d 失败", para.OpenId, para.UserId)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}