package main

import "testing"

func TestMakeToken(t *testing.T) {
	MakeToken("while(a==b)begin\na:=a+1;#zhushi\nb:=b-1;\nc=c*d;\nd=c/d;\nif(a>b) then  c=C else C=c;\nend")
}

func TestMakeToken2(t *testing.T) {
	MakeToken("for(int i=1;i<=10;i++) begin\na_b++;#zszszszszs\nb_C--;#zszszszszs\nB123:=1234567;\na=@;\n123a=0;\na.b;\nend\n")
}

func TestMakeToken3(t *testing.T) {
	MakeToken("begin\nwhile(aacd<=100)\n\t\t\t\tbegin #hahahaha\n\t\t\t\t\tccdd=aacd-1;\n\t\t\t\t\tbbcc=ccdd+23;\n\t\t\t\t\t3c2b=ccdd<=4;\n\t\t\t\tend\nend")
}

func TestParseGrammer(t *testing.T) {
	//tokens := MakeToken("i*(i-i)/(i+i)*((i+i-i*i/i))")
	//tokens := MakeToken("i*i**")
	//tokens := MakeToken("i+i*i(")
	//tokens := MakeToken("i+i*i/i-i)")
	//tokens := MakeToken("i+)i-i(")
	tokens := MakeToken("(i-i)(i/i)")
	Grammar(tokens)
}
