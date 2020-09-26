# Changelog

## [Unreleased]

## [0.2.0]

### Changed

修正將程式碼分開至不同檔案導致的自動建置錯誤

## [0.1.7]

### Changed

- 分開程式碼，修改 fop 為單獨 library
- 新增一個 main 以便建置成執行檔

## [0.1.6] 

### Changed

- 修正部分指令無效問題
- 將指令與旗幟(command and flag)的設定分開
- 此為 issue 1 的補充修正

## [0.1.5] 

### Added

- 加入 makefile 檔案，將自動建置改為使用 makefile 產生

## [0.1.4] 

### Added

- 利用 CircleCI 自動發布版本(issue 4)

## [0.1.3] 

### Added

- 新增一個編碼使用 utf8 的文字檔避免錯誤

### Changed

- 將指令與旗幟(command and flag)字串改成 const(issue 1)
- 推測檔案如果是文字檔才進行處理(issue 2)
- 使用建置參數設置版本號(issue 3)

## [0.1.2] 

### Added

- 新增錯誤類別，處理路徑為資料夾的錯誤

## [0.1.1]

### Added

- 加入 CircleCI yml 檔案，可自動編譯執行檔(但不會發佈)，自動執行 go test 測試

### Changed

- 並將 checksum 的測試資料分成 windows 與其他

## [0.1.0]

### Changed

- 將程式改為可處理參數

## [0.0.3]

### Added

- 加入 checksum 指令

## [0.0.2]

### Added

- 加入可搭配 go test 指令的測試檔案

## [0.0.1]

### Added

- 簡單的 linecount 功能，不能處理輸入命令，只會處理 myfile.txt

