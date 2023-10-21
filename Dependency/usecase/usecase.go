package usecase

import "aqary/repository"

type Usecase struct {
	Repo repository.RepositoryMethods
}

type UsecaseMethods interface {
}

func InitUsecase(repo repository.RepositoryMethods) Usecase {
	return Usecase{
		Repo: repo,
	}
}
