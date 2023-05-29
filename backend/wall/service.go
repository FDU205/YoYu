package wall

func WallCreate(wallModel *Wall) error {
	err := CreateWall(wallModel)
	return err
}

func GetWallByPage(date string, pageNum int, pageSize int) ([]Wall, error) {
	wall, err := GetWall(date, pageNum, pageSize)
	return wall, err
}
