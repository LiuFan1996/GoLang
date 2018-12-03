package main

func map_equal(x,y map[string] string) bool{
	if(len(x))!= len(y){
		return false
	}
	for k,xv :=range x{
		if yv,ok:=y[k];!ok||yv!=xv{
			return false
		}
	}
	return true
}

