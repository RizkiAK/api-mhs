package user

type User struct {
	Nim      int    `json:"nim" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInput struct {
	Nim      int    `json:"nim"`
	Password string `json:"password"`
}
type UserUri struct {
	Nim int `uri:"nim"`
}
