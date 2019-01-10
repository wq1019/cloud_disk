package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wq1019/cloud_disk/handler/middleware"
	"github.com/wq1019/cloud_disk/model"
	"github.com/wq1019/cloud_disk/pkg/bytesize"
	"github.com/wq1019/cloud_disk/service"
	"github.com/zm-dev/go-image_uploader/image_url"
	"net/http"
)

type meHandler struct {
	imageUrl image_url.URL
}

func (m *meHandler) Show(c *gin.Context) {
	uid := middleware.UserId(c)
	user, err := service.UserLoadAndRelated(c.Request.Context(), uid)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, convert2UserResp(user, m.imageUrl))
}

func convert2UserResp(user *model.User, imageUrl image_url.URL) map[string]interface{} {
	fmt.Println(user.UserInfo.AvatarHash, imageUrl.Generate(user.UserInfo.AvatarHash))
	return map[string]interface{}{
		"id":                user.Id,
		"name":              user.Name,
		"created_at":        user.CreatedAt,
		"updated_at":        user.UpdatedAt,
		"nickname":          user.UserInfo.Nickname,
		"avatar_hash":       user.UserInfo.AvatarHash,
		"avatar_url":        imageUrl.Generate(user.UserInfo.AvatarHash),
		"profile":           user.UserInfo.Profile,
		"email":             user.UserInfo.Email,
		"used_storage":      bytesize.ByteSize(user.UserInfo.UsedStorage),
		"group_name":        user.Group.Name,
		"max_allow_storage": bytesize.ByteSize(user.Group.MaxStorage),
		"is_allow_share":    user.Group.AllowShare,
	}
}

func NewMeHandler(imageUrl image_url.URL) *meHandler {
	return &meHandler{imageUrl: imageUrl}
}
