// handlers/questions.go
package routes

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	quiz "main/quizzes"

	"github.com/gin-gonic/gin"
)

type CreateQuestionRequest struct {
    QuizID        int32  `json:"quiz_id"`
    QuestionText  string `json:"question_text"`
    OptionA       string `json:"option_a"`
    OptionB       string `json:"option_b"`
    OptionC       string `json:"option_c"`
    OptionD       string `json:"option_d"`
    CorrectOption string `json:"correct_option"`
}

type UpdateQuestionParams struct {
	QuestionText  string `json:"question_text"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string `json:"option_d"`
	CorrectOption string `json:"correct_option"`
	ID            int32  `json:"id"`
}

type GetRandomQuestionsRow struct {
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


func CreateQuestion(queries quiz.Querier) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req CreateQuestionRequest
        if err := c.BindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        question, err := queries.CreateQuestion(c, quiz.CreateQuestionParams{
            QuizID:        req.QuizID,
            QuestionText:  req.QuestionText,
            OptionA:       req.OptionA,
            OptionB:       req.OptionB,
            OptionC:       req.OptionC,
            OptionD:       req.OptionD,
            CorrectOption: req.CorrectOption,
        })
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, question)
    }
}

func UpdateQuestion(queries quiz.Querier) gin.HandlerFunc{
    return func(c *gin.Context){
        var req CreateQuestionRequest
        if err := c.BindJSON(&req); err!= nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        id , err:= strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
            return
        }

        question, err := queries.UpdateQuestion(c, quiz.UpdateQuestionParams{
            ID:            int32(id),
            QuestionText:  req.QuestionText,
            OptionA:       req.OptionA,
            OptionB:       req.OptionB,
            OptionC:       req.OptionC,
            OptionD:       req.OptionD,
            CorrectOption: req.CorrectOption,
        })
        if err!= nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, question)
    }
}


func GetRandomQuestions(queries quiz.Querier) (gin.HandlerFunc ){

    return func(c *gin.Context) {
     
    
        id , err:= strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
            return
        }


        randomQuestions, err := queries.GetRandomQuestions(c, int32(id))
        if err != nil {
            log.Fatalf("Failed to fetch random questions: %v", err)
        }

        c.JSON(http.StatusOK, randomQuestions)
    }

}