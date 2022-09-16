package GeeRPC

import "GeeRPC/codec"

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int        //MagicNumber marks this is a geerpc request
	CodecType   codec.Type //client may choose different Codec to encode body
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func ()
