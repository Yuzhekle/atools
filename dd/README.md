# 滴答数据处理工具

这是一个用于处理滴答酒店信息 Excel 数据的 Go 程序。

## 功能

1. 读取原始 Excel 文件 `dida_hotel_info_20250521095614.xlsx`
2. 通过 `DidaBookingId` 过滤相同订单的数据，将结果保存到 Sheet2
3. 通过过滤后的数据，再过滤相同 `DidaHotelId` 数据，将唯一的 `DidaHotelId` 数据保存到 Sheet3
4. 统计处理结果并保存到新文件 `dida_hotel_info_processed.xlsx`

## 使用方法

1. 确保 `dida_hotel_info_20250521095614.xlsx` 文件在程序同一目录下
2. 安装依赖：
   ```
   go mod tidy
   ```
3. 运行程序：
   ```
   go run main.go
   ```
4. 查看生成的 `dida_hotel_info_processed.xlsx` 文件

## 要求

- Go 1.16+
- github.com/tealeg/xlsx v1.0.5

## 输出示例

程序运行后会输出类似以下信息：

```
原始数据: XX 行
按DidaBookingId过滤后: XX 行
按DidaHotelId过滤后: XX 行
处理完成，结果已保存到: dida_hotel_info_processed.xlsx
```
