package service

import (
	"errors"

	"github.com/candelatorrez/northwind-app/internal/domain"
	"github.com/candelatorrez/northwind-app/internal/repository"
	"gorm.io/gorm"
)

var (
	ErrClientNotFound      = errors.New("client not found")
	ErrInvoiceNotFound     = errors.New("invoice not found")
	ErrInvoiceAlreadyPaid  = errors.New("invoice already paid")
	ErrInvalidClientStatus = errors.New("invalid client status")
	ErrInvalidActionType   = errors.New("invalid action type")
)

type CollectionActionService struct {
	actionRepo *repository.CollectionActionRepository
	clientRepo *repository.ClientRepository
}

func NewCollectionActionService(
	actionRepo *repository.CollectionActionRepository,
	clientRepo *repository.ClientRepository,
) *CollectionActionService {
	return &CollectionActionService{
		actionRepo: actionRepo,
		clientRepo: clientRepo,
	}
}

func (s *CollectionActionService) CreateAction(
	clientID uint,
	actionType domain.ActionType,
	notes string,
	performedBy string,
) (*domain.CollectionAction, error) {
	if !isValidActionType(actionType) {
		return nil, ErrInvalidActionType
	}

	if _, err := s.clientRepo.FindByID(clientID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClientNotFound
		}
		return nil, err
	}

	action := &domain.CollectionAction{
		ClientID:    clientID,
		Type:        actionType,
		Notes:       notes,
		PerformedBy: performedBy,
	}

	if err := s.actionRepo.Create(action); err != nil {
		return nil, err
	}

	return action, nil
}

func isValidActionType(actionType domain.ActionType) bool {
	switch actionType {
	case domain.ActionCall, domain.ActionEmail, domain.ActionNote:
		return true
	default:
		return false
	}
}
