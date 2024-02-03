package ultima_v_save

/*
enum <ubyte> Gender { Male = 0x0B, Female = 0x0C };
enum <ubyte> Class { Avatar = 'A', Fighter = 'F', Bard = 'B', Wizard = 'M' };
enum <ubyte> Status { Good = 'G', Poisoned = 'P', Charmed = 'C', Asleep = 'S', Dead = 'D' };
enum <ubyte> PartyStatus { InTheParty = 0x00, HasntJoinedYet = 0xFF, AtTheInn = 0x01 };
enum <ubyte> Location { };
enum <ubyte> Boolean { False = 0, True = 1 };
enum <ubyte> BritOrUnderworld { Britannia = 0, Underworld = 0xFF };
enum <ubyte> MoonstoneStatus { Buried = 0, InInventory = 0xFF };
enum <ubyte> DungeonOrientation { North = 0, East = 1, South = 2, West = 2 };
enum <ubyte> DungeonType { Cave = 1, Mine = 2, Dungeon = 3 };
*/

var CharacterClasses []CharacterClass = []CharacterClass{
	{
		identifier:   'A',
		FriendlyName: "Avatar",
	},
	{
		identifier:   'F',
		FriendlyName: "Fighter",
	},
	{
		identifier:   'B',
		FriendlyName: "Bard",
	},
	{
		identifier:   'M',
		FriendlyName: "Wizard",
	},
}
