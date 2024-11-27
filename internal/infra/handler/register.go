package handler

import (
	"github.com/jinfeijie/pangu/internal/constant"
	"github.com/jinfeijie/pangu/pkg"
)

func Register(confName string, handler pkg.Handler, sortId constant.SortType) {
	if Chain[sortId] == nil {
		Chain[sortId] = make(map[string]pkg.Handler)
	}
	Chain[sortId][confName] = handler
}
