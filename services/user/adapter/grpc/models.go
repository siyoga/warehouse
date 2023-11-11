package models

import (
	"time"
	"warehouseai/user/adapter/grpc/gen"
	"warehouseai/user/model"

	"github.com/gofrs/uuid"
)

func UserToProto(m *model.User) *gen.User {
	var ownedAi []string
	var favoriteAi []string

	for _, s := range m.OwnedAi {
		ownedAi = append(ownedAi, s.String())
	}

	for _, s := range m.FavoriteAi {
		favoriteAi = append(favoriteAi, s.String())
	}

	return &gen.User{
		Id:         m.ID.String(),
		Username:   m.Username,
		Firstname:  m.Firstname,
		Lastname:   m.Lastname,
		Picture:    m.Picture,
		Email:      m.Email,
		OwnedAi:    ownedAi,
		FavoriteAi: favoriteAi,
		ViaGoogle:  m.ViaGoogle,
		CreatedAt:  m.CreatedAt.String(),
		UpdatedAt:  m.UpdatedAt.String(),
	}
}

func ProtoToUser(m *gen.User) *model.User {
	createdAt, _ := time.Parse(time.RFC3339, m.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, m.UpdatedAt)
	var ownedAi []uuid.UUID
	var favoriteAi []uuid.UUID

	for _, s := range m.OwnedAi {
		ownedAi = append(ownedAi, uuid.FromStringOrNil(s))
	}

	for _, s := range m.FavoriteAi {
		favoriteAi = append(favoriteAi, uuid.FromStringOrNil(s))
	}

	return &model.User{
		ID:         uuid.FromStringOrNil(m.Id),
		Username:   m.Username,
		Firstname:  m.Firstname,
		Lastname:   m.Lastname,
		Picture:    m.Picture,
		Email:      m.Email,
		ViaGoogle:  m.ViaGoogle,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		OwnedAi:    ownedAi,
		FavoriteAi: favoriteAi,
	}
}