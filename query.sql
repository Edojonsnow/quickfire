-- query.sql
-- name: CreateQuiz :one
INSERT INTO quizzes (title, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetQuiz :one
SELECT * FROM quizzes
WHERE id = $1;

-- name: ListQuizzes :many
SELECT * FROM quizzes
ORDER BY created_at DESC;

-- name: CreateQuestion :one
INSERT INTO questions (quiz_id, question_text, option_a, option_b, option_c, option_d, correct_option)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetQuizWithQuestions :many
SELECT q.*, qs.id as question_id, qs.question_text, 
       qs.option_a, qs.option_b, qs.option_c, qs.option_d, qs.correct_option
FROM quizzes q
JOIN questions qs ON q.id = qs.quiz_id
WHERE q.id = $1;

-- name: UpdateQuestion :one
UPDATE questions
SET question_text = COALESCE($1, question_text),
    option_a = COALESCE($2, option_a),
    option_b = COALESCE($3, option_b),
    option_c = COALESCE($4, option_c),
    option_d = COALESCE($5, option_d),
    correct_option = COALESCE($6, correct_option)
WHERE id = $7
RETURNING *;


-- name: GetRandomQuestions :many
SELECT q.*, qs.id as question_id, qs.question_text, 
       qs.option_a, qs.option_b, qs.option_c, qs.option_d, qs.correct_option
FROM quizzes q
JOIN questions qs ON q.id = qs.quiz_id
WHERE q.id = $1
ORDER BY RANDOM()
LIMIT 5;