package routes

import (
	"database/sql"
	quiz "main/quizzes"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)


type CreateQuizParams struct {
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
}

type GetQuizWithQuestionsRow struct {
	ID            int32          `json:"id"`
	Title         string         `json:"title"`
	Description   sql.NullString `json:"description"`
	CreatedAt     time.Time      `json:"created_at"`
	QuestionID    int32          `json:"question_id"`
	QuestionText  string         `json:"question_text"`
	OptionA       string         `json:"option_a"`
	OptionB       string         `json:"option_b"`
	OptionC       string         `json:"option_c"`
	OptionD       string         `json:"option_d"`
	CorrectOption string         `json:"correct_option"`
}


func CreateQuiz(queries quiz.Querier) gin.HandlerFunc{
    return func (c *gin.Context){

		var req CreateQuizParams

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
		}

		quiz , err := queries.CreateQuiz(c , quiz.CreateQuizParams{
			Title:       req.Title,
            Description: req.Description,
		})

		if err != nil {
		  c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create quiz."})
		}

		c.JSON(http.StatusCreated, quiz)



    }
}

func GetQuiz(queries quiz.Querier) gin.HandlerFunc{

	return func(c *gin.Context){
		id , err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		}

		quiz , err := queries.GetQuiz(c , int32(id))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Could not get quiz."})
		}

		c.JSON(http.StatusAccepted, quiz)


	}
}

func ListQuizzes(queries quiz.Querier) gin.HandlerFunc{
	return func (c *gin.Context){

		quizzes , err := queries.ListQuizzes(c)

		if err != nil {
		  c.JSON(http.StatusInternalServerError, gin.H{"error": "Error listing quizzes"})
		}

		c.JSON(http.StatusAccepted, quizzes)
		
	}
}

func GetQuizWithQuestions( queries quiz.Querier) gin.HandlerFunc{
	return func (c *gin.Context){

		id , err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		}

		quiz, err := queries.GetQuizWithQuestions(c , int32(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Could not get quiz."})
		}

		c.JSON(http.StatusAccepted, quiz)


	}
}