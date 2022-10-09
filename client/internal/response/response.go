// Package response - ответ сервиса.
package response

// Response - Структура ответа.
type Response struct {
	Version     string   `json:"version"`
	ExecuteDate string   `json:"execute_date"`
	Results     []Result `json:"resp"`
}

// Result - структура выходных агрегатов.
type Result struct {
	// Признак достаточности денег для ежемесячного платежа в дату платежа в разбивке по месяцам
	// Формат: 1 - достаточно , 0 - не достаточно
	EnoughMoneyByMonths [6]int32 `json:"enough_money_by_months"`
	// Признак просрочки платежа в разбивке по месяцам
	// Формат: 1 - есть просрочка, 0 - нет
	DelinquencyByMonths [6]int32 `json:"delinquency_by_months"`
	// Количество дней в просрочке за весь период
	DelinquencyDurationDaysTotal int32 `json:"delinquency_duration_days_total"`
	// Итоговая сумма просрочек за весь период
	DelinquencySumTotal int32 `json:"delinquency_sum_total"`
}
