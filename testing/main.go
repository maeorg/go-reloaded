package main

import "reloaded"

func main(){

    input := ""
    want := ""
	test(input, want)

	input = "1E (hex) files were added"
	want = "30 files were added"
	test(input, want)

	input = "It has been 10 (bin) years"
	want = "It has been 2 years"
	test(input, want)

	input = "Ready, set, go (up) !"
	want = "Ready, set, GO !"
	test(input, want)

	input = "I should stop SHOUTING (low)"
	want = "I should stop shouting"
	test(input, want)

	input = "Welcome to the Brooklyn bridge (cap)"
	want = "Welcome to the Brooklyn Bridge"
	test(input, want)

	input = "This is so exciting (up, 2)"
	want = "This is SO EXCITING"
	test(input, want)

	input = "This Is sO EXCITING (low, 3)"
	want = "This is so exciting"
	test(input, want)

	input = "This Is sO EXCITING (low, 3) yeah let's go !"
	want = "This is so exciting yeah let's go!"
	test(input, want)

	input = "I was sitting over there ,and then BAMM !!"
	want = "I was sitting over there, and then BAMM!!"
	test(input, want)

	input = "I was thinking ... You were right"
	want = "I was thinking... You were right"
	test(input, want)

	input = "I am exactly how they describe me: ' awesome '"
	want = "I am exactly how they describe me: 'awesome'"
	test(input, want)

	input = "As Elton John said: ' I am the most well-known homosexual in the world '"
	want = "As Elton John said: 'I am the most well-known homosexual in the world'"
	test(input, want)

	input = "There it was. A amazing rock!"
	want = "There it was. An amazing rock!"
	test(input, want)

	input = "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	want = "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."
	test(input, want)

	input = "it (cap) was the best of times,"
	want = "It was the best of times,"
	test(input, want)

	input = "it was the worst of times (up) , it was the age of wisdom,"
	want = "it was the worst of TIMES, it was the age of wisdom,"
	test(input, want)
	
	input = "it was the age of foolishness (cap, 6) , it was the epoch of belief,"
	want = "It Was The Age Of Foolishness, it was the epoch of belief,"
	test(input, want)
	
	input = "it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope,"
	want = "it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope,"
	test(input, want)

	input = "IT WAS THE (low, 3) winter of despair."
	want = "it was the winter of despair."
	test(input, want)

	input = "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	want = "Simply add 66 and 2 and you will see the result is 68."
	test(input, want)

	input = "There is no greater agony than bearing a untold story inside you."
	want = "There is no greater agony than bearing an untold story inside you."
	test(input, want)

	input = "Punctuation tests are ... kinda boring ,don't you think !?"
	want = "Punctuation tests are... kinda boring, don't you think!?"
	test(input, want)

	input = "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?"
	want = "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"
	test(input, want)

	input = "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure"
	want = "I have to pack 5 outfits. Packed 26 just to be sure"
	test(input, want)

	input = "Don not be sad ,because sad backwards is das . And das not good"
	want = "Don not be sad, because sad backwards is das. And das not good"
	test(input, want)

	input = "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '"
	want = "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"
	test(input, want)

	input = ""
	want = ""
	test(input, want)

}

func test(input string, want string) {
	// output := reloaded.FormatText(input)
	temp := reloaded.FormatText(input)
	output := reloaded.FormatText(temp)

	if output != want {
		println("----------Failed!----------")
		println("Want: ")
		println(want)
		println("Got: ")
		println(output)
	} 
	// else {
	// 	println("Passed!")
	// }
}
