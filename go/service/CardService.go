package service

import (
	"web-backend/go/dao"
	"web-backend/go/entity"
)

func CreateCard(card *entity.Card) (err error) {
	if err = dao.SqlSession.Create(card).Error; err != nil {
		return err
	}
	return
}

func GetAllCards() (cardList []*entity.Card, err error) {
	if err := dao.SqlSession.Find(&cardList).Error; err != nil {
		return nil, err
	}
	return
}

func DeleteUserById(id string) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.Card{}).Error
	return
}

func GetCardById(id string) (card []*entity.Card, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(card).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateUser(card *entity.Card) (err error) {
	err = dao.SqlSession.Save(card).Error
	return
}
