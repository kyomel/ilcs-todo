package task

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kyomel/ilcs-todo/internal/domain/task/model"
	"github.com/kyomel/ilcs-todo/internal/domain/task/repository/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	testTitle  = "Test Task"
	testDesc   = "Test Description"
	dateFormat = "2006-01-02"
	testStatus = "pending"
)

type TaskUseCaseTestSuite struct {
	suite.Suite
	mockRepo *mock.MockTaskRepository
	useCase  Usecase
	ctx      context.Context
}

func (suite *TaskUseCaseTestSuite) SetupTest() {
	suite.mockRepo = &mock.MockTaskRepository{}
	suite.useCase = NewTaskUseCase(suite.mockRepo, 2*time.Second)
	suite.ctx = context.Background()
}

func TestTaskUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseTestSuite))
}

func (suite *TaskUseCaseTestSuite) TestCreateTask() {
	// Create a future date for testing
	futureDate := time.Now().Add(24 * time.Hour)
	futureDateStr := futureDate.Format(dateFormat)

	req := &model.TaskRequest{
		Title:       testTitle,
		Description: testDesc,
		DueDate:     futureDateStr,
	}

	expectedTask := &model.Task{
		ID:          uuid.New(),
		Title:       testTitle,
		Description: testDesc,
		Status:      testStatus,
		DueDate:     futureDate.Truncate(24 * time.Hour),
	}

	suite.mockRepo.On("PostTask", suite.ctx, req).Return(expectedTask, nil)

	result, err := suite.useCase.PostTask(suite.ctx, req)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), expectedTask.ID, result.ID)
	assert.Equal(suite.T(), expectedTask.Title, result.Title)
	assert.Equal(suite.T(), expectedTask.Description, result.Description)
	assert.Equal(suite.T(), expectedTask.Status, result.Status)
}

func (suite *TaskUseCaseTestSuite) TestCreateTaskValidationError() {
	// Create a past date for testing
	pastDate := time.Now().Add(-24 * time.Hour)
	pastDateStr := pastDate.Format(dateFormat)

	req := &model.TaskRequest{
		Title:       testTitle,
		Description: testDesc,
		DueDate:     pastDateStr,
	}

	result, err := suite.useCase.PostTask(suite.ctx, req)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Contains(suite.T(), err.Error(), "due date cannot be in the past")
}

func (suite *TaskUseCaseTestSuite) TestGetTasks() {
	page := 1
	limit := 10
	status := testStatus
	search := "test"

	futureDate := time.Now().Add(24 * time.Hour)

	expectedTasks := []*model.Task{
		{
			ID:          uuid.New(),
			Title:       testTitle,
			Description: testDesc,
			Status:      testStatus,
			DueDate:     futureDate.Truncate(24 * time.Hour),
		},
	}

	suite.mockRepo.On("GetTasksPaginated", suite.ctx, page, limit, status, search).Return(expectedTasks, nil)
	suite.mockRepo.On("GetTotalTasksWithFilter", suite.ctx, status, search).Return(1, nil)

	result, err := suite.useCase.GetAllTasks(suite.ctx, page, limit, status, search)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), 1, len(result.Tasks))
	assert.Equal(suite.T(), 1, result.Pagination.TotalTasks)
}

func (suite *TaskUseCaseTestSuite) TestGetTaskByID() {
	taskID := uuid.New()
	futureDate := time.Now().Add(24 * time.Hour)

	expectedTask := &model.Task{
		ID:          taskID,
		Title:       testTitle,
		Description: testDesc,
		Status:      testStatus,
		DueDate:     futureDate.Truncate(24 * time.Hour),
	}

	suite.mockRepo.On("GetTaskByID", suite.ctx, taskID).Return(expectedTask, nil)

	result, err := suite.useCase.GetTaskByID(suite.ctx, taskID)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), expectedTask.ID, result.ID)
}

func (suite *TaskUseCaseTestSuite) TestUpdateTask() {
	taskID := uuid.New()
	futureDate := time.Now().Add(24 * time.Hour)
	futureDateStr := futureDate.Format(dateFormat)

	req := &model.TaskRequest{
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     futureDateStr,
	}

	expectedTask := &model.Task{
		ID:          taskID,
		Title:       req.Title,
		Description: req.Description,
		Status:      testStatus,
		DueDate:     futureDate.Truncate(24 * time.Hour),
	}

	suite.mockRepo.On("UpdateTask", suite.ctx, taskID, req).Return(expectedTask, nil)

	result, err := suite.useCase.UpdateTask(suite.ctx, taskID, req)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), expectedTask.ID, result.ID)
	assert.Equal(suite.T(), expectedTask.Title, result.Title)
}

func (suite *TaskUseCaseTestSuite) TestDeleteTask() {
	taskID := uuid.New()

	suite.mockRepo.On("DeleteTask", suite.ctx, taskID).Return(nil)

	err := suite.useCase.DeleteTask(suite.ctx, taskID)

	assert.NoError(suite.T(), err)
}
