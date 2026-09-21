// Package printshop предоставляет функции для расчетов стоимости печати в типографии:
// вычисление стоимости страниц, общего заказа с учетом тиража, скидок и отчетов.
package printshop

import (
	"fmt"
)

// PageCost вычисляет базовую стоимость печати страниц для одного экземпляра.
// Возвращает ошибку, если количество страниц или цена за страницу отрицательные.
func PageCost(pages int, pricePerPage float64) (float64, error) {
	if pages <= 0 {
		return 0, fmt.Errorf("количество страниц должно быть больше нуля: %d", pages)
	}
	if pricePerPage < 0 {
		return 0, fmt.Errorf("цена за страницу не может быть отрицательной: %.2f", pricePerPage)
	}
	return float64(pages) * pricePerPage, nil
}

// OrderCost рассчитывает общую стоимость заказа с учетом количества копий (тиража).
func OrderCost(pages, copies int, pricePerPage float64) (float64, error) {
	if copies <= 0 {
		return 0, fmt.Errorf("тираж (количество копий) должен быть больше нуля: %d", copies)
	}

	singleCost, err := PageCost(pages, pricePerPage)
	if err != nil {
		return 0, err
	}

	return singleCost * float64(copies), nil
}

// ApplyBulkDiscount изменяет общую стоимость заказа с учетом оптовой скидки через указатель.
// Возвращает ошибку, если процент скидки некорректный или указатель равен nil.
func ApplyBulkDiscount(cost *float64, percent float64) error {
	if cost == nil {
		return fmt.Errorf("указатель на стоимость не может быть nil")
	}
	if percent < 0 || percent > 100 {
		return fmt.Errorf("процент скидки должен быть в диапазоне от 0 до 100: %.2f", percent)
	}

	discountAmount := *cost * (percent / 100.0)
	*cost = *cost - discountAmount
	return nil
}

// FormatPrintReport формирует итоговую строку отчета по заказу типографии.
func FormatPrintReport(orderID string, pages, copies int, cost float64) (string, error) {
	if orderID == "" {
		return "", fmt.Errorf("номер заказа не может быть пустым")
	}
	if pages <= 0 || copies <= 0 || cost < 0 {
		return "", fmt.Errorf("некорректные данные для формирования отчета")
	}

	return fmt.Sprintf("Заказ №: %s | Страниц: %d | Тираж: %d экз. | Итоговая стоимость: %.2f руб.", orderID, pages, copies, cost), nil
}
