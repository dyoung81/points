package points

import (
	"context"
	"fmt"
	"time"
)

type PointsTransaction struct {
	ID        string
	player    string
	points    points
	timestamp time.Time
}

type points struct {
	points int
}

type Store interface {
	GetPointsTransactions(context.Context) ([]PointsTransaction, error)
	PostPointsTransaction(context.Context, PointsTransaction) (PointsTransaction, error)
	UpdatePointsTransaction(context.Context, string) (PointsTransaction, error)
	DeletePointsTransaction(context.Context, string) error
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetPointsTransactions(ctx context.Context) ([]PointsTransaction, error) {
	fmt.Println("retrieving points transactions")

	pttxs, err := s.Store.GetPointsTransactions(ctx)
	if err != nil {
		fmt.Println(err)
		return []PointsTransaction{}, err
	}

	return pttxs, nil
}

func (s *Service) PostPointsTransaction(ctx context.Context, pttx PointsTransaction) (PointsTransaction, error) {

	return PointsTransaction{}, nil
}

func (s *Service) UpdatePointsTransaction(ctx context.Context, id string) (PointsTransaction, error) {

	return PointsTransaction{}, nil
}

func (s *Service) DeletePointsTransaction(ctx context.Context, id string) error {

	return nil
}
