package models

import (
    "github.com/jinzhu/gorm"
)

type Meet struct {
    gorm.Model

    UserID      int `gorm:"column:uid"`
    Type        int
    Banner      string
    Name        string
    StartAt     time.Time
    Place       string
    Longitude   string
    Latitude    string
    Fee         int
    Person      int
    Content     string `gorm:"size:1000"`
    State       int
}

func GetMeetTotal(maps interface{}) (int, error) {
    var count int
    if err := db.Model(&Meet{}).Where(maps).Count(&count).Error; err != nil {
        return 0, err
    }

    return count, nil
}

func GetMeets(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
    var meets []*Meet
    err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&meets).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }

    return articles, nil
}

func GetMeet(id int) (*Article, error) {
    var meet Meet
    err := db.First(&meet, id).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }

    return &meet, nil
}

func UpdateMeet(id int, data map[string]interface{}) error {
    if err := db.Model(&Meet{}).Where("id = ?", id).Updates(data).Error; err != nil {
        return err
    }

    return nil
}

func AddMeet(data map[string]interface{}) error {
    article := Article{
        Type:       data["type"].(int),
        Banner:     data["banner"].(string),
        Name:       data["name"].(string),
        StartAt:    data["start_at"].(time.Time),
        Place:      data["place"].(string),
        Longitude:  data["longitude"].(string),
        Latitude:   data["latitude"].(string),
        Fee:        data["fee"].(int),
        Person:     data["person"].(int),
        content:    data["content"].(string),
        State:      data["state"].(int),
    }
    if err := db.Create(&article).Error; err != nil {
        return err
    }

    return nil
}

func DeleteMeet(id int) error {
    if err := db.Delete(&Meet{ID: id}).Error; err != nil {
        return err
    }

    return nil
}
