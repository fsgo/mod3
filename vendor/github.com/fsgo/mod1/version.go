/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/7/3
 */

package mod1

import (
	"runtime"
)


func Version()string{
	return "mod1{ mod_version:1.0.1 go:"+runtime.Version()+" }"
}
