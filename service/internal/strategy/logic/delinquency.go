package logic

import "service/internal/strategy/utils/functions"

// Delinquency - структура для расчета внутренних агрегатов.
type Delinquency struct {
	// Признак достаточности денег для ежемесячного платежа в дату платежа в разбивке по месяцам
	EnoughMoney [6]bool
	// Признак просрочки платежа в разбивке по месяцам
	Delinquency [6]bool
	// Количество дней в просрочке в разбивке по месяцам
	DelinquencyDurationDays [6]int32
	// Сумма просрочек в разбивке по месяцам
	DelinquencySum [6]int32
}

// NewLocal - Конструктор структуры Delinquency.
func NewLocal() *Delinquency {
	return &Delinquency{}
}

// EnoughMoneyCount - расчет флага достаточности денег для ежемесячного платежа по месяцам.
func (d *Delinquency) EnoughMoneyCount(data *Data) {
	for i := 0; i < len(data.Request.Loans); i++ {
		if *data.Request.Accounts[i].PaymentDateBalance < data.Request.Loans[i].Payment {
			d.EnoughMoney[i] = false
		} else {
			d.EnoughMoney[i] = true
		}
	}
}

// DelinquencyCount - Расчет количества просрочек по месяцам.
func (d *Delinquency) DelinquencyCount(data *Data) {
	for i := 0; i < len(data.Request.Loans); i++ {
		if data.Request.Loans[i].ActualPaymentDate.After(data.Request.Loans[i].PaymentDate) {
			d.Delinquency[i] = true
		}
	}
}

// DelinquencyDurationCount - Расчет длительности просрочек в днях по месяцам.
func (d *Delinquency) DelinquencyDurationCount(data *Data) {
	for i := 0; i < len(data.Request.Loans); i++ {
		d.DelinquencyDurationDays[i] = functions.DateDiff(
			data.Request.Loans[i].PaymentDate,
			data.Request.Loans[i].ActualPaymentDate,
		)
	}
}

// DelinquencySumCount - Расчет суммы просрочки в разбивке по месяцам.
func (d *Delinquency) DelinquencySumCount(data *Data) {
	for index, value := range d.DelinquencyDurationDays {
		if value != 0 {
			d.DelinquencySum[index] =
				data.Request.Loans[index].Payment - *data.Request.Accounts[index].PaymentDateBalance
		}
	}
}
