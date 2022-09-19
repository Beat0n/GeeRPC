package codec

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
)

type JsonCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *json.Decoder
	enc  *json.Encoder
}

func (j JsonCodec) Close() error {
	return j.conn.Close()
}

func (j JsonCodec) ReadHeader(header *Header) error {
	return j.dec.Decode(header)
}

func (j JsonCodec) ReadBody(body interface{}) error {
	return j.dec.Decode(body)
}

func (j JsonCodec) Write(header *Header, body interface{}) (err error) {
	defer func() {
		j.buf.Flush()
		if err != nil {
			j.Close()
		}
	}()

	if err := j.enc.Encode(header); err != nil {
		log.Println("rpc codec: gob error encoding header", err)
		return err
	}
	if err := j.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body", err)
		return err
	}
	return nil
}

//var _ Codec = (*JsonCodec)(nil)

func NewJsonCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &JsonCodec{
		conn,
		buf,
		json.NewDecoder(conn),
		json.NewEncoder(buf),
	}
}
