package usecase

type ProfileUsecase struct{}

func (receiver ProfileUsecase) Register() {}
func (receiver ProfileUsecase) Delete()   {}
func (receiver ProfileUsecase) Modify()   {}
func (receiver ProfileUsecase) FindOne()  {}
func (receiver ProfileUsecase) FindAll()  {}
