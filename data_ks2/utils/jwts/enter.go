package jwts

type JwyPayLoad struct {
	NickName string `json:"nickName"`
	RoleID   int    `json:"roleID"`
	UserID   int    `json:"userID"`
}

var Secret []byte
