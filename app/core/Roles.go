package core

const (
	RoleAdmin     = uint8(1)
	RoleModerator = uint8(2)
	RoleAuthor    = uint8(3)
	RoleCommenter = uint8(4)
)

type Role struct {
	ID    uint8
	Title string
}

var UserRolesCanPost = []uint8{RoleAdmin, RoleModerator, RoleAuthor}
var AllRoles = []Role{
	{ID: RoleAdmin, Title: "Admin"},
	{ID: RoleModerator, Title: "Moderator"},
	{ID: RoleAuthor, Title: "Author"},
	{ID: RoleCommenter, Title: "Commenter"},
}

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

func IsAdminOrModer(role uint8) bool {
	return role == RoleAdmin || role == RoleModerator
}
