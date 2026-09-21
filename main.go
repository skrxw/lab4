package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/google/uuid"

	"github.com/skrxw/lab4/pkg/printshop"
)

func main() {
	// Генерация уникального ID заказа с помощью стороннего пакета uuid
	orderUUID := uuid.New()
	color.Cyan("=== Типография: Обработка заказа (UUID: %s) ===", orderUUID.String())

	pages := 50
	copies := 500
	pricePerPage := 2.50

	// 1. Расчет стоимости страниц для одной копии
	singleCost, err := printshop.PageCost(pages, pricePerPage)
	if err != nil {
		log.Fatalf("Ошибка расчета стоимости страниц: %v", err)
	}
	fmt.Printf("Стоимость печати 1 экземпляра (%d стр. по %.2f руб.): %.2f руб.\n", pages, pricePerPage, singleCost)

	// 2. Расчет общей стоимости тиража
	totalCost, err := printshop.OrderCost(pages, copies, pricePerPage)
	if err != nil {
		log.Fatalf("Ошибка расчета стоимости заказа: %v", err)
	}
	fmt.Printf("Общая стоимость тиража (%d экз.): %.2f руб.\n", copies, totalCost)

	// 3. Применение функции с указателем (скидка за опт 15%)
	discountPercent := 15.0
	err = printshop.ApplyBulkDiscount(&totalCost, discountPercent)
	if err != nil {
		log.Fatalf("Ошибка применения скидки: %v", err)
	}

	color.Green("Стоимость с учетом оптовой скидки (%.1f%%): %.2f руб.", discountPercent, totalCost)
	// 4. Формирование отчета
	report, err := printshop.FormatPrintReport("ORD-2026-001", pages, copies, totalCost)
	if err != nil {
		log.Fatalf("Ошибка формирования отчета: %v", err)
	}

	color.Yellow("\n--- Отчет о печати ---")
	fmt.Println(report)
}
