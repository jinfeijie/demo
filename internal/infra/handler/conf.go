package handler

import (
	"github.com/jinfeijie/pangu/internal/constant"
	"github.com/jinfeijie/pangu/pkg"
)

var Chain = make(map[constant.SortType]map[string]pkg.Handler)
