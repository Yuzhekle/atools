package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tealeg/xlsx"
)

func main() {
	// 打开Excel文件
	fileName := "dida_hotel_info_20250529115135.xlsx"
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		log.Fatalf("无法打开Excel文件: %v", err)
	}

	// 确保文件至少有一个sheet
	if len(xlFile.Sheets) == 0 {
		log.Fatalf("Excel文件中没有sheet")
	}

	// 获取第一个sheet (原始数据)
	sheet1 := xlFile.Sheets[0]
	fmt.Println("成功读取原始数据表格")

	// 创建新的sheet2用于存储按DidaBookingId过滤后的数据
	sheet2, err := xlFile.AddSheet("过滤后的订单")
	if err != nil {
		log.Fatalf("创建Sheet2失败: %v", err)
	}
	fmt.Println("已创建'过滤后的订单'表格")

	// 创建新的sheet3用于存储按DidaHotelId过滤后的数据
	sheet3, err := xlFile.AddSheet("唯一酒店ID")
	if err != nil {
		log.Fatalf("创建Sheet3失败: %v", err)
	}
	fmt.Println("已创建'唯一酒店ID'表格")

	// 复制表头到sheet2
	if len(sheet1.Rows) > 0 {
		headerRow := sheet2.AddRow()
		for _, cell := range sheet1.Rows[0].Cells {
			headerRow.AddCell().SetValue(cell.Value)
		}
	}

	// ==================== 过滤DidaBookingId阶段 ====================
	fmt.Println("开始过滤重复的DidaBookingId...")

	// 通过DidaBookingId进行过滤
	bookingIdMap := make(map[string]bool)
	var filteredByBookingId []*xlsx.Row

	// 查找DidaBookingId列的索引
	bookingIdIdx := -1
	if len(sheet1.Rows) > 0 {
		for i, cell := range sheet1.Rows[0].Cells {
			if cell.Value == "DidaBookingId" {
				bookingIdIdx = i
				fmt.Printf("找到DidaBookingId列，索引为: %d\n", bookingIdIdx)
				break
			}
		}
	}

	if bookingIdIdx == -1 {
		log.Fatalf("在表格中未找到DidaBookingId列")
	}

	// 遍历数据行（跳过表头）
	for i := 1; i < len(sheet1.Rows); i++ {
		row := sheet1.Rows[i]
		if len(row.Cells) <= bookingIdIdx {
			continue // 跳过没有足够列的行
		}

		bookingId := row.Cells[bookingIdIdx].Value
		if bookingId != "" && !bookingIdMap[bookingId] {
			bookingIdMap[bookingId] = true
			filteredByBookingId = append(filteredByBookingId, row)
		}
	}

	// 将按DidaBookingId过滤后的数据写入sheet2
	for _, row := range filteredByBookingId {
		newRow := sheet2.AddRow()
		for _, cell := range row.Cells {
			newRow.AddCell().SetValue(cell.Value)
		}
	}
	fmt.Printf("成功过滤DidaBookingId, 从%d行减少到%d行\n", len(sheet1.Rows)-1, len(filteredByBookingId))

	// ==================== 过滤DidaHotelId阶段 ====================
	fmt.Println("\n开始从过滤后的订单数据中过滤重复的DidaHotelId...")

	// 查找DidaHotelId列的索引
	hotelIdIdx := -1
	if len(sheet1.Rows) > 0 {
		for i, cell := range sheet1.Rows[0].Cells {
			if cell.Value == "DidaHotelId" {
				hotelIdIdx = i
				fmt.Printf("找到DidaHotelId列，索引为: %d\n", hotelIdIdx)
				break
			}
		}
	}

	if hotelIdIdx == -1 {
		log.Fatalf("在表格中未找到DidaHotelId列")
	}

	// 从过滤后的数据中再次过滤DidaHotelId
	hotelIdMap := make(map[string]bool)

	// 复制表头到sheet3
	if len(sheet1.Rows) > 0 {
		headerRow := sheet3.AddRow()
		for _, cell := range sheet1.Rows[0].Cells {
			headerRow.AddCell().SetValue(cell.Value)
		}
	}

	// 遍历过滤后的行（重要：是从filteredByBookingId中过滤，而不是原始数据）
	for _, row := range filteredByBookingId {
		if len(row.Cells) <= hotelIdIdx {
			continue // 跳过没有足够列的行
		}

		hotelId := row.Cells[hotelIdIdx].Value
		if hotelId != "" && !hotelIdMap[hotelId] {
			hotelIdMap[hotelId] = true

			// 将唯一的DidaHotelId数据写入sheet3
			newRow := sheet3.AddRow()
			for _, cell := range row.Cells {
				newRow.AddCell().SetValue(cell.Value)
			}
		}
	}
	fmt.Printf("成功过滤DidaHotelId, 从%d行减少到%d行\n", len(filteredByBookingId), len(hotelIdMap))

	// 统计结果
	fmt.Println("\n=== 处理结果统计 ===")
	fmt.Printf("原始数据: %d 行\n", len(sheet1.Rows)-1)
	fmt.Printf("按DidaBookingId过滤后: %d 行\n", len(filteredByBookingId))
	fmt.Printf("按DidaHotelId过滤后: %d 行\n", len(hotelIdMap))

	// 保存到新文件
	// 文件名 自动加上处理时间
	outputFileName := "dida_hotel_info_processed_" + time.Now().Format("20060102150405") + ".xlsx"
	err = xlFile.Save(outputFileName)
	if err != nil {
		log.Fatalf("保存Excel文件失败: %v", err)
	}
	fmt.Printf("\n处理完成，结果已保存到: %s\n", outputFileName)
}
