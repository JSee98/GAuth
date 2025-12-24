package types


/*
Please read if unsure https://argon2-cffi.readthedocs.io/en/stable/parameters.html
*/

type ARGON2_DEFAULT Argon2Params

var RFC_9106_LOW_MEMORY ARGON2_DEFAULT = ARGON2_DEFAULT{
    Memory: 64,
    Iterations: 3,
    Parallelism: 4,
    KeyLength: 128,
    SaltLength: 128,
}

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}
