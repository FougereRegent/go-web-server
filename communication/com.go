package communication

import (
	"errors"
	"net"
)

type Protocol string

const (
	MESSAGE_SIZE = 128
)

const (
	TCP4 Protocol = "tcp4"
	TCP6 Protocol = "tcp6"
)

type Socket struct {
	Ln             net.Listener
	Port           uint
	Listen_adresse string
	Proto          Protocol
}

func InitSocket(host string, port uint, proto Protocol) (*Socket, error) {
	if port > 65635 {
		return nil, errors.New("Ports have to be lower than 65635")
	}

	var s Socket = Socket{
		Ln:             nil,
		Port:           port,
		Listen_adresse: host,
		Proto:          proto,
	}

	ln, err := net.Listen(string(proto), "0.0.0.0:8080")
	if err != nil {
		return nil, err
	}
	s.Ln = ln

	return &s, nil
}

func (s *Socket) Listen() (net.Conn, error) {
	con, err := s.Ln.Accept()

	if con == nil {
		return nil, err
	}
	return con, nil
}

func RecvMessage(con net.Conn) ([]byte, error) {
	result := make([]byte, MESSAGE_SIZE)
	total_byte := 0

	for {
		buffer := make([]byte, MESSAGE_SIZE)
		nb_byte, err := con.Read(buffer)

		if err != nil {
			return nil, err
		}

		total_byte += nb_byte

		if total_byte < MESSAGE_SIZE {
			copy(result, buffer)
			break
		} else if nb_byte < MESSAGE_SIZE {
			result = append(result, buffer...)
			break
		} else if total_byte == MESSAGE_SIZE {
			copy(result, buffer)
		} else {
			result = append(result, buffer...)
		}
	}
	return result, nil
}

func SendMessage(con net.Conn, message string) error {
	_, err := con.Write([]byte(message))
	if err != nil {
		return err
	}
	return nil
}

func CloseConnection(con net.Conn) {
	con.Close()
}
