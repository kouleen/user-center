package utils

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/mileusna/useragent"
)

func ParseUA(ctx context.Context) (os, browser string) {
	uaStr, ok := metainfo.GetPersistentValue(ctx, "User-Agent")
	if !ok {
		return
	}
	ua := useragent.Parse(uaStr)
	return ua.OS, ua.Name
}

func GetClientIp(ctx context.Context) (ip string) {
	ip, ok := metainfo.GetPersistentValue(ctx, "X-Real-IP")
	if !ok {
		return
	}
	return ip
}
