package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ShyamSundhar1411/chime-space-go-backend/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewChimeUseCase(chimeRepository domain.ChimeRepository, timeout time.Duration) domain.ChimeUsecase {
	return &chimeUsecase{
		chimeRepository: chimeRepository,
		contextTimeout:  timeout,
	}
}

func (cu *chimeUsecase) CreateChime(c context.Context, request domain.ChimeCreateOrUpdateRequest) (*domain.ChimeWithAuthor, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	var chime domain.Chime
	userId, ok := c.Value("userId").(string)
	if !ok {
		return nil, fmt.Errorf("user id not found")
	}
	primitiveUserId, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	chime = domain.Chime{
		ID:           bson.NewObjectID(),
		ChimeTitle:   request.ChimeTitle,
		ChimeContent: request.ChimeContent,
		IsPrivate:    request.IsPrivate,
		Author:       primitiveUserId,
		CreatedAt:    bson.NewDateTimeFromTime(time.Now()),
	}

	return cu.chimeRepository.CreateChime(ctx, &chime)
}

func (cu *chimeUsecase) Fetch(c context.Context) ([]domain.ChimeWithAuthor, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chimeRepository.Fetch(ctx)
}

func (cu *chimeUsecase) GetById(c context.Context, id string) (*domain.ChimeWithAuthor, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chimeRepository.GetById(ctx, id)
}

func (cu *chimeUsecase) FetchChimeFromUser(c context.Context) ([]domain.ChimeWithAuthor, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chimeRepository.GetChimeFromUserId(ctx)
}

func (cu *chimeUsecase) UpdateChime(c context.Context,request domain.ChimeCreateOrUpdateRequest,id string)(*domain.ChimeWithAuthor,error){
	ctx, cancel := context.WithTimeout(c,cu.contextTimeout)
	defer cancel()
	return cu.chimeRepository.UpdateChime(ctx, request, id)
}