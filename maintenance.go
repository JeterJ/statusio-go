package statusio

import (
	"errors"
	"fmt"
)

func (a StatusioApi) MaintenanceList(statuspage_id string) (r MaintenanceListResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("maintenance/list/%s", statuspage_id), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceMessage(statuspage_id string, message_id string) (r MaintenanceMessageResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("maintenance/message/%s/%s", statuspage_id, message_id), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceSchedule(task Maintenance) (r MaintenanceScheduleResponse, err error) {
	err = a.apiRequest("POST", "maintenance/schedule", task, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceStart(control MaintenanceContol) (r MaintenanceStartResponse, err error) {
	err = a.apiRequest("POST", "maintenance/start", control, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceUpdate(control MaintenanceContol) (r MaintenanceUpdateResponse, err error) {
	err = a.apiRequest("POST", "maintenance/update", control, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceFinish(control MaintenanceContol) (r MaintenanceFinishResponse, err error) {
	err = a.apiRequest("POST", "maintenance/finish", control, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceDelete(control MaintenanceContol) (r MaintenanceDeleteResponse, err error) {
	err = a.apiRequest("POST", "maintenance/delete", control, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}
