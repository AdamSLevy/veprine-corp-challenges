package palindrome

import "testing"

var palindromes = []string{
	"Don't nod",
	"Dogma: I am God",
	"Never odd or even",
	"Too bad – I hid a boot",
	"Rats live on no evil star",
	"No trace; not one carton",
	"Was it Eliot's toilet I saw?",
	"Murder for a jar of red rum",
	"May a moody baby doom a yam?",
	"Go hang a salami; I'm a lasagna hog!",
	"Satan, oscillate my metallic sonatas!",
	"A Toyota! Race fast... safe car: a Toyota",
	"Straw? No, too stupid a fad; I put soot on warts",
	"Are we not drawn onward, we few, drawn onward to new era?",
	"Doc Note: I dissent. A fast never prevents a fatness. I diet on cod",
	"No, it never propagates if I set a gap or prevention",
	"Anne, I vote more cars race Rome to Vienna",
	"Sums are not set as a test on Erasmus",
	"Kay, a red nude, peeped under a yak",
	"Some men interpret nine memos",
	"Campus Motto: Bottoms up, Mac",
	"Go deliver a dare, vile dog!",
	"Madam, in Eden I'm Adam",
	"Oozy rat in a sanitary zoo",
	"Ah, Satan sees Natasha",
	"Lisa Bonet ate no basil",
	"Do geese see God?",
	"God saw I was dog",
	"Dennis sinned",
	"Never odd or even.",
}
var jabberwocky = []string{
	"’Twas brillig, and the slithy toves",
	"Did gyre and gimble in the wabe:",
	"All mimsy were the borogoves,",
	"And the mome raths outgrabe.",
	"“Beware the Jabberwock, my son!",
	"The jaws that bite, the claws that catch!",
	"Beware the Jubjub bird, and shun",
	"The frumious Bandersnatch!”",
	"He took his vorpal sword in hand;",
	"Long time the manxome foe he sought—",
	"So rested he by the Tumtum tree",
	"And stood awhile in thought.",
	"And, as in uffish thought he stood,",
	"The Jabberwock, with eyes of flame,",
	"Came whiffling through the tulgey wood,",
	"And burbled as it came!",
	"One, two! One, two! And through and through",
	"The vorpal blade went snicker-snack!",
	"He left it dead, and with its head",
	"He went galumphing back.",
	"“And hast thou slain the Jabberwock?",
	"Come to my arms, my beamish boy!",
	"O frabjous day! Callooh! Callay!”",
	"He chortled in his joy.",
	"’Twas brillig, and the slithy toves",
	"Did gyre and gimble in the wabe:",
	"All mimsy were the borogoves,",
	"And the mome raths outgrabe.",
}

func TestIs(t *testing.T) {
	for _, phrase := range palindromes {
		if !Is(phrase) {
			t.Fatalf("Is(%q) returned false but should be true", phrase)
		}
	}
	for _, line := range jabberwocky {
		if Is(line) {
			t.Fatalf("Is(%q) returned true but should be false", line)
		}
	}
}
