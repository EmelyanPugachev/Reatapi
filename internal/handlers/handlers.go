package handlers

import (
	"awesomeProject1/internal/taskservice"
	"awesomeProject1/internal/web/tasks"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskservice.TaskService
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Taska{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	if taskRequest == nil || taskRequest.Task == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "task is required")
	}
	isDone := false
	if taskRequest.IsDone != nil {
		isDone = *taskRequest.IsDone
	}
	taskToCreate := taskservice.Taska{
		Task:   *taskRequest.Task,
		IsDone: isDone,
	}
	createdTask, err := h.service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func NewTaskHandler(s taskservice.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) PatchTask(c echo.Context) error {
	id := c.Param("id")
	var body taskservice.TaskaRequest
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	updatedTaska, err := h.service.UpdateTask(id, body.Task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}
	return c.JSON(http.StatusOK, updatedTaska)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}
	return c.NoContent(http.StatusNoContent)
}
