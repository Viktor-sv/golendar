package containers
/*
vector::assign
vector::at
vector::back
vector::begin
vector::capacity
vector::cbegin
vector::cend
vector::clear
vector::crbegin
vector::crend
vector::data
vector::emplace
vector::emplace_back
vector::empty
vector::end
vector::erase
vector::front
vector::get_allocator
vector::insert
vector::max_size
vector::operator=
vector::operator[]
vector::pop_back
vector::push_back
vector::rbegin
vector::rend
vector::reserve
vector::resize
vector::shrink_to_fit
vector::size
vector::swap
*/
type vector []int32

func (v vector) push_back(i int32){
	v = append(v, i)
}

func (v vector) size() int32{
	return int32(len(v))
}

func (v vector) del(){
	return
}

func (v vector) pop_back() int32{
	return v[len(v)-1]
}
