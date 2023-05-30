package messagebox

// 创建提问箱
func MessageBoxCreate(messageBoxModel *MessageBox) error {
	err := CreateMessageBox(messageBoxModel)
	return err
}

// 根据ID获取提问箱
func MessageBoxGetByID(messageBoxID uint) (MessageBox, error) {
	messageBox, err := GetMessageBoxByID(messageBoxID)
	return messageBox, err
}

// 搜索提问箱
func MessageBoxSearch(title string, ownerID uint, pageNum int, pageSize int) ([]MessageBox, error) {
	offset := (pageNum - 1) * pageSize
	limit := pageSize
	messageBox, err := SearchMessageBox(title, ownerID, offset, limit)
	return messageBox, err
}

// 根据ID删除提问箱
func MessageBoxDeleteByID(messageBoxID uint, ownerID uint) error {
	err := DeleteMessageBoxByID(messageBoxID, ownerID)
	return err
}

// 根据ID更新提问箱
func MessageBoxUpdateByID(messageBoxID uint, ownerID uint, title string) error {
	err := UpdateMessageBoxByID(messageBoxID, ownerID, title)
	return err
}
