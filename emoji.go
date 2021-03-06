package lineemoji

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/text/encoding/unicode/utf32"
)

// LineEmojiHexCode
// Reference: https://devdocs.line.me/files/emoticon.pdf
var LineEmojiHexCode = map[string]string{
	"(moon heart eyes)":      "0x00100078",
	"(moon smile)":           "0x00100090",
	"(cony kiss)":            "0x00100096",
	"(brown sleeping)":       "0x00100086",
	"(wink)":                 "0x00100005",
	"(sweat)":                "0x00100010",
	"(crossed-out eyes)":     "0x0010001B",
	"(dog)":                  "0x0010005E",
	"(moon laugh)":           "0x00100079",
	"(moon sweat)":           "0x00100091",
	"(cony grin)":            "0x00100097",
	"(james sweat)":          "0x00100087",
	"(heart eyes)":           "0x00100006",
	"(looking down)":         "0x00100011",
	"(horrified face)":       "0x0010001C",
	"(ghost)":                "0x001000A0",
	"(moon pleading)":        "0x0010007A",
	"(moon depressed)":       "0x00100092",
	"(cony freezing)":        "0x00100098",
	"(james exhausted)":      "0x00100088",
	"(kiss)":                 "0x00100007",
	"(frown)":                "0x00100012",
	"(angry frown)":          "0x0010001D",
	"(halloween)":            "0x001000A1",
	"(moon shocked)":         "0x0010007B",
	"(moon eating)":          "0x00100093",
	"(cony exhausted)":       "0x00100099",
	"(james panicking)":      "0x00100089",
	"(kiss2)":                "0x00100008",
	"(unpleasant face)":      "0x00100013",
	"(scary angry face)":     "0x0010001E",
	"(imp)":                  "0x00100024",
	"(moon cry)":             "0x0010007C",
	"(moon streaming tears)": "0x00100094",
	"(cony angry)":           "0x0010009A",
	"(james wink heart)":     "0x0010008A",
	"(big eyes)":             "0x00100009",
	"(sweat2)":               "0x00100014",
	"(sleeping face)":        "0x0010001F",
	"(skull)":                "0x001000A2",
	"(moon gaunt)":           "0x0010007D",
	"(moon laugh2)":          "0x00100095",
	"(cony sweat)":           "0x0010009B",
	"(james heart eyes)":     "0x0010008B",
	"(closed eyes smile)":    "0x0010000A",
	"(pale sweat)":           "0x00100015",
	"(sick)":                 "0x00100020",
	"(poop)":                 "0x001000A3",
	"(moon furious)":         "0x0010007E",
	"(cony laugh)":           "0x0010007F",
	"(cony sleeping)":        "0x0010009C",
	"(james thinking)":       "0x0010009F",
	"(big grin)":             "0x0010000B",
	"(pale)":                 "0x00100016",
	"(stone face)":           "0x00100021",
	"(fire)":                 "0x001000A4",
	"(moon smirk)":           "0x0010008C",
	"(cony sparkling eyes)":  "0x00100080",
	"(cony cute smile)":      "0x0010009D",
	"(laugh)":                "0x00100001",
	"(wink with tongue out)": "0x0010000C",
	"(distressed)":           "0x00100017",
	"(cat face)":             "0x00100022",
	"(yes)":                  "0x001000A5",
	"(moon big smile)":       "0x0010008D",
	"(cony kiss2)":           "0x00100081",
	"(cony furious)":         "0x0010009E",
	"(smile)":                "0x00100002",
	"(tongue out)":           "0x0010000D",
	"(tears)":                "0x00100018",
	"(eating face)":          "0x00100023",
	"(no)":                   "0x001000A6",
	"(moon oops)":            "0x0010008E",
	"(cony with tongue out)": "0x00100082",
	"(brown blush)":          "0x00100084",
	"(big smile)":            "0x00100003",
	"(looking away)":         "0x0010000E",
	"(crying)":               "0x00100019",
	"(pig)":                  "0x0010005D",
	"(toilet)":               "0x001000A7",
	"(moon wink)":            "0x0010008F",
	"(cony shocked)":         "0x00100083",
	"(brown surprised)":      "0x00100085",
	"(blushing smile)":       "0x00100004",
	"(smirk)":                "0x0010000F",
	"(tears of laughter)":    "0x0010001A",
	"(cat)":                  "0x0010005F",
	"(rage)":                 "0x00100026",
	"(surprise)":             "0x00100027",
	"(rain)":                 "0x001000AA",
	"(eighthnote)":           "0x00100039",
	"(calculator)":           "0x001000B4",
	"(bicycle)":              "0x0010004B",
	"(chocolate cake)":       "0x001000B6",
	"(rose stalk)":           "0x001000B9",
	"(videogame)":            "0x0010006E",
	"(drop)":                 "0x00100029",
	"(snow)":                 "0x001000AB",
	"(heart)":                "0x00100037",
	"(loveletter)":           "0x00100040",
	"(building)":             "0x0010004C",
	"(cocktail)":             "0x00100057",
	"(wave)":                 "0x00100064",
	"(cinema)":               "0x0010006F",
	"(dash)":                 "0x0010002A",
	"(cloud)":                "0x001000AC",
	"(broken heart)":         "0x00100038",
	"(pencil)":               "0x00100041",
	"(postoffice)":           "0x0010004D",
	"(beer)":                 "0x00100058",
	"(cigar)":                "0x00100065",
	"(headphones)":           "0x00100070",
	"(sleepy)":               "0x0010002B",
	"(thumbs up)":            "0x00100033",
	"(heart2)":               "0x001000AF",
	"(baseball)":             "0x00100042",
	"(hospital)":             "0x0010004E",
	"(coffee)":               "0x00100059",
	"(nosmoking)":            "0x00100066",
	"(clock)":                "0x00100071",
	"(lips)":                 "0x0010002C",
	"(thumbs down)":          "0x001000AD",
	"(three hearts)":         "0x001000B0",
	"(3 hearts)":             "0x001000B0",
	"(golfball)":             "0x00100043",
	"(school)":               "0x0010004F",
	"(medicine)":             "0x001000B7",
	"(highheels)":            "0x00100067",
	"(purse)":                "0x00100072",
	"(shiny)":                "0x0010002D",
	"(peace sign)":           "0x00100030",
	"(two hearts)":           "0x001000B1",
	"(2 hearts)":             "0x001000B1",
	"(tennisball)":           "0x00100044",
	"(bank)":                 "0x00100050",
	"(apple)":                "0x0010005B",
	"(lipstick)":             "0x00100068",
	"(crown)":                "0x00100073",
	"(eyes)":                 "0x0010002E",
	"(open hand)":            "0x00100031",
	"(star)":                 "0x001000B2",
	"(soccerball)":           "0x00100045",
	"(spa)":                  "0x00100051",
	"(chick)":                "0x0010005C",
	"(ribbon)":               "0x00100069",
	"(ring)":                 "0x00100074",
	"(ear)":                  "0x0010002F",
	"(fist)":                 "0x00100032",
	"(2 stars)":              "0x001000B3",
	"(bed)":                  "0x001000B5",
	"(gasstation)":           "0x00100053",
	"(flower)":               "0x00100060",
	"(camera)":               "0x0010006A",
	"(gift)":                 "0x00100075",
	"(thunder)":              "0x0010003A",
	"(clapping hands)":       "0x001000AE",
	"(PC)":                   "0x0010003B",
	"(train)":                "0x00100047",
	"(noodles)":              "0x00100054",
	"(fourleafclover)":       "0x00100061",
	"(bag)":                  "0x0010006B",
	"(birthday)":             "0x00100076",
	"(moon)":                 "0x001000A8",
	"(!)":                    "0x00100035",
	"(phone)":                "0x0010003C",
	"(car)":                  "0x00100049",
	"(hamburger)":            "0x00100055",
	"(tulip)":                "0x00100062",
	"(book)":                 "0x0010006C",
	"(lightbulb)":            "0x00100077",
	"(sun)":                  "0x001000A9",
	"(?)":                    "0x00100036",
	"(cellphone)":            "0x0010003D",
	"(plane)":                "0x0010004A",
	"(cake)":                 "0x00100056",
	"(rose)":                 "0x001000B8",
	"(TV)":                   "0x0010006D",
}

// GetEmoji return unicode character.
func GetEmoji(text string) (emoji []byte, err error) {
	hexstr := LineEmojiHexCode[text]
	if hexstr == "" {
		return nil, fmt.Errorf("Emoji %s not found.", text)
	}

	val := hexstr[2:]
	data, err := hex.DecodeString(val)
	if err != nil {
		return nil, err
	}
	utf32BEIB := utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM)
	dec := utf32BEIB.NewDecoder()

	return dec.Bytes(data)
}
