package meet_service

import (
    "encoding/json"

    "meeting/models"
    "meeting/pkg/gredis"
    "meeting/pkg/logging"
    "meeting/service/cache_service"
)

type Meet struct {
    ID          int
    Type        int
    Banner      string
    Name        string
    StartAt     time.Time
    Place       string
    Longitude   string
    Latitude    string
    Fee         int
    Person      int
    Content     string
    State       int

    PageNum  int
    PageSize int
}

func (a *Meet) Add() error {
    item := map[string]interface{}{
        "type":       a.Type,
        "banner":     a.Banner,
        "name":       a.Name,
        "start_at":   a.StartAt,
        "place":      a.Place,
        "longitude":  a.Longitude,
        "latitude":   a.Latitude,
        "fee":        a.Fee,
        "person":     a.Person,
        "content":    a.Content,
        "state":      a.State,
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
        "start_at":   a.StartAt,
        "place":      a.Place,
        "longitude":  a.Longitude,
        "latitude":   a.Latitude,
        "fee":        a.Fee,
        "person":     a.Person,
        "content":    a.Content,
        "state":      a.State,
    })
}

func (a *Meet) Get() (*models.Article, error) {
    var cacheArticle *models.Article

    cache := cache_service.Article{ID: a.ID}
    key := cache.GetArticleKey()
    if gredis.Exists(key) {
        data, err := gredis.Get(key)
        if err != nil {
            logging.Info(err)
        } else {
            json.Unmarshal(data, &cacheArticle)
            return cacheArticle, nil
        }
    }

    article, err := models.GetArticle(a.ID)
    if err != nil {
        return nil, err
    }

    gredis.Set(key, article, 3600)
    return article, nil
}

func (a *Meet) GetAll() ([]*models.Meet, error) {
    var (
        articles, cacheArticles []*models.Article
    )

    cache := cache_service.Article{
        TagID: a.TagID,
        State: a.State,

        PageNum:  a.PageNum,
        PageSize: a.PageSize,
    }
    key := cache.GetArticlesKey()
    if gredis.Exists(key) {
        data, err := gredis.Get(key)
        if err != nil {
            logging.Info(err)
        } else {
            json.Unmarshal(data, &cacheArticles)
            return cacheArticles, nil
        }
    }

    articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
    if err != nil {
        return nil, err
    }

    gredis.Set(key, articles, 3600)
    return articles, nil
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
