package redkacli

type Client interface {
}

type Value interface {
	Bool() (bool, error)
	Bytes() []byte
	Float() (float64, error)
	Int() (int, error)
	IsZero() bool
	MustBool() bool
	MustFloat() float64
	MustInt() int
	String() string
}

type StringCmd interface {
	Result() (string, error)
}
