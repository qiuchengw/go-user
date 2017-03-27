// 提供生成 userid 的 api.
package userid

import "github.com/qiuchengw/go-user/util"

func GetId() (int64, error) {
	return util.Int64ID(), nil
}
