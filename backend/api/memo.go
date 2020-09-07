package api

import (
	"backend/model/auth"
	"backend/model/memo"
	"backend/proto/pb"
	"backend/util/appret"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMemo(c *gin.Context) *pb.TAppRet {
	year, _ := strconv.Atoi(c.PostForm("year"))
	mouth, _ := strconv.Atoi(c.PostForm("mouth"))

	userSession := auth.GetAuthSession(c)

	memoList := pb.NewTMemoList()
	memoObj, err := memo.FindMemoByMouth(userSession.Uid, int64(year), int64(mouth))
	if err != nil {
		return appret.MakeErrRet(http.StatusOK, "GetMemoDbError")
	}
	for _, memo := range memoObj {
		oneMemo := pb.NewTMemo()
		oneMemo.ID = memo.ID
		oneMemo.Text = memo.Text
		oneMemo.RemindTime = memo.RemindTime

		memoList.Memos = append(memoList.Memos, oneMemo)
	}
	return appret.MakeSuccessRet(http.StatusOK, memoList)
}

func CreateMemo(c *gin.Context) *pb.TAppRet {
	RemindTime, _ := strconv.Atoi(c.PostForm("remind_time"))
	text := c.PostForm("text")

	userSession := auth.GetAuthSession(c)

	id, err := memo.CreateMemo(userSession.Uid, text, int64(RemindTime))
	if err != nil {
		return appret.MakeErrRet(http.StatusOK, "CreateMemoDbError")
	}
	resp := pb.NewTMemoCreateRet()
	resp.ID = id

	return appret.MakeSuccessRet(http.StatusOK, resp)
}
