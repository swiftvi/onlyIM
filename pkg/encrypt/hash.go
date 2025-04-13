package encrypt
import(
	"golang.org/x/crypto/bcrypt"
)

func GenHashedPwd(pwd []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
}

func VerifyHashedPwd(hashedPwd, pwd []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPwd, pwd)
	return err == nil
}