package ux

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
)

type Table struct {
	Headers           []interface{}
	GenerateAutoIndex bool
}

//func getRowsAsMap(data []interface{}, numOfColumn int) []map[string]string {
//	var mapWithRowValues []map[string]string
//
//	for _, object := range data {
//		var objectValue []interface{}
//		objectValue = append(objectValue, object)
//
//		// Iterate on every object, per property (or struct attributes)
//		for _, item := range objectValue {
//			// Based on the number of headers, for this table, iterate on every property
//			mapOfValues := make(map[string]string)
//			for i := 0; i < numOfColumn; i++ {
//				mapOfValues[reflect.TypeOf(item).Field(i).Name] = reflect.ValueOf(item).Field(i).
//					String()
//			}
//
//			mapWithRowValues = append(mapWithRowValues, mapOfValues)
//		}
//	}
//	return mapWithRowValues
//}

//func getRowNames(rows []map[string]string, numberOfColumns int) []string {
//	var keyNames []string
//
//	for _, item := range rows {
//		for key, _ := range item {
//			keyNames = append(keyNames, key)
//		}
//
//		if len(keyNames) == numberOfColumns {
//			break
//		}
//	}
//
//	return keyNames
//}

//func getRows(data []interface{}, numOfColumn int) []table.Row {
//	// [] of maps, which includes the name of each row and its value
//	rowsAsMaps := getRowsAsMap(data, numOfColumn)
//
//	// [] of strings, which includes the 'keys' or name of the properties that holds the values
//	//to put in a row.
//	rowNames := getRowNames(rowsAsMaps, numOfColumn)
//
//	var tableRowsToAppend []table.Row // Slices of different rowsAsMaps
//
//	for _, item := range rowsAsMaps {
//		var rowValues []string
//
//		for _, rowName := range rowNames {
//			rowValues = append(rowValues, item[rowName])
//		}
//
//		tableRowsToAppend = append(tableRowsToAppend, table.Row{rowValues})
//	}
//
//	return tableRowsToAppend
//}

func MakeTable(options Table) table.Writer {
	// Main structure of the table
	t := table.NewWriter()

	// Headers
	t.Style().Color.Header = text.Colors{text.BgHiCyan, text.FgBlack}
	t.AppendHeader(table.Row(options.Headers))
	t.AppendSeparator()

	// Style and other options. TODO: Make this vary depending on the format, or option selected.
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)

	return t
}
