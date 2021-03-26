package ulog

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

// 滚动文件输出, 此输出器应作为最终输出用
type RollingOutput struct {
	maxFileSize int
	fileName    string
	fileHandle  *os.File
}

func (self *RollingOutput) checkFile() {
	// 文件已经创建
	if self.fileHandle != nil {

		// 设定了文件大小约束
		if self.maxFileSize > 0 {
			info, err := self.fileHandle.Stat()
			if err != nil {
				return
			}

			// 文件还没到达约束
			if info.Size() < int64(self.maxFileSize) {
				return
			}
		} else { // 没有设置约束, 直接返回
			return
		}

	}

	// 找到可用的文件名
	fileName, fileIndex := self.foundUseableName()

	// 关闭之前的日志文件
	if self.fileHandle != nil {
		self.fileHandle.Close()
	}

	// 首次创建/尺寸超过后新创建
	self.fileHandle, _ = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	// 更新max文件的当前最大文件索引
	self.setFileIndex(fileIndex)
}

func (self *RollingOutput) Sync() error {
	if self.fileHandle == nil {
		return nil
	}

	return self.fileHandle.Sync()
}

func (self *RollingOutput) foundUseableName() (retFileName string, retFileIndex int) {

	fileIndex := self.getFileIndex()

	for i := fileIndex; ; i++ {

		retFileName = GetRollingFile(i, self.fileName)

		if _, err := os.Stat(retFileName); err != nil {
			if os.IsNotExist(err) {

				retFileIndex = i

				return
			}
		}
	}
}

// 根据当前日志名和文件索引, 获得实际文件名
func GetRollingFile(fileIndex int, fileName string) string {
	if fileIndex == 0 {
		return fileName
	} else {
		return fmt.Sprintf("%s.%d", fileName, fileIndex)
	}
}

func (self *RollingOutput) Write(b []byte) (n int, err error) {

	self.checkFile()

	if self.fileHandle == nil {
		return 0, nil
	}

	return self.fileHandle.Write(b)
}

func (self *RollingOutput) createDir() {
	// 自动创建日志目录
	logDir := filepath.Dir(self.fileName)

	_, err := os.Stat(logDir)
	if err != nil && os.IsNotExist(err) {
		os.MkdirAll(logDir, 0777)
	}
}

func (self *RollingOutput) getFileIndex() int {
	maxFile := fmt.Sprintf("%s.max", self.fileName)
	data, err := ioutil.ReadFile(maxFile)
	if err != nil {
		return 0
	}

	v, err := strconv.Atoi(string(data))
	if err != nil {
		return 0
	}

	return v
}

func (self *RollingOutput) setFileIndex(index int) {
	// 更新最大日志文件
	maxFile := fmt.Sprintf("%s.max", self.fileName)

	maxLogIndex := strconv.Itoa(index)

	ioutil.WriteFile(maxFile, []byte(maxLogIndex), 0666)
}

func NewRollingOutput(fileName string, maxFileSize int) *RollingOutput {
	self := &RollingOutput{
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}

	if maxFileSize == 0 {
		panic("maxFileSize should > 0")
	}

	self.createDir()

	return self
}
