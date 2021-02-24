package chatting

import (
	"api/model/chatting"
	"api/model/user"
	"api/utlis/jwtParser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Rooms struct {
	ID int `json:"id"`
	About string `json:"about"`
	Name string `json:"name"`
	Participants []user.User `json:"participants"`
}

func CreateGroup(context *gin.Context) {
	var roomRequest Rooms
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	err := context.ShouldBindJSON(&roomRequest)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}

	userCount := len(roomRequest.Participants)

	if userCount == 0 {
		context.JSON(422, gin.H{"error": "No users were provided"})
		return
	}
	var newRoom chatting.Rooms
	var userArr []int
	var participantArr []chatting.RoomParticipants

	newRoom.Name = roomRequest.Name
	newRoom.About = roomRequest.About
	participantArr = append(participantArr, chatting.RoomParticipants{
		UserID: int(claims["id"].(float64)),
		IsAdmin: true,
		RoomsID: newRoom.ID,
		IsDeleted: false,
	})
	newRoom.IsDeleted = false

	chatting.SaveRoom(&newRoom)

	userArr = append(userArr, int(claims["id"].(float64)))

	for _, iter := range roomRequest.Participants {
		if findIfDuplicate(iter.ID ,userArr) {
			chatting.DeleteRoom(newRoom)
			context.JSON(500, gin.H{"error": "User duplicates"})
			return
		}
		participantArr = append(participantArr, chatting.RoomParticipants{
			IsAdmin: false,
			UserID: iter.ID,
			RoomsID: newRoom.ID,
			IsDeleted: false,
		})
	}
	participantArr[0].RoomsID = newRoom.ID

	chatting.SaveParticipants(participantArr)

	context.JSON(200, gin.H{"message": "Group Created"})
	return
}

func GetGroups(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	rooms := chatting.GetRooms(claims["id"])

	context.JSON(200, gin.H{"rooms": rooms})
	return
}

func GetGroup(context *gin.Context) {
	var room chatting.Rooms
	err := context.ShouldBindJSON(&room)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}

	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	rooms := chatting.GetRoom(claims["id"], room.ID)
	arr:= chatting.GetUnViewedMessages(int(claims["id"].(float64)),rooms.ID)
	var idArr []int

	for _, iter := range arr {
		idArr = append(idArr, iter.ID)
	}
	chatting.UpdateViews(int(claims["id"].(float64)), idArr)

	context.JSON(200, gin.H{"room": rooms})
	return
}

func SaveMessageRegular(context *gin.Context) {
	var message chatting.RoomMessages
	err := context.ShouldBindJSON(&message)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}

	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	if message.Message == "" && message.RoomsID != 0 {
		context.JSON(422, gin.H{"error": "Fields were not filled"})
		return
	}
	message.UserID = int(claims["id"].(float64))
	message.Sent = time.Now()
	message, arr:= chatting.SaveMessage(message)
	chatting.UpdateAt(message.RoomsID)
	for _, iter := range arr {
		newView := chatting.MessageViews{
			UserID: iter.UserID,
			MessageID: message.ID,
			RoomsID: message.RoomsID,
		}
		chatting.SaveView(newView)
	}

	context.JSON(200, gin.H{"message": "message saved"})
}

func GetUnViewed(context *gin.Context) {
	var message chatting.RoomMessages
	err := context.ShouldBindJSON(&message)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}

	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	if message.RoomsID == 0 {
		context.JSON(422, gin.H{"error": "Fields were not filled"})
		return
	}

	arr:= chatting.GetUnViewedMessages(int(claims["id"].(float64)),message.RoomsID)
	var idArr []int

	for _, iter := range arr {
		idArr = append(idArr, iter.ID)
	}
	chatting.UpdateViews(int(claims["id"].(float64)), idArr)

	context.JSON(200, gin.H{"messages": arr})
}

func findIfDuplicate(id int,arr []int ) bool{
	for _,iter := range arr {
		if iter == id {
			return true
		}
	}
	return false
}

func ChangeName(context *gin.Context) {
	var request chatting.Rooms
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	user := chatting.GetMember(request.ID, claims["id"])
	if !(user.IsAdmin) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	if request.Name == "" {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Field is empty"})
		return
	}

	chatting.UpdateName(request)

	context.JSON(200, gin.H{"message": "Name was changed successfully"})
}

func ChangeAbout(context *gin.Context) {
	var request chatting.Rooms
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	user := chatting.GetMember(request.ID, claims["id"])
	if !(user.IsAdmin) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	if request.About == "" {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Field is empty"})
		return
	}

	chatting.UpdateAbout(request)

	context.JSON(200, gin.H{"message": "Name was changed successfully"})
}

func ChangeAvatar(context *gin.Context) {
	var rooms int
	rooms, _ = strconv.Atoi(context.PostForm("id"))

	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	user := chatting.GetMember(rooms, claims["id"])
	if !(user.IsAdmin) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	file, _ := context.FormFile("avatar")
	if file != nil {
		path := "/chatting/%s/"
		path = fmt.Sprintf(path, fmt.Sprint(rooms))
		if _, err := os.Stat("storage" + path); os.IsNotExist(err) {
			os.MkdirAll("storage"+path, os.ModeDir)
		}
		path = path + file.Filename
		err := context.SaveUploadedFile(file, "storage"+path)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the avatar"})
			return
		}
		var request chatting.Rooms
		request.ID = rooms
		request.Avatar = path
		chatting.UpdateAvatar(request)
		context.JSON(200, gin.H{"message": "Avatar was changed successfully"})
	}
	context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error getting the Avatar"})
	return
}

func GetNonMembers(context *gin.Context) {
	var request chatting.Rooms
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	user := chatting.GetNonMembers(request.ID)
	context.JSON(200,gin.H{"users": user})
}

type RequestInvitees struct {
	Users []user.User `json:"users"`
	RoomID int `json:"room_id"`
}

func AddUsers(context *gin.Context) {
	var users RequestInvitees
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	err := context.ShouldBindJSON(&users)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	user := chatting.GetMember(users.RoomID, claims["id"])
	if !(user.IsAdmin) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	var members []chatting.RoomParticipants

	for _, iter := range users.Users {
		var newMember chatting.RoomParticipants
		newMember.UserID = iter.ID
		newMember.IsAdmin = false
		newMember.RoomsID = users.RoomID
		members = append(members, newMember)
	}

	chatting.SaveParticipants(members)
	chatting.UpdateAt( users.RoomID)
	context.JSON(200, gin.H{"message": "Members Added"})
}

func AddParticipants(context *gin.Context) {
	var roomRequest Rooms
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	err := context.ShouldBindJSON(&roomRequest)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	
	user := chatting.GetMember(roomRequest.ID, claims["id"])
	if !(user.IsAdmin) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	userCount := len(roomRequest.Participants)

	if userCount == 0 {
		context.JSON(422, gin.H{"error": "No users were provided"})
		return
	}
	var participantArr []chatting.RoomParticipants

	for _, iter := range roomRequest.Participants {
		participantArr = append(participantArr, chatting.RoomParticipants{
			IsAdmin: false,
			UserID: iter.ID,
			RoomsID: roomRequest.ID,
			IsDeleted: false,
		})
	}

	chatting.SaveParticipants(participantArr)
	chatting.UpdateAt( roomRequest.ID)

	context.JSON(200, gin.H{"message": "Member added"})
	return
}