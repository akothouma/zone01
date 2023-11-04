package piscine

func SplitWhiteSpaces(s string) []string {
  var requiredstring []string
  word:="" //empty string called word
  for _,ch:=range s{
	if ch == ' ' || ch == '\t' || ch== '\n'{
		if word != ""{ // checks if word is not empty
			requiredstring=append(requiredstring, word)//appends word to required string
		word= "" //clears word and makes it empty again
		}else {
		word +=string(ch) //prints word
		}
	}
}
	if word != " "{ //
		requiredstring=append(requiredstring, word)
	}	
	return requiredstring
}


