package entities

// Response - response struct.
type Response struct {
	Version     string   `json:"version"`
	ExecuteDate string   `json:"execute_date"`
	Results     []Result `json:"resp"`
}

// Result - output aggregates struct.
type Result struct {
	// A sign of the sufficiency of money for a monthly payment on the payment date by month
	// Format: 1 is enough , 0 is not enough
	// Признак достаточности денег для ежемесячного платежа в дату платежа в разбивке по месяцам
	// Формат: 1 - достаточно , 0 - не достаточно
	EnoughMoneyByMonths [6]int32 `json:"enough_money_by_months"`
	// Indication of late payment by month
	// Format: 1 - there is a delay, 0 - no
	// Признак просрочки платежа в разбивке по месяцам
	// Формат: 1 - есть просрочка, 0 - нет
	DelinquencyByMonths [6]int32 `json:"delinquency_by_months"`
	// Number of days overdue for the entire period
	// Количество дней в просрочке за весь период
	DelinquencyDurationDaysTotal int32 `json:"delinquency_duration_days_total"`
	// The total amount of delays for the entire period
	// Итоговая сумма просрочек за весь период
	DelinquencySumTotal int32 `json:"delinquency_sum_total"`
}
