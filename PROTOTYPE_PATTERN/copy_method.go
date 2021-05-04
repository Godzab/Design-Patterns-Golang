package main

type Address1 struct{
	StreetAddress, City, Country string
}

type Person1 struct{
	Name string
	Address *Address1
	Friends []string
}

func(ad *Address1) DeepCopy() *Address1{
	return &Address1{ad.StreetAddress, ad.City, ad.Country}
}

func(p *Person1) DeepCopy() *Person1{
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}


func main(){

}
