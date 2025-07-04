package employee



import "errors"

var (
	ErrUserCannotSave      = errors.New("error user can not save")
	ErrGetUsers            = errors.New("error get users")
	ErrDuplicateUser       = errors.New("user already exists")
	ErrSavingUser          = errors.New("error saving user")
	ErrUserCannotGet       = errors.New("error can not get user")
	ErrUserCannotFound     = errors.New("error can no found user")
	ErrGettingUserByEmail  = errors.New("error getting user by the email")
	ErrNotFoundUserByEmail = errors.New("error not found user by email")
	ErrUserCannotLogin     = errors.New("error user can not login")
	ErrValidationUser      = errors.New("error validation user")
	ErrInvalidJson         = errors.New("error invalid json")
)
