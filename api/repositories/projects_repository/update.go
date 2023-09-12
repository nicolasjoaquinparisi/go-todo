package projects_repository

import (
	"errors"
	"go-todo/api/models"
	"go-todo/api/utils/requests/projects_requests"
	"go-todo/database"
	"reflect"
)

func Update(project *models.Project, updateData projects_requests.UpdateBodyStruct) (*models.Project, error) {
	updateMap := make(map[string]interface{})

	// get reflect value from updateData
	updateValue := reflect.ValueOf(updateData)

	if updateValue.Kind() != reflect.Struct {
		return nil, errors.New("updateData must be a struct")
	}

	for i := 0; i < updateValue.NumField(); i++ {
		fieldName := updateValue.Type().Field(i).Name
		fieldValue := updateValue.Field(i).Interface()

		// add key and value to map if value is not nil
		if !reflect.DeepEqual(fieldValue, reflect.Zero(updateValue.Type().Field(i).Type).Interface()) {
			updateMap[fieldName] = fieldValue
		}
	}

	if len(updateMap) > 0 {
		if err := database.Instance.Model(&project).Updates(updateMap).Error; err != nil {
			return nil, err
		}
	}

	return project, nil
}
