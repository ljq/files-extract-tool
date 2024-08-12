/**
* @Time : 2024/08/07 15:05
* @autor : Jack Liu
* @email: ljqlab@gmail.com
* @Software: VsCode
* @Description:
 */
package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/xuri/excelize/v2"
)

func readINIValue(filename string, key string) (string, error) {
	// Open the INI file
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open configuration file: %w", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore empty lines and comments
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}

		// Parse key-value pairs
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])

		// If the specified key is found, return its value
		if k == key {
			return v, nil
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading configuration file: %w", err)
	}

	// If the key was not found, return an error
	return "", fmt.Errorf("key not found: %s", key)
}

// Calculate the MD5 and SHA-256 values of the file
func calculateHashes(filePath string) (string, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	md5Hash := md5.New()
	sha256Hash := sha256.New()

	// Copy file contents to hash calculator
	if _, err := io.Copy(md5Hash, file); err != nil {
		return "", "", err
	}
	file.Seek(0, 0) // reset file pointer
	if _, err := io.Copy(sha256Hash, file); err != nil {
		return "", "", err
	}

	return hex.EncodeToString(md5Hash.Sum(nil)), hex.EncodeToString(sha256Hash.Sum(nil)), nil
}

// Read hash values from Excel files
func readExcelHashes(filePath string) (map[string]map[string]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	hashes := make(map[string]map[string]string)
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	for _, row := range rows[1:] { // jump header
		if len(row) < 3 {
			continue
		}
		fileName := row[0]
		md5Value := row[1]
		sha256Value := row[2]

		hashes[fileName] = map[string]string{
			"md5":     md5Value,
			"sha-256": sha256Value,
		}
	}

	return hashes, nil
}

func verifyFiles(dir string, hashes map[string]map[string]string) []string {
	var results []string

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			md5Value, sha256Value, err := calculateHashes(path)
			//fmt.Println("md5Value:" + md5Value)
			//fmt.Println("sha256Value:" + sha256Value)

			if err != nil {
				return err
			}
			//fmt.Println("hashes:", hashes)
			for fileName, hashValues := range hashes {
				if md5Value == hashValues["md5"] || sha256Value == hashValues["sha-256"] {
					//results = append(results, fmt.Sprintf("File[%s]: %s, MD5: %s, SHA-256: %s", fileName, path, md5Value, sha256Value))
					results = append(results, fmt.Sprintf("%s, %s, %s, %s", fileName, path, md5Value, sha256Value))
				}
			}
		}
		return nil
	})
	//fmt.Println("results", results)
	return results
}

// Write the results into a new Excel file
func writeResultsToExcel(results []string, outputFile string) error {
	f := excelize.NewFile()
	index, _ := f.NewSheet("Sheet1")

	// Set active sheet
	f.SetActiveSheet(index)

	// Set headers
	headers := map[string]string{
		"A1": "file_name[verifyed]",
		"B1": "file_path",
		"C1": "md5",
		"D1": "sha-256",
	}

	// Write headers to worksheet
	for cell, value := range headers {
		f.SetCellValue("Sheet1", cell, value)
	}

	// Write results to worksheet
	for i, result := range results {
		rowIndex := i + 2                     // Start writing from the second row (after headers)
		columns := strings.Split(result, ",") // Assuming result is comma-separated values

		// Assuming the order is fileName, md5, sha-256
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", rowIndex), columns[0])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", rowIndex), columns[1])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowIndex), columns[2])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", rowIndex), columns[3])
	}

	// Save the file
	err := f.SaveAs(outputFile)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) == 2 {
		// 获取命令行的第一个参数
		arg := os.Args[1]
		if arg == "-v" || arg == "-V" || arg == "--version" {
			arg = "The current version: 1.0.0"
			fmt.Println(arg)
			return
		}
	}

	dirToVerify, err := readINIValue("config.ini", "dirToVerify")
	if err != nil {
		fmt.Println("Failed to read INI value:", err)
		return
	}
	fmt.Println("Dir To Verify:", dirToVerify)

	// read source file: .xlsx
	excelFilePath := "hashes.xlsx"
	// output verifyed file: .xlsx
	outputExcelFile := "results.xlsx"
	hashes, err := readExcelHashes(excelFilePath)

	if err != nil {
		fmt.Println("Error reading Excel:", err)
		return
	}

	results := verifyFiles(dirToVerify, hashes)

	if err := writeResultsToExcel(results, outputExcelFile); err != nil {
		fmt.Println("Error writing results to Excel:", err)
		return
	}

	fmt.Println("Verification completed. Results written to", outputExcelFile)
}
