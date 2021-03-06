package message

import (
    "io"
)


type Message interface {
    Decode(io.Reader, int16) error
    Encode(io.Writer, int16) error
}

type TagBuffer = map[string]interface{}
