package admin_v1

import (
    // "net/http"
    "fmt"
    "time"
    "github.com/unknwon/com"
    // "github.com/astaxie/beego/validation"
    "github.com/gin-gonic/gin"

    "meeting/pkg/app"
    "meeting/pkg/e"
    "meeting/pkg/setting"
    "meeting/pkg/util"
    "meeting/service"
)

func Meet_List(c *gin.Context) {
    response := app.Response{C: c}

    // 默认获取状态为1的
    state := 1
    if  arg := c.PostForm("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
    }
    meetServer := service.Meet{
        State: state,
        PageNum:  util.GetPage(c),
        PageSize: setting.AppSetting.PageSize,
    }

    total, err := meetServer.Count()
    if err != nil {
        fmt.Println(err)
        response.SendErr(0, "获取Count失败")
        return
    }

    meets, err := meetServer.List()
    if err != nil {
        fmt.Println(err)
        response.SendErr(0, "获取List失败")
        return
    }

    data := make(map[string]interface{})
    data["lists"] = meets
    data["total"] = total
    response.SendSucc(data)
}

func Meet_Detail(c *gin.Context) {
    response := app.Response{C: c}

    id := com.StrTo(c.PostForm("id")).MustInt()
    meetServer := service.Meet{ID: id}
    exists, err := meetServer.ExistByID()
    if err != nil {
        response.SendErr(e.ERROR_DB, "")
        return
    }
    if !exists {
        response.SendErr(e.DATA_NOT_EXIST, "")
        return
    }

    meet, err := meetServer.Detail()
    if err != nil {
        response.SendErr(e.ERROR_DB, "")
        return
    }
    response.SendSucc(meet)
}

type MeetFormValid struct {
    ID        int       `form:"id"`
    Type      int       `form:"type" valid:"Required;Min(1)"`
    Banner    string    `form:"banner" valid:"Required;MaxSize(255)"`
    Title     string    `form:"title" valid:"Required;MaxSize(255)"`
    StartAt   time.Time `form:"start_at"`
    Place     string    `form:"place" valid:"Required;MaxSize(255)"`
    Fee       int       `form:"fee"`
    Person    int       `form:"person" valid:"Range(0,200)"`
    Content   string    `form:"content"`
    State     int       `form:"state"`
}
func Meet_Add(c *gin.Context) {
    var (
        response = app.Response{C: c}
        form MeetFormValid
    )

    _, errCode := app.BindAndValid(c, &form)
    if errCode != e.SUCCESS {
        response.SendErr(0, "参数异常")
        return
    }

    meetService := service.Meet{
        Type:      form.Type,
        Banner:    form.Banner,
        Title:     form.Title,
        StartAt:   form.StartAt,
        Place:     form.Place,
        Fee:       form.Fee,
        Person:    form.Person,
        Content:   form.Content,
        State:     form.State,
    }
    fmt.Println(meetService)
    if err := meetService.Add(); err != nil {
        response.SendErr(e.ERROR_DB, "")
        return
    }
    response.SendSucc(nil)
}

func Meet_Update(c *gin.Context) {
    var (
        response = app.Response{C: c}
        form MeetFormValid
    )

    _, errCode := app.BindAndValid(c, &form)
    if errCode != e.SUCCESS {
        response.SendErr(0, "参数异常")
        return
    }

    meetServer := service.Meet{
        ID:        form.ID,
        Type:      form.Type,
        Banner:    form.Banner,
        Title:     form.Title,
        StartAt:   form.StartAt,
        Place:     form.Place,
        Fee:       form.Fee,
        Person:    form.Person,
        Content:   form.Content,
        State:     form.State,
    }
    exists, err := meetServer.ExistByID()
    if err != nil {
        response.SendErr(e.ERROR_DB, "")
        return
    }
    if !exists {
        response.SendErr(e.DATA_NOT_EXIST, "")
        return
    }

    err = meetServer.Update()
    if err != nil {
        response.SendErr(e.ERROR_DB, "")
        return
    }
    response.SendSucc(nil)
}


func Meet_Delete(c *gin.Context) {
    response := app.Response{C: c}
    id := com.StrTo(c.Param("id")).MustInt()
    meetServer := service.Meet{ID: id}
    exists, err := meetServer.ExistByID()
    if err != nil {
        response.SendErr(e.ERROR_DB, "")
        return
    }
    if !exists {
        response.SendErr(e.DATA_NOT_EXIST, "")
        return
    }

    err = meetServer.Delete()
    if err != nil {
        response.SendErr(e.ERROR_DB, "")
        return
    }
    response.SendSucc(nil)
}