package meet_service

import (
    "time"
    // "encoding/json"

    "meeting/models"
    // "meeting/pkg/gredis"
    // "meeting/pkg/logging"
    // "meeting/service/cache_service"
)

type Meet struct {
    ID          int
    Type        int
    Banner      string
    Name        string
    // StartAt     time.Time
    // Place       string
    // Longitude   string
    // Latitude    string
    // Fee         int
    // Person      int
    // Content     string
    // State       int

    PageNum  int
    PageSize int
}

func (a *Meet) Add() error {
    item := map[string]interface{}{
        "type":       a.Type,
        "banner":     a.Banner,
        "name":       a.Name,
        // "start_at":   a.StartAt,
        // "place":      a.Place,
        // "longitude":  a.Longitude,
        // "latitude":   a.Latitude,
        // "fee":        a.Fee,
        // "person":     a.Person,
        // "content":    a.Content,
        // "state":      a.State,
    }

    if err := models.AddMeet(item); err != nil {
        return err
    }

    return nil
}

func (a *Meet) Edit() error {
    return models.EditMeet(a.ID, map[string]interface{}{
        "type":       a.Type,
        "banner":     a.Banner,
        "name":       a.Name,
        // "start_at":   a.StartAt,
        // "place":      a.Place,
        // "longitude":  a.Longitude,
        // "latitude":   a.Latitude,
        // "fee":        a.Fee,
        // "person":     a.Person,
        // "content":    a.Content,
        // "state":      a.State,
    })
}

func (a *Meet) Get() (*models.Meet, error) {
    // var cacheArticle *models.Article

    // cache := cache_service.Article{ID: a.ID}
    // key := cache.GetArticleKey()
    // if gredis.Exists(key) {
    //     data, err := gredis.Get(key)
    //     if err != nil {
    //         logging.Info(err)
    //     } else {
    //         json.Unmarshal(data, &cacheArticle)
    //         return cacheArticle, nil
    //     }
    // }

    meet, err := models.GetMeet(a.ID)
    if err != nil {
        return nil, err
    }

    // gredis.Set(key, meet, 3600)
    return meet, nil
}

func (a *Meet) GetAll() ([]*models.Meet, error) {
    // var (
    //     articles, cacheArticles []*models.Meet
    // )

    // cache := cache_service.Article{
    //     TagID: a.TagID,
    //     State: a.State,

    //     PageNum:  a.PageNum,
    //     PageSize: a.PageSize,
    // }
    // key := cache.GetArticlesKey()
    // if gredis.Exists(key) {
    //     data, err := gredis.Get(key)
    //     if err != nil {
    //         logging.Info(err)
    //     } else {
    //         json.Unmarshal(data, &cacheArticles)
    //         return cacheArticles, nil
    //     }
    // }
    maps := make(map[string]interface{})
    maps["state"] = 1
    meets, err := models.GetMeets(a.PageNum, a.PageSize, maps)
    if err != nil {
        return nil, err
    }

    // gredis.Set(key, articles, 3600)
    return meets, nil
}

func (a *Meet) Delete() error {
    return models.DeleteMeet(a.ID)
}

func (a *Meet) ExistByID() (bool, error) {
    return models.ExistMeetByID(a.ID)
}

func (a *Meet) Count() (int, error) {
    return models.GetMeetTotal()
}
