package api

// you can add fields to structures, if it's necessary

type Mo struct {
	MoID int      `json:"mo_id"`
	User BodyUser `json:"user"`
}

type BodyUser struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
}

type Body struct {
	Data   BodyData `json:"DATA"`
	Status string   `json:"STATUS"`
}

type BodyData struct {
	RowsCount int   `json:"rows_count"`
	Rows      []*Mo `json:"rows"`
}

type SaveIndicatorInstanceField struct {
	AuthUserId           string `json:"auth_user_id"`
	PeriodStart          string `json:"period_start"`
	PeriodEnd            string `json:"period_end"`
	PeriodKey            string `json:"period_key"`
	IndicatorToMoId      string `json:"indicator_to_mo_id"`
	FieldName            string `json:"field_name"`
	FieldValue           string `json:"field_value"`
	ApplyToFuturePeriods bool   `json:"apply_to_future_periods"`
}

// MoPayload request payload base structure
// You can embed one in other request structures
type MoPayload struct {
	PeriodStart  string `json:"period_start"`
	PeriodEnd    string `json:"period_end"`
	PeriodKey    string `json:"period_key"`
	MoChatFilter bool   `json:"mo_chart_filter,omitempty"`
}

type SaveInterpretation struct {
	AuthUserId                    string `json:"auth_user_id"`
	IndicatorInterpretationID     string `json:"indicator_interpretation_id"`
	IndicatorInterpretationAreaID string `json:"indicator_interpretation_area_id"`
	Protected                     string `json:"protected"`
	Scale                         bool   `json:"scale"`
	Name                          string `json:"name"`
	FieldName                     string `json:"field_name"`
	FieldValue                    string `json:"field_value"`
}
