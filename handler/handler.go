package handler
import (
  "time"
  "golang.org/x/crypto/bcrypt"
 "github.com/yasserhejlah/goVosProject/database"
 "github.com/yasserhejlah/goVosProject/model"
 "github.com/gofiber/fiber/v2"
 "github.com/google/uuid"
 "github.com/golang-jwt/jwt/v4"
)

 func CreateUser(c *fiber.Ctx) error {
      db := database.DB.Db
      user := new(model.User)
      err := c.BodyParser(user)
      hash, _ := HashPassword(user.Password)
      if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message":  "Something's wrong with your input", "data": err})
      }
      user.Password = hash
      err = db.Create(&user).Error
      if err != nil {
        return c.Status(500).JSON(fiber.Map{"status": "error", "message":  "Could not create user", "data": err})
      }
      return c.Status(201).JSON(fiber.Map{"status": "success", "message":  "User has created", "data": user})
 }
func GetAllUsers(c *fiber.Ctx) error {
        db := database.DB.Db
        var users []model.User
        db.Find(&users)
        if len(users) == 0 {
          return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
        }
 return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}

func GetSingleUser(c *fiber.Ctx) error {
        db := database.DB.Db
        id := c.Params("id")
        var user model.User
        db.Find(&user, "id = ?", id)
        if user.ID == uuid.Nil {
          return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
        }
        return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

func UpdateUser(c *fiber.Ctx) error {
        type updateUser struct {
          Username string `json:"username"`
        }
        db := database.DB.Db
        var user model.User
        id := c.Params("id")
        db.Find(&user, "id = ?", id)
        if user.ID == uuid.Nil {
          return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
        }
        var updateUserData updateUser
        err := c.BodyParser(&updateUserData)
        if err != nil {
          return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
        }
        user.Username = updateUserData.Username
        db.Save(&user)
        return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})
}
func DeleteUserByID(c *fiber.Ctx) error {
          db := database.DB.Db
          var user model.User
          id := c.Params("id")
          db.Find(&user, "id = ?", id)
          if user.ID == uuid.Nil {
            return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
          }
          err := db.Delete(&user, "id = ?", id).Error
          if err != nil {
            return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
          }
return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
func Login(c *fiber.Ctx) error {
  db := database.DB.Db
  usr := SigninData{}
  var user model.User
  err := c.BodyParser(&usr)
	email := usr.Email
	// password := c.Params("password")
 db.Find(&user, "email = ?", email)
  if user.ID == uuid.Nil {
    return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
  }
  user.Password = ""
	claims := jwt.MapClaims{
		"data": user,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("wgcBKnzlyvv7tGZk5Zmevz30eC2O8N"))
  
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t,"data":user,"message": "Log in Succefully"})
}
func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}
type SigninData struct {
  Email  string `json:"email" xml:"email" form:"email"`
  Password string `json:"password" xml:"password" form:"password"`
}
