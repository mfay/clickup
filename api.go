package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Task struct {
	ID                  string      `json:"id"`
	CustomID            interface{} `json:"custom_id"`
	CustomItemID        int         `json:"custom_item_id"`
	Name                string      `json:"name"`
	TextContent         string      `json:"text_content"`
	Description         string      `json:"description"`
	MarkdownDescription string      `json:"markdown_description"`
	Status              struct {
		ID         string `json:"id"`
		Status     string `json:"status"`
		Color      string `json:"color"`
		Orderindex int    `json:"orderindex"`
		Type       string `json:"type"`
	} `json:"status"`
	Orderindex  string      `json:"orderindex"`
	DateCreated string      `json:"date_created"`
	DateUpdated string      `json:"date_updated"`
	DateClosed  interface{} `json:"date_closed"`
	DateDone    interface{} `json:"date_done"`
	Archived    bool        `json:"archived"`
	Creator     struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Color          string `json:"color"`
		Email          string `json:"email"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"creator"`
	Assignees []struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Color          string `json:"color"`
		Initials       string `json:"initials"`
		Email          string `json:"email"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"assignees"`
	GroupAssignees []interface{} `json:"group_assignees"`
	Watchers       []struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Color          string `json:"color"`
		Initials       string `json:"initials"`
		Email          string `json:"email"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"watchers"`
	Checklists []interface{} `json:"checklists"`
	Tags       []interface{} `json:"tags"`
	Parent     interface{}   `json:"parent"`
	Priority   struct {
		Color      string `json:"color"`
		ID         string `json:"id"`
		Orderindex string `json:"orderindex"`
		Priority   string `json:"priority"`
	} `json:"priority"`
	DueDate      string      `json:"due_date"`
	StartDate    interface{} `json:"start_date"`
	Points       interface{} `json:"points"`
	TimeEstimate interface{} `json:"time_estimate"`
	TimeSpent    int         `json:"time_spent"`
	CustomFields []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		TypeConfig struct {
			Default     int         `json:"default"`
			Placeholder interface{} `json:"placeholder"`
			NewDropDown bool        `json:"new_drop_down"`
			Options     []struct {
				ID         string      `json:"id"`
				Name       string      `json:"name"`
				Color      interface{} `json:"color"`
				Orderindex int         `json:"orderindex"`
			} `json:"options"`
		} `json:"type_config"`
		DateCreated    string `json:"date_created"`
		HideFromGuests bool   `json:"hide_from_guests"`
		Required       bool   `json:"required"`
	} `json:"custom_fields"`
	Dependencies []interface{} `json:"dependencies"`
	LinkedTasks  []interface{} `json:"linked_tasks"`
	Locations    []interface{} `json:"locations"`
	TeamID       string        `json:"team_id"`
	URL          string        `json:"url"`
	Sharing      struct {
		Public               bool        `json:"public"`
		PublicShareExpiresOn interface{} `json:"public_share_expires_on"`
		PublicFields         []string    `json:"public_fields"`
		Token                interface{} `json:"token"`
		SeoOptimized         bool        `json:"seo_optimized"`
	} `json:"sharing"`
	PermissionLevel string `json:"permission_level"`
	List            struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"list"`
	Project struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
		Access bool   `json:"access"`
	} `json:"project"`
	Folder struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
		Access bool   `json:"access"`
	} `json:"folder"`
	Space struct {
		ID string `json:"id"`
	} `json:"space"`
	Subtasks []struct {
		ID           string `json:"id"`
		CustomItemID int    `json:"custom_item_id"`
		Name         string `json:"name"`
		Status       struct {
			Status     string `json:"status"`
			Orderindex int    `json:"orderindex"`
			Color      string `json:"color"`
			Type       string `json:"type"`
		} `json:"status"`
		Orderindex  string      `json:"orderindex"`
		DateCreated string      `json:"date_created"`
		DateUpdated string      `json:"date_updated"`
		DateClosed  interface{} `json:"date_closed"`
		DateDone    interface{} `json:"date_done"`
		Archived    bool        `json:"archived"`
		Creator     struct {
			ID             int    `json:"id"`
			Username       string `json:"username"`
			Color          string `json:"color"`
			Email          string `json:"email"`
			ProfilePicture string `json:"profilePicture"`
		} `json:"creator"`
		Assignees []struct {
			ID             int    `json:"id"`
			Username       string `json:"username"`
			Color          string `json:"color"`
			Initials       string `json:"initials"`
			Email          string `json:"email"`
			ProfilePicture string `json:"profilePicture"`
		} `json:"assignees"`
		GroupAssignees []interface{} `json:"group_assignees"`
		Watchers       []interface{} `json:"watchers"`
		Checklists     []interface{} `json:"checklists"`
		Tags           []interface{} `json:"tags"`
		Parent         string        `json:"parent"`
		DueDate        interface{}   `json:"due_date"`
		StartDate      interface{}   `json:"start_date"`
		Points         interface{}   `json:"points"`
		TimeEstimate   interface{}   `json:"time_estimate"`
		TimeSpent      int           `json:"time_spent"`
		CustomFields   []interface{} `json:"custom_fields"`
		Dependencies   []interface{} `json:"dependencies"`
		LinkedTasks    []interface{} `json:"linked_tasks"`
		Locations      []interface{} `json:"locations"`
		URL            string        `json:"url"`
	} `json:"subtasks"`
	Attachments []struct {
		ID               string      `json:"id"`
		Date             string      `json:"date"`
		Title            string      `json:"title"`
		Type             int         `json:"type"`
		Source           int         `json:"source"`
		Version          int         `json:"version"`
		Extension        string      `json:"extension"`
		ThumbnailSmall   string      `json:"thumbnail_small"`
		ThumbnailMedium  string      `json:"thumbnail_medium"`
		ThumbnailLarge   string      `json:"thumbnail_large"`
		IsFolder         interface{} `json:"is_folder"`
		Mimetype         string      `json:"mimetype"`
		Hidden           bool        `json:"hidden"`
		ParentID         string      `json:"parent_id"`
		Size             int         `json:"size"`
		TotalComments    int         `json:"total_comments"`
		ResolvedComments int         `json:"resolved_comments"`
		User             struct {
			ID             int    `json:"id"`
			Username       string `json:"username"`
			Email          string `json:"email"`
			Initials       string `json:"initials"`
			Color          string `json:"color"`
			ProfilePicture string `json:"profilePicture"`
		} `json:"user"`
		Deleted     bool        `json:"deleted"`
		Orientation interface{} `json:"orientation"`
		URL         string      `json:"url"`
		EmailData   struct {
			ID          string        `json:"id"`
			Msg         string        `json:"msg"`
			From        string        `json:"from"`
			Link        string        `json:"link"`
			Email       string        `json:"email"`
			Client      string        `json:"client"`
			Subject     string        `json:"subject"`
			Platform    string        `json:"platform"`
			AccountType string        `json:"accountType"`
			Attachments []interface{} `json:"attachments"`
		} `json:"email_data"`
		URLWQuery string `json:"url_w_query"`
		URLWHost  string `json:"url_w_host"`
	} `json:"attachments"`
}

type Api struct {
	ApiKey string
}

func NewApi(apiKey string) Api {
	return Api{ApiKey: apiKey}
}

func (a Api) GetTask(taskId string) Task {
	reqUrl := "https://api.clickup.com/api/v2/task/" + taskId
	req, err := http.NewRequest("GET", reqUrl, nil)

	query := req.URL.Query()
	query.Add("custom_task_ids", "true")
	query.Add("include_subtasks", "true")
	query.Add("include_markdown_description", "false")
	query.Add("custom_fields", "string")
	req.URL.RawQuery = query.Encode()

	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", a.ApiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var task Task
	json.Unmarshal(body, &task)
	return task
}


