package game

var avatarsAnim = &Anim{
	Key:       "avatars",
	Frames:    5,
	FrameSize: vec.I2{34, 64},
}

type avatar int

const (
	avatarNone     = avatar(iota - 1)
	avatarAwakeman // 0
	avatarSJ
	avatarDucky
	avatarAlamore
	avatarAwakemanOnPhone
)
