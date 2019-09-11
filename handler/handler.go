package handler

import (
	"net/http"
	"strconv"
  //"fmt"
	"github.com/labstack/echo"
  // c-color さんに依存しちゃってる...
  // なんとか再現性があるように作り変えることができないもんかね...?
  // でも, wifi なくても起動できるんだよなぁ
  // どこを参照しているんだろう..... 
  // "github.com/x-color/simple-webapp/model"
  // ? できたのか ?
  "github.com/monkukui/procon-qa/model"
  "fmt"
)

func AddTodo(c echo.Context) error {
  fmt.Println(c)
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}

	if todo.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

  todo.UID = uid
  model.CreateTodo(todo)
	return c.JSON(http.StatusCreated, todo)
}

func GetTodos(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	todos := model.FindTodos(&model.Todo{UID: uid})
	return c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.DeleteTodo(&model.Todo{ID: todoID, UID: uid}); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdateTodo(c echo.Context) error {
    uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

    todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

    todos := model.FindTodos(&model.Todo{ID: todoID, UID: uid})
    if len(todos) == 0 {
        return echo.ErrNotFound
    }
    todo := todos[0]
    todo.Completed = !todos[0].Completed
    if err := model.UpdateTodo(&todo); err != nil {
		return echo.ErrNotFound
	}

    return c.NoContent(http.StatusNoContent)
}


// 以下自作
// 質問を全取得する
func GetAllQuestions(c echo.Context) error {
	// todos := model.FindTodos(&model.Todo{UID: uid})
  // &model.Question{} とすることで, 条件なしで取得する <=> 全取得 となる
  uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}
  questions := model.FindQuestions(&model.Question{})
	return c.JSON(http.StatusOK, questions)
}

// 質問を 1 つ 取得する
func GetQuestion(c echo.Context) error {

  uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

  QuestionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

  fmt.Println(QuestionID)
  question := model.FindQuestions(&model.Question{ID: QuestionID})[0]
  fmt.Println(question)
  return c.JSON(http.StatusOK, question)
}

// 質問を投稿する
func PostQuestion(c echo.Context) error {
	question := new(model.Question)

  // question に 送信されてきたデータを bind している
	if err := c.Bind(question); err != nil {
		return err
	}

  // 妥当性判定
  // Title, Body が空欄ではないことをチェックする
	if question.Title == "" || question.Body == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

  question.UID = uid
  model.CreateQuestion(question)

	return c.JSON(http.StatusCreated, question)
}
