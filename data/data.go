package data



type Link struct{
	Id string
	Short string
	FullUrl string
	Active bool
}


type LinkStorer interface{
	SaveLink(link Link)error
	GetLinkById(id string)(*Link,error)
	GetLinkByShort(short string)(*Link,error)

}


func NewLink(full string)Link{
	id:=nanoid

}