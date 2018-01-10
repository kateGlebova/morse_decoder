package morse_decoder 
	
import (
	"testing"
	"io/ioutil"
	"strings"
	"gopkg.in/mgo.v2"
)

func read_test_cases_file(t *testing.T, file string) string {
	b_test, err := ioutil.ReadFile(file) 
    if err != nil {
        t.Errorf(err.Error())
    }

   	return string(b_test)
}

func prepare_test_cases(str_test_cases string) ([]string, []string){
   	var morse, text []string
   	for _, test_case := range strings.Split(str_test_cases, "\n") {
   		test_case_slice := strings.Split(test_case, ":") 
   		if len(test_case_slice) == 2 {
   			morse_tc, text_tc := test_case_slice[0], test_case_slice[1]
   			morse = append(morse, morse_tc)
   			text = append(text, text_tc)
   		}
   	}
   	return morse, text
}

func test_one(t *testing.T, encoded, expected string) {
	db, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Cannot dial MongoDB", err)
	}
	defer db.Close() 

	morse_code := NewMorseCode(db)
	decoded := morse_code.Decode(encoded)
	if decoded != expected {
		t.Errorf("got='%s'	, expected='%s'\n", decoded, expected	)
	}
}

func TestDecode(t *testing.T) {
	morse, text := prepare_test_cases(read_test_cases_file(t, "test.txt"))

	for i := 0; i < len(morse); i++ {
		test_one(t, morse[i], text[i])
	}
}
