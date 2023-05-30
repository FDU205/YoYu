package wall

func WallCreate(wallModel *Wall) error {
	err := CreateWall(wallModel)
	return err
}

func GetWallByPage(date string, pageNum int, pageSize int) ([]Wall, error) {
	offset := (pageNum - 1) * pageSize
	limit := pageSize
	wall, err := GetWall(date, offset, limit)
	return wall, err
}
