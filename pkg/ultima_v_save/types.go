package ultima_v_save

type StartingMemoryAddressUb uint16
type StartingMemoryAddressU16 uint16

const NPlayers = 6
const NMaxPlayerNameSize = 9

const (
	ubPartyMembers StartingMemoryAddressUb = 0x2B5
	ubActivePlayer                         = 0x2D5
)

const (
	u16CurrentYear StartingMemoryAddressU16 = 0x2CE
)

type ByteIdToFriendly struct {
	identifier   byte
	FriendlyName string
}

type CharacterStats struct {
	Name      string
	Class     ByteIdToFriendly
	CurrentMp byte
	MaximumMp byte

	CurrentHp byte
	MaximumHp byte

	ExperiencePoints uint16
}

type PartyStatus byte

const (
	InTheParty     PartyStatus = 0x00
	HasntJoinedYet             = 0xFF
	AtTheInn                   = 0x01
)

type Location byte

//goland:noinspection ALL
const (
	Britannia_Underworld  Location = 0x00
	Moonglow                       = 1
	Britain                        = 2
	Jhelom                         = 3
	Yew                            = 4
	Minoc                          = 5
	Trinsic                        = 6
	Skara_Brae                     = 7
	New_Magincia                   = 8 // Town
	Fogsbane                       = 9
	Stormcrow                      = 10
	Greyhaven                      = 11
	Waveguide                      = 12
	Iolos_Hut                      = 13
	Suteks_Hut                     = 14
	SinVraals_Hut                  = 15
	Grendels_Hut                   = 16 // Dwelling
	Lord_Britishs_Castle           = 17
	Palace_of_Blackthorn           = 18
	West_Britanny                  = 19
	North_Britanny                 = 20
	East_Britanny                  = 21
	Paws                           = 22
	Cove                           = 23
	Buccaneers_Den                 = 24 // Castle
	Ararat                         = 25
	Bordermarch                    = 26
	Farthing                       = 27
	Windemere                      = 28
	Stonegate                      = 29
	Lycaeum                        = 30 // Keep
	Empath_Abbey                   = 31
	Serpents_Hold                  = 32
	Deceit                         = 33 // Dungeons
	Despise                        = 34
	Destard                        = 35
	Wrong                          = 36
	Covetous                       = 37
	Shame                          = 38
	Hythloth                       = 39
	Doom                           = 40
	Combat_resting_shrine          = 41
)

type BritOrUnderworld byte

const (
	Britannia  BritOrUnderworld = 0x00
	Underworld                  = 0xFF
)

type MoonstoneStatus byte

const (
	Buried      MoonstoneStatus = 0x00
	InInventory                 = 0xFF
)

type SaveGame struct {
	Characters       [NPlayers]PlayerCharacter
	Location         Location
	BritOrUnderworld BritOrUnderworld
	MoonstoneStatus  MoonstoneStatus
}

type PlayerCharacter struct {
	Name         [NMaxPlayerNameSize]byte
	Gender       CharacterGender
	Class        CharacterClass
	Status       CharacterStatus
	Strength     byte
	Dexterity    byte
	Intelligence byte
	CurrentMp    byte
	CurrentHp    uint16
	MaxHp        uint16
	Exp          uint16
	Level        byte
	MonthsAtInn  byte
	Unknown      byte
	Helmet       byte
	Armor        byte
	Weapon       byte
	Shield       byte
	Ring         byte
	Amulet       byte
	PartyStatus  PartyStatus
}
