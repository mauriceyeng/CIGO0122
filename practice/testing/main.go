package main


var test_cases=[]struct{
	input string,
	expected string,
	description string,
}
{
	{
		"WWWWWWWWWWWWBWWWBBBWWWWWWWBB",
		"12WB3W3B7W2B",
		"Mixed: Repeating and non repeating"
	},
	{
		"XYZ",
		"XYZ",
		"All Non repeating"
	},
	{
		"AABBBCCCC",
		"2A3B4C",
		"string with no single character"
	},
	{
		"A",
		"A",
		"single char"
	},
	{
		"",
		"",
		"empty string"
	},
	{
		"  heee  heee  ",
		"2 h3e2 h3e2 ",
		"white spaces"
	},
	{
		"AaaaaBBeee",
		"A4a2B3e",
		"mixed"
	},
}

func main(){

}