/**********************************************************************************************************************
 *
 * Copyright (c) 2010 babeltime.com, Inc. All Rights Reserved
 * $
 *
 **********************************************************************************************************************/

/**
 * @file $
 * @author $(liujianyong@babeltime.com)
 * @date $
 * @version $
 * @brief
 *
 **/
package api

import (
	"firstWeb/module/memo"
	"firstWeb/proto/pb"
	"firstWeb/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMemo(c *gin.Engine) {
	c.POST("/memoList", func(c *gin.Context) {
		year := c.PostForm("year")
		mouth := c.PostForm("mouth")

		sess := sessions.Default(c)
		userObj := sess.Get("user").(pb.TAuthInfo)

		memoList := pb.NewTMemoList()
		defer memoList.Put()
		memoList, err := memo.FindMemoByMouth(userObj.Uid, int(year), int8(mouth))
		if err != nil {
			memoList.SetMessage("FindError")
			c.ProtoBuf(http.StatusOK, memoList)
			return
		}

		memoList.set(authObj.GetPhoneNum())
		memoList.SetMail(authObj.GetMail())
		memoList.SetToken(token)
		memoList.SetUid(string(authObj.GetUid()))
		memoList.SetNickName(authObj.GetAccount())

		sess := sessions.Default(c)
		sess.Set("user", retAuth)
		err = sess.Save()

		c.ProtoBuf(http.StatusOK, retAuth)
		authObj.Release()
	})
}
