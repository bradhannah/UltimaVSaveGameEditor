package ultima_v_save

/*
enum <ubyte> PartyStatus { InTheParty = 0x00, HasntJoinedYet = 0xFF, AtTheInn = 0x01 };
enum <ubyte> Location { };
enum <ubyte> Boolean { False = 0, True = 1 };
enum <ubyte> BritOrUnderworld { Britannia = 0, Underworld = 0xFF };
enum <ubyte> MoonstoneStatus { Buried = 0, InInventory = 0xFF };
enum <ubyte> DungeonOrientation { North = 0, East = 1, South = 2, West = 2 };
enum <ubyte> DungeonType { Cave = 1, Mine = 2, Dungeon = 3 };
*/

/*
CharacterClass

enum <ubyte> Class { Avatar = 'A', Fighter = 'F', Bard = 'B', Wizard = 'M' };
*/

type CharacterClass byte

const (
	Avatar  CharacterClass = 'A'
	Fighter CharacterClass = 'F'
	Bard    CharacterClass = 'B'
	Wizard  CharacterClass = 'M'
)

var CharacterClassMap = map[CharacterClass]string{
	Avatar:  "Avatar",
	Fighter: "Fighter",
	Bard:    "Bard",
	Wizard:  "Wizard",
}

/*
CharacterStatus

enum <ubyte> Status { Good = 'G', Poisoned = 'P', Charmed = 'C', Asleep = 'S', Dead = 'D' };
*/
type CharacterStatus byte

const (
	Good     CharacterStatus = 'G'
	Poisoned CharacterStatus = 'P'
	Charmed  CharacterStatus = 'C'
	Sleep    CharacterStatus = 'S'
	Dead     CharacterStatus = 'D'
)

var CharacterStatusMap = map[CharacterStatus]string{
	Good:     "Good",
	Poisoned: "Poisoned",
	Charmed:  "Charmed",
	Sleep:    "Sleep",
	Dead:     "Dead",
}

/*
Gender
enum <ubyte> Gender { Male = 0x0B, Female = 0x0C };
*/
type CharacterGender byte

const (
	Male   CharacterGender = 0x0B
	Female CharacterGender = 0x0C
)

var CharacterGenderMap = map[CharacterGender]string{
	Male:   "Male",
	Female: "Female",
}

/*
Other smaller definitions
*/
const (
	brit BritOrUnderworld = 'F'
)
