package main

func main() {
	println("Day 2: I Was Told There Would Be No Math")

	present_lenght, present_width, present_height := 1, 1, 10
	surface_area_of_box := SurfaceAreaOfBox(present_lenght, present_width, present_height)
	println("The surface area of the box is: ", surface_area_of_box)
	smallest_side_area := SmallestSideArea(present_lenght, present_width)
	println("The smallest side area of the box is: ", smallest_side_area)
	total_square_feet_of_wrapping_paper := surface_area_of_box + smallest_side_area
	println("The total square feet of wrapping paper needed is: ", total_square_feet_of_wrapping_paper)
}

func SurfaceAreaOfBox(length int, width int, height int) int {
	return (2 * length * width) + (2 * width * height) + (2 * height * length)
}

func SmallestSideArea(length int, width int) int {
	return length * width
}
