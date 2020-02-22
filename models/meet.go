package models

import (
    "time"
    "fmt"
    "github.com/jinzhu/gorm"
)

type Meet struct {
    gorm.Model
    Type      int       `json:"type"`
    Banner    string    `json:"banner"`
    Title     string    `json:"title"`
    StartAt   time.Time     `json:"start_at"`
    Place     string    `json:"place"`
    Fee       int       `json:"fee"`
    Person    int       `json:"person"`
    Content   string    `json:"content" gorm:"size:1000"`
    State     int       `json:"state"`
}

func Meet_ExistByID(id int) (bool, error) {
    var meet Meet
    err := db.Find(&meet, id).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return false, err
    }

    if meet.ID > 0 {
        return true, nil
    }
    return false, nil
}

func Meet_Count() (int, error) {
    var count int
    if err := db.Model(&Meet{}).Count(&count).Error; err != nil {
        return 0, err
    }

    return count, nil
}

func Meet_List(pageNum int, pageSize int, maps interface{}) ([]*Meet, error) {
    var meets []*Meet
    err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&meets).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }

    return meets, nil
}

func Meet_Detail(id int) (*Meet, error) {
    var meet Meet
    err := db.First(&meet, id).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }

    return &meet, nil
}

func Meet_Update(id int, data map[string]interface{}) error {
    if err := db.Model(&Meet{}).Where("id = ?", id).Updates(data).Error; err != nil {
        return err
    }

    return nil
}

func Meet_Add(data map[string]interface{}) error {
    meet := Meet{
        Type:       data["type"].(int),
        Banner:     data["banner"].(string),
        Title:      data["title"].(string),
        StartAt:    data["start_at"].(time.Time),
        Place:      data["place"].(string),
        Fee:        data["fee"].(int),
        Person:     data["person"].(int),
        Content:    data["content"].(string),
        State:      data["state"].(int),
    }
    fmt.Println(meet)
    if err := db.Create(&meet).Error; err != nil {
        return err
    }

    return nil
}

func Meet_Delete(id int) error {
    if err := db.Where("id = ?", id).Delete(&Meet{}).Error; err != nil {
        return err
    }

    return nil
}
