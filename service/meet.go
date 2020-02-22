package service

import (
    "time"

    "meeting/models"
)

type Meet struct {
    ID       int
    Type     int
    Banner   string
    Title    string
    StartAt  time.Time
    Place    string
    Fee      int
    Person   int
    Content  string
    State    int

    PageNum  int
    PageSize int
}

func (m *Meet) List() ([]*models.Meet, error) {
    maps := make(map[string]interface{})
    maps["state"] = m.State
    meets, err := models.Meet_List(m.PageNum, m.PageSize, maps)
    if err != nil {
        return nil, err
    }
    return meets, nil
}

func (m *Meet) Add() error {
    item := map[string]interface{}{
        "type":       m.Type,
        "banner":     m.Banner,
        "title":      m.Title,
        "start_at":   m.StartAt,
        "place":      m.Place,
        "fee":        m.Fee,
        "person":     m.Person,
        "content":    m.Content,
        "state":      m.State,
    }

    if err := models.Meet_Add(item); err != nil {
        return err
    }

    return nil
}

func (m *Meet) Update() error {
    return models.Meet_Update(m.ID, map[string]interface{}{
        "type":       m.Type,
        "banner":     m.Banner,
        "title":      m.Title,
        "start_at":   m.StartAt,
        "place":      m.Place,
        "fee":        m.Fee,
        "person":     m.Person,
        "content":    m.Content,
        "state":      m.State,
    })
}

func (m *Meet) Detail() (*models.Meet, error) {
    meet, err := models.Meet_Detail(m.ID)
    if err != nil {
        return nil, err
    }
    return meet, nil
}


func (m *Meet) Delete() error {
    return models.Meet_Delete(m.ID)
}

func (m *Meet) ExistByID() (bool, error) {
    return models.Meet_ExistByID(m.ID)
}

func (_ *Meet) Count() (int, error) {
    return models.Meet_Count()
}
