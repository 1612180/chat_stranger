package service

import (
	"github.com/1612180/chat_stranger/internal/dtos"
	"github.com/1612180/chat_stranger/internal/models"
	"github.com/1612180/chat_stranger/internal/repository"
)

type RoomService struct {
	roomRepo repository.RoomRepo
	userRepo repository.UserRepo
	msgRepo  repository.MessageRepo
}

func NewRoomService(roomRepo repository.RoomRepo, userRepo repository.UserRepo, msgRepo repository.MessageRepo) *RoomService {
	return &RoomService{
		roomRepo: roomRepo,
		userRepo: userRepo,
		msgRepo:  msgRepo,
	}
}

func (s *RoomService) FetchAll() ([]*dtos.RoomResponse, []error) {
	rooms, errs := s.roomRepo.FetchAll()
	if len(errs) != 0 {
		return nil, errs
	}

	var roomRess []*dtos.RoomResponse
	for _, room := range rooms {
		roomRess = append(roomRess, room.ToResponse())
	}

	return roomRess, nil
}

func (s *RoomService) Find(id int) (*dtos.RoomResponse, []error) {
	room, errs := s.roomRepo.Find(id)
	if len(errs) != 0 {
		return nil, errs
	}

	return room.ToResponse(), nil
}

func (s *RoomService) Create() (int, []error) {
	return s.roomRepo.Create()
}

func (s *RoomService) Delete(id int) []error {
	return s.roomRepo.Delete(id)
}

func (s *RoomService) FindEmpty() (int, []error) {
	id, errs := s.roomRepo.FindEmpty()
	if len(errs) != 0 {
		return id, errs
	}

	return id, errs
}

func (s *RoomService) Join(usedid, roomid int) []error {
	return s.roomRepo.Join(usedid, roomid)
}

func (s *RoomService) Leave(userid, roomid int) []error {
	return s.roomRepo.Leave(userid, roomid)
}

func (s *RoomService) SendLatestMsg(userid, roomid, latest int) (*dtos.MessageResponse, int, []error) {
	// make sure user in room when receive msg
	if errs := s.roomRepo.Check(userid, roomid); len(errs) != 0 {
		return nil, 0, errs
	}

	msg, newLatest, errs := s.msgRepo.FetchLatest(roomid, latest)
	if len(errs) != 0 {
		return nil, 0, errs
	}

	fromUser, errs := s.userRepo.Find(msg.FromUserID)
	if len(errs) != 0 {
		return nil, 0, errs
	}

	msgRes, errs := msg.ToResponse(fromUser.FullName)
	if len(errs) != 0 {
		return nil, 0, errs
	}

	return msgRes, newLatest, nil
}

func (s *RoomService) ReceiveMsg(msgReq *dtos.MessageRequest) []error {
	// make sure user in room when send msg
	if errs := s.roomRepo.Check(msgReq.FromUserID, msgReq.RoomID); len(errs) != 0 {
		return errs
	}

	msg := (&models.Message{}).FromRequest(msgReq)
	return s.msgRepo.Create(msg)
}
