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
	"backend/model/auth"
	"backend/model/memo"
	"backend/proto/pb"
	"backend/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMemo(c *gin.Context, ret *pb.TAppRet) error {
	year, _ := strconv.Atoi(c.PostForm("year"))
	mouth, _ := strconv.Atoi(c.PostForm("mouth"))

	userSession := auth.GetAuthSession(c)

	memoList := pb.NewTMemoList()
	memoObj, err := memo.FindMemoByMouth(userSession.Uid, int64(year), int64(mouth))
	if err != nil {
		return util.MakeErrRet(ret, http.StatusOK, "GetMemoDbError")
	}
	for _, memo := range memoObj {
		oneMemo := pb.NewTMemo()
		oneMemo.ID = memo.ID
		oneMemo.Text = memo.Text
		oneMemo.RemindTime = memo.RemindTime

		memoList.Memos = append(memoList.Memos, oneMemo)
	}
	return util.MakeSuccessRet(ret, http.StatusOK, memoList)
}

func CreateMemo(c *gin.Context, ret *pb.TAppRet) error {
	RemindTime, _ := strconv.Atoi(c.PostForm("remind_time"))
	text := c.PostForm("text")

	userSession := auth.GetAuthSession(c)

	id, err := memo.CreateMemo(userSession.Uid, text, int64(RemindTime))
	if err != nil {
		return util.MakeErrRet(ret, http.StatusOK, "CreateMemoDbError")
	}
	resp := pb.NewTMemoCreateRet()
	resp.ID = id

	return util.MakeSuccessRet(ret, http.StatusOK, resp)
}
