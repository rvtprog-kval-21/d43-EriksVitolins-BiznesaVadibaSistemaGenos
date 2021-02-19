package chatting

import (
	"api/model/chatting"
	"api/model/user"
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Rooms struct {
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
