package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

/*
 * Copyright (c) 2009, Mathew Tinsley (tinsley@tinsology.net)
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *    * Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *    * Redistributions in binary form must reproduce the above copyright
 *      notice, this list of conditions and the following disclaimer in the
 *      documentation and/or other materials provided with the distribution.
 *    * Neither the name of the organization nor the
 *      names of its contributors may be used to endorse or promote products
 *      derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY MATHEW TINSLEY ''AS IS'' AND ANY
 * EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL <copyright holder> BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

// LoremIpsum represents the dogipsum generator
type LoremIpsum struct{}

// Constants for sentence generation
const (
	WordsPerSentenceAvg = 8
	WordsPerSentenceStd = 4
)

// List of possible dog-themed words
var words = []string{
	"doggo",
	"shibe",
	"shoob",
	"shoober",
	"doggorino",
	"shooberino",
	"long doggo",
	"length boy",
	"noodle horse",
	"long water shoob",
	"aqua doggo",
	"pupper",
	"yapper",
	"pupperino",
	"wrinkler",
	"puggorino",
	"puggo",
	"corgo",
	"porgo",
	"woofer",
	"long woofer",
	"sub woofer",
	"heckin angery woofer",
	"heckin good boys",
	"floofs",
	"fluffer",
	"waggy wags",
	"long bois",
	"clouds",
	"boofers",
	"smol",
	"big ol",
	"doge",
	"bork",
	"borkf",
	"mlem",
	"blep",
	"blop",
	"pats",
	"tungg",
	"snoot",
	"ruff",
	"borkdrive",
	"thicc",
	"boof",
	"h*ck",
	"heck",
	"heckin",
	"vvv",
	"heckin good boys and girls",
	"big ol pupper",
	"you are doing me a frighten",
	"doing me a frighten",
	"you are doing me the shock",
	"ur givin me a spook",
	"you are doin me a concern",
	"stop it fren",
	"maximum borkdrive",
	"very good spot",
	"adorable doggo",
	"what a nice floof",
	"the neighborhood pupper",
	"borking doggo",
	"many pats",
	"lotsa pats",
	"he made many woofs",
	"dat tungg tho",
	"smol borking doggo with a long snoot for pats",
	"most angery pupper I have ever seen",
	"wow such tempt",
	"much ruin diet",
	"wow very biscit",
	"very hand that feed shibe",
	"such treat",
	"very taste wow",
	"I am bekom fat",
	"extremely cuuuuuute",
	"very jealous pupper",
	"super chub",
	"fat boi",
	"fluff nugget",
	"snifferino",
	"derp doggo",
	"smol fren",
	"chonky woofer",
	"danger noodle (friend-shaped)",
	"heckhound",
	"tail wiggle machine",
	"snout booper",
	"sneaky blepper",
	"yapster",
	"wigglebutt",
	"snack detector",
	"zoomie launcher",
	"bean toes",
	"toe beans",
	"sniffer unit",
	"snuggle beast",
	"borkmaster 3000",
	"whistle pupper",
	"drama floof",
	"borkinator",
	"treat goblin",
	"soggy doggo",
	"smol bean",
	"lick wizard",
	"boop machine",
	"happy bark engine",
	"awoo",
	"urha",
	"roo",
	"chonkerino",
	"heckhound deluxe",
	"fluff missile",
	"bark beast",
	"snoot missile",
	"wiggly woofer",
	"smooch bandit",
	"heckin chomper",
	"sassy shibe",
	"derpy snuffler",
	"long snoot lord",
	"tail-thruster",
	"scritch-magnet",
	"honker floof",
	"danger floof",
	"lick cannon",
	"snort doggo",
	"waggy unit",
	"heckin cuteness overload",
	"treat-powered pupper",
	"couch potato doggo",
	"mighty yapper",
	"booper trooper",
	"spook floof",
	"joy noodle",
	"gib treat pls",
	"no thoughts just snoot",
	"emergency smol dog",
	"chonky snuggle tube",
	"hot doggo (not edible)",
	"majestic floof",
	"smol zoomer",
	"heckin wiggle pup",
	"floof dragon",
	"chub lord",
	"wiggle king",
	"queen of snoots",
	"frenly chonker",
	"sneaky pupperino",
	"floof patrol",
	"chonky floofster",
	"snuggle commander",
	"supreme borker",
	"heckhound junior",
	"mini floofball",
	"chonkinator",
	"mighty sniffer",
	"beeg fren",
	"fluff sergeant",
	"snoot officer",
	"danger wiggle",
	"micro woof",
	"mega woofer",
	"super smol fren",
	"heckin fluff factory",
	"smol bark engine",
	"snacc hound",
	"tater tot doggo",
	"carpet shark",
	"supreme snuggler",
	"mighty chomp champ",
	"snoot whisperer",
	"undercover doggo",
	"professional blepper",
	"boop-ready floof",
	"smol storm of chaos",
	"blessing of floof",
	"wag master",
	"tippy tap champion",
	"tail wag wizard",
	"unicorn doggo",
	"party pupper",
	"midnight zoomer",
	"chirpy yipper",
	"sneaky snack hunter",
	"certified good boye",
	"floofin menace",
	"bean sprint master",
	"airborne woofer",
	"grumble pup",
	"snuggle-powered unit",
}

// Generate creates "Lorem ipsum" style words with dog-themed content
func (l *LoremIpsum) Generate(numWords int) string {
	if numWords <= 0 {
		numWords = 100 // Default value
	}

	// Start with first two words
	generatedWords := []string{words[0], words[1]}
	remaining := numWords - 2

	// Add remaining words
	for i := 0; i < remaining; i++ {
		position := rand.Intn(len(words))
		word := words[position]

		// Avoid repeating the same word consecutively
		if len(generatedWords) > 0 && generatedWords[len(generatedWords)-1] == word {
			i--
			continue
		}

		generatedWords = append(generatedWords, word)
	}

	sentences := []string{}
	current := 0
	wordsLeft := len(generatedWords)

	// Create sentences from the generated words
	for wordsLeft > 0 {
		sentenceLength := l.getRandomSentenceLength()

		if wordsLeft < sentenceLength {
			sentenceLength = wordsLeft
		}

		if sentenceLength < 1 {
			break
		}

		sentence := generatedWords[current : current+sentenceLength]
		sentence = l.punctuate(sentence)
		current += sentenceLength
		wordsLeft -= sentenceLength
		sentences = append(sentences, strings.Join(sentence, " "))
	}

	return strings.Join(sentences, " ")
}

// punctuate inserts commas and periods in the given sentence
func (l *LoremIpsum) punctuate(sentence []string) []string {
	wordLength := len(sentence)

	// Handle empty sentences or single word sentences
	if wordLength == 0 {
		return sentence
	}

	// End the sentence with a period
	sentence[wordLength-1] = sentence[wordLength-1] + "."

	if wordLength < 4 {
		// Capitalize the first letter of the first word
		if len(sentence[0]) > 0 {
			sentence[0] = strings.ToUpper(sentence[0][:1]) + sentence[0][1:]
		}
		return sentence
	}

	numCommas := l.getRandomCommaCount(wordLength)

	for i := 0; i <= numCommas; i++ {
		position := int(math.Round(float64(i) * float64(wordLength) / float64(numCommas+1)))

		if position < (wordLength-1) && position > 0 {
			// Add the comma
			sentence[position] = sentence[position] + ","
		}
	}

	// Capitalize the first letter of the first word
	if len(sentence[0]) > 0 {
		sentence[0] = strings.ToUpper(sentence[0][:1]) + sentence[0][1:]
	}

	return sentence
}

// getRandomCommaCount produces a random number of commas
func (l *LoremIpsum) getRandomCommaCount(wordLength int) int {
	base := 6.0 // Arbitrary

	average := math.Log(float64(wordLength)) / math.Log(base)
	standardDeviation := average / base

	return int(math.Round(l.gaussMS(average, standardDeviation)))
}

// getRandomSentenceLength produces a random sentence length
func (l *LoremIpsum) getRandomSentenceLength() int {
	length := int(math.Round(l.gaussMS(WordsPerSentenceAvg, WordsPerSentenceStd)))
	if length < 1 {
		return 1 // Ensure we always have at least one word per sentence
	}
	return length
}

// gauss produces a random number
func (l *LoremIpsum) gauss() float64 {
	return (rand.Float64()*2 - 1) +
		(rand.Float64()*2 - 1) +
		(rand.Float64()*2 - 1)
}

// gaussMS produces a random number with Gaussian distribution
func (l *LoremIpsum) gaussMS(mean, standardDeviation float64) float64 {
	return l.gauss()*standardDeviation + mean
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Parse command line arguments
	numWords := flag.Int("words", 100, "Number of words to generate")
	flag.Parse()

	// Create lorem ipsum generator
	lorem := &LoremIpsum{}

	// Generate and print the dog-themed lorem ipsum text
	fmt.Println(lorem.Generate(*numWords))
}
