package base

type UserName struct{
  
  name string
  sex int
}

func getNew(name string,sex int)*UserName{
  return &UserName{"zhangming",1}
  }
