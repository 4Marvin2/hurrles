package repository

import (
	"errors"
	"fmt"
	"hurrles/config"
	"hurrles/internal/user/models"
	"log"

	"github.com/tarantool/go-tarantool"
)

const success = "Connection success (drip_tarantool) on: "

type ISessionRepository interface {
	GetSessionByCookie(sessionCookie string) (models.Session, error)
	NewSessionCookie(sessionCookie string, userId uint64) error
	DeleteSessionCookie(sessionCookie string) error
}

type sessionRepository struct {
	TarantoolConn *tarantool.Connection
}

func NewTarantoolConnection(tntConfig config.TarantoolConfig) (ISessionRepository, error) {
	addrPort := fmt.Sprintf("%s%s", tntConfig.Host, tntConfig.Port)

	conn, err := tarantool.Connect(addrPort, tarantool.Opts{
		User: "guest",
	})

	seesManager := sessionRepository{conn}

	if err != nil {
		return &sessionRepository{}, err
	} else {
		log.Printf("%s%s", success, addrPort)
	}

	resp, err := conn.Eval("return init()", []interface{}{})
	if err != nil {
		fmt.Println("Error", err)
		fmt.Println("Code", resp.Code)
		return &sessionRepository{}, err
	}

	return &seesManager, nil
}

func (conn *sessionRepository) GetSessionByCookie(sessionCookie string) (session models.Session, err error) {
	resp, err := conn.TarantoolConn.Call("check_session", []interface{}{sessionCookie})
	if err != nil {
		return models.Session{}, err
	}

	if len(resp.Data) == 0 {
		return models.Session{}, errors.New("not exixsts this cookie")
	}

	data := resp.Data[0]
	if data == nil {
		return models.Session{}, nil
	}

	sessionDataSlice, ok := data.([]interface{})
	if !ok {
		return models.Session{}, fmt.Errorf("cannot cast data: %v", sessionDataSlice)
	}

	if len(sessionDataSlice) == 0 {
		return models.Session{}, nil
	}

	cookie, ok := sessionDataSlice[0].(string)
	if !ok {
		return models.Session{}, fmt.Errorf("cannot cast data: %v", sessionDataSlice)
	}
	userId, ok := sessionDataSlice[1].(uint64)
	if !ok {
		return models.Session{}, fmt.Errorf("cannot cast data: %v", sessionDataSlice)
	}

	return models.Session{Cookie: cookie, UserID: userId}, nil
}

func (conn *sessionRepository) NewSessionCookie(sessionCookie string, id uint64) error {
	resp, err := conn.TarantoolConn.Call("new_session", []interface{}{sessionCookie, id})
	if err != nil {
		return err
	}

	if len(resp.Data) == 0 {
		return errors.New("this cookie already exists")
	}
	return nil
}

func (conn *sessionRepository) DeleteSessionCookie(sessionCookie string) error {
	resp, err := conn.TarantoolConn.Call("delete_session", []interface{}{sessionCookie})
	if err != nil {
		return err
	}

	if len(resp.Data) == 0 {
		return errors.New("this cookie is not exists")
	}

	return nil
}
