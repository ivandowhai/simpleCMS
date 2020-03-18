package core

const (
	RoleAdmin     = uint8(1)
	RoleModerator = uint8(2)
	RoleAuthor    = uint8(3)
	//RoleCommenter = uint8(4)
)

var UserRolesCanPost = []uint8{RoleAdmin, RoleModerator, RoleAuthor}

func CanUserPost(role uint8) bool {
	for i := range UserRolesCanPost {
		if UserRolesCanPost[i] == role {
			return true
		}
	}

	return false
}

func IsAdmin(role uint8) bool {
	return role == RoleAdmin
}
