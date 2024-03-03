package ultima_v_save

type IdToString[T comparable] struct {
	Id           T
	FriendlyName string
}

type OrderedMapping[T comparable] []IdToString[T]

//func (o *OrderedMapping[T]) GetIndexByFriendlyString(friendlyName string) int {
//
//	for i, val := range *o {
//		if val.FriendlyName == friendlyName {
//			return i
//		}
//	}
//	return -1
//}

func (o *OrderedMapping[T]) GetIndex(thing T) int {

	for i, val := range *o {
		if val.Id == thing {
			return i
		}
	}
	return -1
}

func (o *OrderedMapping[T]) GetById(thing T) *IdToString[T] {
	for _, val := range *o {
		if val.Id == thing {
			return &val
		}
	}

	return nil
}
