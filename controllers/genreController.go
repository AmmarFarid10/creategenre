package controllers
import (
    "context"
    "log"
    "net/http"
    "shive-app/database"
    helper "shive-app/helpers"
    "shive-app/models"
    "time"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)
var genreCollection *mongo.Collection = database.OpenCollection(database.Client, "genre")
/*
func CreateGenre() gin.HandlerFunc {
    return func(c *gin.Context) {
        if err := helper.VerifyUserType(c, "ADMIN"); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var genre models.Genre
        defer cancel()
        if err := c.BindJSON(&genre); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "Status":  http.StatusBadRequest,
                "Message": "error",
                "Data":    map[string]interface{}{"data": err.Error()}})
            return
        }
        //Check to see if name exists
        regexMatch := bson.M{"$regex": primitive.Regex{Pattern: *genre.Name, Options: "i"}}
        count, err := genreCollection.CountDocuments(ctx, bson.M{"Name": regexMatch})
        defer cancel()
        if err != nil {
            log.Panic(err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "error occured while checking for the genre name"})
        }
        if count > 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "this genre name already exists", "count": count})
            return
        }
        if validationError := validate.Struct(&genre); validationError != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "Status":  http.StatusBadRequest,
                "Message": "error",
                "Data":    map[string]interface{}{"data": validationError.Error()}})
            return
        }
        newGenre := models.Genre{
            Id:   primitive.NewObjectID(),
            Name: genre.Name,
        }
        result, err := genreCollection.InsertOne(ctx, newGenre)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "Status":  http.StatusBadRequest,
                "Message": "error",
                "Data":    map[string]interface{}{"data": err.Error()}})
            return
        }
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "Status":  http.StatusInternalServerError,
                "Message": "error",
                "Data":    map[string]interface{}{"data": err.Error()}})
            return
        }
        c.JSON(http.StatusCreated, gin.H{
            "Status":  http.StatusCreated,
            "Message": "genre created successfully",
            "Data":    map[string]interface{}{"data": result}})
    }
}*/
func CreateGenre() gin.HandlerFunc {
    return func(c *gin.Context) {
        if err := helper.VerifyUserType(c, "ADMIN"); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var genre models.Genre
        defer cancel()
        if err := c.BindJSON(&genre); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "Status":  http.StatusBadRequest,
                "Message": "error",
                "Data":    map[string]interface{}{"data": err.Error()}})
            return
        }
        //Check to see if name exists
        regexMatch := bson.M{"$regex": primitive.Regex{Pattern: *genre.Name, Options: "i"}}
        count, err := genreCollection.CountDocuments(ctx, bson.M{"Name": regexMatch})
        defer cancel()
        if err != nil {
            log.Panic(err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "error occured while checking for the genre name"})
        }
        if count > 0 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "this genre name already exists", "count": count})
            return
        }
        if validationError := validate.Struct(&genre); validationError != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "Status":  http.StatusBadRequest,
                "Message": "error",
                "Data":    map[string]interface{}{"data": validationError.Error()}})
            return
        }
        newGenre := models.Genre{
            Id:         primitive.NewObjectID(),
            Name:       genre.Name,
            Created_at: time.Now(),
            Updated_at: time.Now(),
            Genre_id:   genre.Genre_id, // Insert genre_id
        }
        result, err := genreCollection.InsertOne(ctx, newGenre)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "Status":  http.StatusBadRequest,
                "Message": "error",
                "Data":    map[string]interface{}{"data": err.Error()}})
            return
        }
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "Status":  http.StatusInternalServerError,
                "Message": "error",
                "Data":    map[string]interface{}{"data": err.Error()}})
            return
        }
        c.JSON(http.StatusCreated, gin.H{
            "Status":  http.StatusCreated,
            "Message": "genre created successfully",
            "Data":    map[string]interface{}{"data": result}})
    }
}